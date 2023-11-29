<script lang="ts">
	import { onMount } from 'svelte';
	import TileComponent from '../components/TileComponent.svelte';
	import type { Tile } from '../models/Tile';

	export let data;
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
		tiles = data.tiles;
		console.log(tiles);
	});
</script>

<div class="bg-black h-[500px] relative">
	{#each tiles as tile}
		<TileComponent {...tile} />
	{/each}
</div>
