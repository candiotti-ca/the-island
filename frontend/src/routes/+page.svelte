<div class="hexagon"></div>

<style>
	.hexagon {
		margin: 40px;
		position: relative;
		width: 92px; /* 46% de 200 */
		height: 160px; /*80% de 200 */
		background-color: #fda3b2;
		cursor: pointer;
		border: 3px solid black;
	}

	.hexagon:before,
	.hexagon:after {
		position: absolute;
		content: '';
		background: inherit;
		height: 100%;
		width: 100%;
		transform-origin: center;
		border-top: 3px solid black;
		border-bottom: 3px solid black;
	}
	.hexagon:before {
		transform: rotateZ(60deg);
	}
	.hexagon:after {
		transform: rotateZ(-60deg);
	}
</style>
