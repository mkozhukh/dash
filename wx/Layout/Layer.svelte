<script>
	import { fly } from "svelte/transition";

	export let bottom = false;
</script>

{#if !bottom}
	<div class="shadow" />
{/if}
<div
	class="box"
	class:relative={!bottom}
	transition:fly={{ x: 100, duration: 500 }}>
	<div class="layer">
		<slot />
	</div>
</div>

<style>
	.box {
		position: relative;
		border-top-left-radius: 4px;
		background: var(--wx-back-color);
		box-shadow: var(--wx-popup-shadow);
		min-height: 100%;
	}
	.shadow {
		width: 100%;
		height: 100%;
		position: absolute;
		top: 0px;
		left: 0px;
		background: var(--wx-layer-shadow);
	}
	.relative {
		position: absolute;
		top: 40px;
		left: 20px;
		width: calc(100% - 20px);
		min-height: calc(100% - 40px);
	}
</style>
