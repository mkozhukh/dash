<script>
	import App from "../src/components/App.svelte";
	export let remote;

	window._x_remote = remote;
</script>

<App {remote}/>
