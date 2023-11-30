<script lang="ts">
	import { onMount } from 'svelte';
	import TileComponent from '../components/TileComponent.svelte';
	import { Tile, type BackendTile } from '../models/Tile';

	let tiles: Tile[] = [
		// { x: 0, y: 0 },
		// { x: 0, y: 1 },
		// { x: 0, y: 2 },
		// { x: 0, y: 3 },
		// { x: 1, y: 0 },
		// { x: 1, y: 1 },
		// { x: 1, y: 2 },
		// { x: 1, y: 3 },
		// { x: 2, y: 0 },
		// { x: 2, y: 1 },
		// { x: 2, y: 2 },
		// { x: 2, y: 3 }
	];

	onMount(async () => {
		fetch('http://localhost:8080/map')
			.then((response) => response.json())
			.then((tiles: BackendTile[]) => tiles.map((tile) => new Tile(tile)))
			.then((param: Tile[]) => {
				tiles = param;
				console.log(tiles);
			})
			.catch((error) => console.error(error));
	});
</script>

<div class="bg-black h-[500px] relative">
	{#each tiles as tile}
		<TileComponent {...tile} />
	{/each}
</div>
