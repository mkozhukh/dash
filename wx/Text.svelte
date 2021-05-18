<script>
	import { uid } from "helpers/common";

	export let value = "";
	export let id = uid();
</script>

<input bind:value {id} />

<style>
	input {
		padding: 4px;
		outline: none;
		flex: 1;
		font: var(--wx-font);
		color: var(--wx-color);

		border: none;
		border-bottom: 1px solid transparent;
		background-color: transparent;
	}
	input:focus {
		border-bottom: 1px solid var(--wx-primary-color);
	}
</style>
