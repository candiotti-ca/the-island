<script lang="ts">
	const size = 50;
	const width = size * 0.8;
	const height = size * 0.46;

	export let x: number = 0;
	export let y: number = 0;

	let top = 250 + y * height * 1.7;
	let left = (y % 2 == 0 ? 1 : 0.45) * width + x * width + x * 4 + 600;
</script>

<div class="tile" style="width:{width}px; height:{height}px; top:{top}px; left:{left}px;"></div>

<style>
	.tile {
		position: absolute;
		background-color: #fda3b2;
		cursor: pointer;
	}

	.tile:before,
	.tile:after {
		position: absolute;
		content: '';
		background: inherit;
		height: 100%;
		width: 100%;
		transform-origin: center;
	}

	.tile:before {
		transform: rotateZ(60deg);
	}

	.tile:after {
		transform: rotateZ(-60deg);
	}
</style>
