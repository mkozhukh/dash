<script>
	import { onMount } from "svelte";

	export let value = "";
	export let ok;
	export let cancel;

	let input;
	onMount(() => {
		input.select();
		input.focus();
	});
</script>

<div class="hover" />
<div class="center">
	<div class="window">
		<div class="header">
			<slot />
		</div>
		<input bind:this={input} bind:value />
		<div class="footer">
			<button on:click={() => ok(value)}> Ok </button>
			<button on:click={() => cancel()}> Cancel </button>
		</div>
	</div>
</div>

<style>
	.center {
		display: flex;
		align-items: center;
		justify-content: center;
	}

	.window {
		margin-top: -50px;
		width: 270px;
		background: #fff;
		padding: 3px;
		border-radius: 4px;
		box-shadow: 0px 0px 10px 1px #888;
	}
	.header {
		font-size: 20px;
		line-height: 21px;
		padding: 10px;
		background: #ddd;
		margin-bottom: 10px;
	}
	input {
		outline: none;
		border: none;
		width: 100%;
		font-size: 16px;
		padding: 0 10px;
	}
	button {
		margin: 15px 10px 5px 10px;
		width: 100px;
		padding: 5px;
		border: 1px solid silver;
	}
	button:nth-child(2) {
		float: right;
	}

	.hover,
	.center {
		position: fixed;
		top: 0px;
		left: 0px;
		width: 100%;
		height: 100%;
	}

	.hover {
		background: rgba(0, 0, 0, 0.2);
	}
</style>
