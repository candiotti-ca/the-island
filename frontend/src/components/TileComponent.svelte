<script lang="ts">
	import grassTexture from '../assets/grass.jpg';
	import rockTexture from '../assets/rock.jpg';
	import sandTexture from '../assets/sand.jpg';
	import waterTexture from '../assets/water.jpg';
	import { TileType } from '../models/TileType';

	const size = 80;
	const width = size * 0.8;
	const height = size * 0.46;

	export let q: number = 0;
	export let r: number = 0;
	export let type: TileType = TileType.WATER;

	$: top = () => {
		const middleOfTheMap = 400;

		return height * ((3 / 2) * r) + middleOfTheMap;
	};

	/**
	 * Computes pixel coordinates of the current tile. It depends whether the map is
	 * pointy top or flat top and whether the map is odd or even.
	 */
	$: left = () => {
		const middleOfTheMap = 850;

		return height * (Math.sqrt(3) * q + (Math.sqrt(3) / 2) * r) + middleOfTheMap;
	};

	$: backgroundUrl = () => {
		switch (type) {
			case TileType.GRASS:
				return grassTexture;
			case TileType.ROCK:
				return rockTexture;
			case TileType.SAND:
				return sandTexture;
			default:
				return waterTexture;
		}
	};

	$: altBackground = () => {
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
	};
</script>

<button
	class="tile"
	style="width:{width}px; height:{height}px; top:{top()}px; left:{left()}px; background: url({backgroundUrl()}); background-color: {altBackground()}; z-index:{type};"
	title="[{q},{r}]"
	on:click
	data-testid="tile"
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
