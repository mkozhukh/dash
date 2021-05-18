<script>
	export let type = "";
	export let click;
	export let shape = "round";
</script>

<button
	class="btn {shape}"
	class:primary={type === 'primary'}
	class:danger={type === 'danger'}
	class:link={type === 'link'}
	class:selected={type === 'selected'}
	on:click={click}>
	<slot />
</button>

<style>
	.btn {
		padding: 8px 16px;
		border: none;
		border-radius: 4px;
		cursor: pointer;
	}

	.btn:active {
		outline: none;
		opacity: 0.7;
	}

	.btn:focus {
		outline: none;
	}
	:global(.square).btn {
		border-radius: 0px;
	}

	.selected {
		background-color: var(--wx-selected-color);
	}

	.primary {
		color: var(--wx-primary-font-color);
		background-color: var(--wx-primary-color);
	}

	.primary:hover {
		color: var(--wx-primary-font-color);
		background-color: var(--wx-primary-color-hover);
	}

	.danger {
		color: var(--wx-danger-font-color);
		background-color: var(--wx-danger-color);
	}

	.danger:hover {
		background-color: var(--wx-danger-color-hover);
	}

	.link {
		background: var(--wx-back-color);
		color: var(--wx-link-color);
		padding-left: 0px;
		padding-right: 5px;
		font: var(--wx-font);
	}
</style>
