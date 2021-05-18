<script>
    import Globals from "wx/Globals.svelte";
    import Top from "./Top.svelte";
    export let remote;
</script>

<Globals>
    <Top {remote}/>
</Globals>