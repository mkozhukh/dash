<script>
	import Notice from "./Notice.svelte";

	export let data;
</script>

<div class="area">
	{#each $data as notice (notice.id)}
		<Notice {notice} />
	{/each}
</div>

<style>
	.area {
		position: fixed;
		top: 0px;
		right: 0px;
	}
</style>
