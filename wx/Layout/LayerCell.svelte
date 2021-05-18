<script>
	import Button from "../Button.svelte";

	export let prefix;
	export let title;
	export let backText;
	export let back;
	export let remove = null;

	let edit = false;
</script>

<div class="header" style="padding: {edit ? 2 : 10}px;">
	{#if backText}
		<button class="left" on:click={() => back()}>{backText}</button>
	{/if}
	<div class="content">
		<span class="type">{prefix}</span>
		<span class="name">{title || '(no title)'}</span>
	</div>
	{#if remove}
		<div class="right">
			<Button type="danger" click={remove}>Delete Model</Button>
		</div>
	{/if}
</div>
<div class="body">
	<slot />
</div>

<style>
	.body {
		padding: 0 10px 50px 90px;
	}
	.right {
		margin-left: auto;
	}
	.content {
		margin-left: 10px;
	}
	.header {
		margin-left: 90px;
		font-size: 20px;
		line-height: 20px;
		padding: 10px;
		background: var(--wx-header-color);
		box-shadow: var(--wx-bottom-shadow);
		margin-bottom: 10px;
		position: relative;
		overflow: visible;

		align-items: center;
		display: flex;
	}
	.left {
		position: absolute;
		left: -75px;
		top: 12px;

		background: var(--wx-back-color);
		color: var(--wx-link-color);
		border: none;

		cursor: pointer;
	}

	.header .type {
		color: var(--wx-header-font-alt-color);
	}
	.header .name {
		color: var(--wx-header-font-color);
	}
</style>
