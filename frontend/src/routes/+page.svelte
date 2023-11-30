<script lang="ts">
	import { onMount } from 'svelte';
	import TileComponent from '../components/TileComponent.svelte';
	import type { Tile } from '../models/Tile';

	let tiles: Tile[] = [];

	onMount(async () => {
		fetch('http://localhost:8080/map')
			.then((response) => response.json())
			.then((param: Tile[]) => {
				tiles = param;
				console.log(tiles);
			})
			.catch((error) => console.error(error));
	});
</script>

<div class="bg-black h-[600px] relative">
	{#each tiles as tile}
		<TileComponent {...tile} />
	{/each}
</div>
