<script lang="ts">
	import TileComponent from '../components/TileComponent.svelte';
	import type { Tile } from '../models/Tile';

	let tiles: Tile[] = [
		{ x: 0, y: 0 },
		{ x: 0, y: 1 },
		{ x: 0, y: 2 },
		{ x: 0, y: 3 },
		{ x: 1, y: 0 },
		{ x: 1, y: 1 },
		{ x: 1, y: 2 },
		{ x: 1, y: 3 },
		{ x: 2, y: 0 },
		{ x: 2, y: 1 },
		{ x: 2, y: 2 },
		{ x: 2, y: 3 }
	];
</script>

<div class="bg-black h-[500px] relative">
	{#each tiles as tile}
		<TileComponent {...tile} />
	{/each}
</div>
