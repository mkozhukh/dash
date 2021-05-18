<script>
	import { fade } from "svelte/transition";
	export let notice = {};

	function onRemove() {
		if (notice.remove) notice.remove();
	}
</script>

<div
	class="message {notice.type || 'info'}"
	role="status"
	aria-live="polite"
	transition:fade>
	<div class="text">{notice.text}</div>
	<button class="close" on:click={onRemove}> &times; </button>
</div>

<style>
	.message {
		background: var(--wx-back-color);
		box-shadow: var(--wx-card-shadow);
		border-radius: 3px;
		margin: 6px 12px;
		display: flex;
		overflow: hidden;
	}

	.message .text {
		width: 240px;
		padding: 10px;
		word-wrap: break-word;
	}

	.message .close {
		display: block;
		width: 40px;
		padding: 0;
		margin: 0;
		border: none;
		outline: none;
		background: none;

		cursor: pointer;
		font-size: 20px;
	}

	.message.warning {
		background: var(--wx-warning-color);
	}

	.message.success {
		background: var(--wx-success-color);
	}

	.message.error {
		background: var(--wx-danger-color);
		color: var(--wx-danger-font-color);
	}
	.message.error .close {
		color: var(--wx-danger-font-color);
	}
</style>
