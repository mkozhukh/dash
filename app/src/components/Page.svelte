<script>
    import Button from "wx/Button.svelte";
    import Layer from "wx/Layout/Layer.svelte";
    import LayerCell from "wx/Layout/LayerCell.svelte";

    export let remote;
    export let key;

    const api = remote.api.admin;

    let server = "";
    let stats = [];
    let actions = [];

    remote.on("update", obj => {
        switch(obj.type){
            case "config":
                server = obj.server
                actions = obj.commands;
                break;
            case "info":
                stats = obj.info;
                break;
            case "exec":
                if (progress.uid !== obj.id) return;
                progress.lines = [ ...progress.lines, { status: obj.status, message: obj.message }];
                break;
            case "done":
                if (progress.uid !== obj.id) return;
                progress.done = { status: obj.status, message: obj.message };
                break;
        }
    })
    // FIXME - on must return a promise 
    setTimeout(() => api.GetInfo(key), 300);

    var progress = null;
    function start(line){
        const uid = (new Date()).valueOf().toString();
        progress = { ...line, uid, lines:[], done:null };

        api.Exec(line.id, uid, key);
    }
</script>

<h2>{server}</h2>

<h3>Server stats</h3>
{#each stats as line}
    <div class="line">
        <span class="name">{line.name}</span>
        <span class="value">{line.value}</span>
    </div>
{/each}

<h3>Actions</h3>
{#each actions as line}
    <div class="line">
        <div class="border">
        {#if line.details}
            <pre>{line.details||""}</pre>
        {/if}
        <Button click={() => start(line)} type={line.danger?"danger":"primary"}>{line.name}</Button>
        </div>
    </div>
{/each}

{#if progress}
    <Layer>
        <LayerCell prefix="exec" title={progress.name} backText="close" back={() => progress = null}>
            {#if progress.done}
            <div class="end {progress.done.status}">
                {progress.done.message || "Done"}
            </div>
            {:else}
            <div class="end wait">
                executing...
            </div>
            {/if}
            {#if progress.lines.length}
                {#each progress.lines as line}
                    <pre class={line.status}>{line.message || "(no response)"}</pre>
                {/each}
            {:else}
                processing...
            {/if}
        </LayerCell>
    </Layer>
{/if}


<style>
    h2{
        text-align: center;
        font-weight: 500;
        font-size: 30px;
        padding: 10px 0; 
        margin: 0 0 10px 0;
        
        box-shadow: var(--wx-bottom-shadow);
        background: var(--wx-secondary-back-color);
    }
    h3{
        padding: 5px 10px;
        margin: 30px 15px 10px;
        border-bottom: 1px solid var(--wx-border-color);
    }
    .line{
        margin-left: 25px;
        margin-bottom: 5px;
    }
    .border{
        border-left: 3px solid var(--wx-border-color);
        padding-left: 10px;
    }
    .name{
        font-weight: 400;
    }
    .value{
        font-weight: 500;
    }
    pre{
        font-family: Roboto;
        margin-bottom: 5px;
    }
    pre.ok {
        border-left: 3px solid #77ea77;
        padding-left: 10px;
        font-family: monospace;
    }
    pre.error {
        border-left: 3px solid #ea7777;
        padding-left: 10px;
        font-family: monospace;
    }
    .end.ok {
        background: #04b304;
        padding: 5px 10px;
        border-radius: 3px;
        color: white;
    }
    .end.error {
        background: #b32704;
        padding: 5px 10px;
        border-radius: 3px;
        color: white;
    }
    .end.wait {
        background: #eec50f;
        padding: 5px 10px;
        border-radius: 3px;
    }
</style>