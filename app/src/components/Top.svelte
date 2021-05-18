<script>
    import { getContext } from "svelte";

    import Login from "./Login.svelte";
    import Page from "./Page.svelte";

    export let remote;
	
	const helpers = getContext("wx-helpers");
	remote.onCall = (_, p) =>
		p.catch(msg => helpers.showNotice({ text: msg, type: "warning" }));
	remote.onError = msg => {
		helpers.showNotice({ text: "communication error", type: "error" });
		console.error(msg);
	};

    let key = sessionStorage.getItem("jwt");
</script>

{#if !key}
    <Login {remote} bind:key />
{:else}
    <Page {remote} {key} />
{/if}