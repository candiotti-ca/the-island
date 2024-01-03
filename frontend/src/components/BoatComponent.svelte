<script lang="ts">
	import { selectedEntity } from '../store/store';

	export let top: number = 0;
	export let left: number = 0;
	export let id: string = '';

	$: selected = $selectedEntity?.entity.id == id;
</script>

<button on:click>
	<svg
		style="top:{top}px; left:{left + 7}px;"
		class="absolute z-20"
		height="50px"
		width="50px"
		viewBox="0 -1.8 20 20"
		xmlns="http://www.w3.org/2000/svg"
		fill="#000"
		><g id="SVGRepo_bgCarrier" stroke-width="0"></g><g
			id="SVGRepo_tracerCarrier"
			stroke-linecap="round"
			stroke-linejoin="round"
		></g><g id="SVGRepo_iconCarrier">
			<g id="boat-2" transform="translate(-2 -3.607)">
				<path
					id="secondary"
					fill="#795146"
					d="M16.25,8a5.73,5.73,0,0,0,3.2-1h0A1,1,0,0,1,21,7.86v1.2A4,4,0,0,1,17.06,13H6.94A4,4,0,0,1,3,9.06V7.86A1,1,0,0,1,4.55,7h0a5.73,5.73,0,0,0,3.2,1Z"
				></path>
				<path
					id="primary"
					d="M16.25,8a5.73,5.73,0,0,0,3.2-1h0A1,1,0,0,1,21,7.86v1.2A4,4,0,0,1,17.06,13H6.94A4,4,0,0,1,3,9.06V7.86A1,1,0,0,1,4.55,7h0a5.73,5.73,0,0,0,3.2,1Zm0-3L8,16.79M6,17.85,6.57,17A1,1,0,0,1,8,16.79H8a1,1,0,0,1,.24,1.39l-.6.82Z"
					fill="none"
					stroke={selected ? '#f00' : '#000'}
					stroke-linecap="round"
					stroke-linejoin="round"
					stroke-width="0.7"
				></path>
			</g>
		</g></svg
	>
</button>
