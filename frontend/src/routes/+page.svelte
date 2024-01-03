<script lang="ts">
	import { onMount } from 'svelte';
	import TileComponent from '../components/TileComponent.svelte';
	import type { SelectedEntity } from '../models/SelectedEntity';
	import type { Tile } from '../models/Tile';
	import { selectedEntity } from '../store/store';

	let tiles: Tile[] = [];

	onMount(async () => {
		fetch('http://localhost:8080/map')
			.then((response) => response.json())
			.then((body) =>
				Object.entries(body).map((entry) => {
					const coord: number[] = JSON.parse(entry[0]);
					const tile = entry[1] as any;

					return {
						q: coord[0],
						r: coord[1],
						...tile
					};
				})
			)
			.then((param: Tile[]) => (tiles = param))
			.catch((error) => console.error(error));
	});

	function clicOnTile(tile: Tile): void {
		if ($selectedEntity) {
			moveEntity($selectedEntity, tile);
		} else {
			flipTile(tile);
		}
	}

	function flipTile(tile: Tile): void {
		fetch(`http://localhost:8080/tiles/${tile.q}/${tile.r}/flip`, { method: 'PATCH' })
			.then((response) => {
				if (response.status == 200) {
					return response.json();
				}
				return Promise.reject(new Error('server did not return 200'));
			})
			.then((newType) => {
				tiles = tiles.map((t) => {
					if (t.id == tile.id) {
						t.type = newType;
					}

					return t;
				});
			})
			.catch((error) => console.error(error));
	}

	function selectEntity(detail: any, tile: Tile): void {
		selectedEntity.update((selected) => {
			if (!selected || selected.entity.id != detail.entity.id) {
				return { ...detail, tile };
			}
			return undefined;
		});
	}

	function moveEntity(entity: SelectedEntity, destination: Tile): void {
		const from = entity.tile;
		fetch(
			`http://localhost:8080/${entity.type}/${from.q}/${from.r}/move/${destination.q}/${destination.r}`,
			{ method: 'PATCH' }
		)
			.then((response) => {
				if (response.status == 200) {
					return response.json();
				}
				return Promise.reject(new Error('server did not return 200'));
			})
			.then((destTile) => {
				selectedEntity.set(undefined);
				tiles = tiles.map((t) => {
					if (t.id == from.id) {
						t.boat = undefined;
					}
					if (t.id == destination.id) {
						Object.assign(t, destTile);
					}

					return t;
				});
			})
			.catch((error) => console.error(error));
	}
</script>

<div class="h-full relative">
	{#each tiles as tile}
		<TileComponent
			{...tile}
			on:click={() => clicOnTile(tile)}
			on:select={(event) => {
				selectEntity(event.detail, tile);
			}}
		/>
	{/each}
</div>
