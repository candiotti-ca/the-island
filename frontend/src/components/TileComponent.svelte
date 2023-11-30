<script lang="ts">
	import { TileType } from '../models/TileType';

	const size = 50;
	const width = size * 0.8;
	const height = size * 0.46;

	export let x: number = 0;
	export let y: number = 0;
	export let type: TileType = TileType.WATER;

	function onClick(): void {
		console.log('X:', x, 'Y:', y);
	}

	function getTop(): number {
		const middleOfTheMap = 250;
		const border = 1.7;

		return middleOfTheMap + y * height * border;
	}

	function getLeft(): number {
		const middleOfTheMap = 600;
		const offset = (y % 2 == 0 ? 1 : 0.45) * width;
		const border = x * 4;

		return offset + x * width + border + middleOfTheMap;
	}

	function getBackground(): string {
		switch (type) {
			case TileType.GRASS:
				return 'green';
			case TileType.ROCK:
				return 'gray';
			case TileType.SAND:
				return 'yellow';
			default:
				return 'blue';
		}
	}
</script>

<button
	class="tile"
	style="width:{width}px; height:{height}px; top:{getTop()}px; left:{getLeft()}px; background-color:{getBackground()};"
	on:click={onClick}
></button>

<style>
	.tile {
		position: absolute;
	}

	.tile:before,
	.tile:after {
		position: absolute;
		content: '';
		background: inherit;
		height: 100%;
		width: 100%;
		left: 0%;
		top: 0%;
	}

	.tile:before {
		transform: rotateZ(60deg);
	}

	.tile:after {
		transform: rotateZ(-60deg);
	}
</style>
