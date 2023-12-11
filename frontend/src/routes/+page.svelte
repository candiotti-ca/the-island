<script lang="ts">
	import { onMount } from 'svelte';
	import TileComponent from '../components/TileComponent.svelte';
	import type { Tile } from '../models/Tile';
	import type { TileType } from '../models/TileType';

	let tiles: Tile[] = [];

	onMount(async () => {
		fetch('http://localhost:8080/map')
			.then((response) => response.json())
			.then((body) =>
				Object.entries(body).map((entry) => {
					const coord: number[] = JSON.parse(entry[0]);
					const type: TileType = (entry[1] as any).type;

					return {
						q: coord[0],
						r: coord[1],
						type
					};
				})
			)
			.then((param: Tile[]) => (tiles = param))
			.catch((error) => console.error(error));
	});

	function flipTile(tile: Tile): void {
		fetch(`http://localhost:8080/tiles/${tile.q}/${tile.r}/flip`, { method: 'PATCH' })
			.then((response) => {
				if (response.status == 200) {
					return response.json();
				}
				return Promise.reject(new Error('server does not return 200'));
			})
			.then((newType) => {
				tiles = tiles.map((t) => {
					if (t.q == tile.q && t.r == tile.r) {
						t.type = newType;
					}

					return t;
				});
			})
			.catch((error) => console.error(error));
	}
</script>

<div class="bg-black h-[600px] relative">
	{#each tiles as tile}
		<TileComponent {...tile} on:click={() => flipTile(tile)} />
	{/each}
</div>
