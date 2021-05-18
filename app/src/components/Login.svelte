<script>
    import { onMount } from "svelte";

    import Button from "wx/Button.svelte"
    import Text from "wx/Text.svelte"

    export let remote;
    export let key;

    let value = "";
    let error = "";
    let id;

    function handleEnter(ev){
        if (ev.key === "Enter")
            doLogin();
    }

    function doLogin(){
        remote.api.admin.Login(value).then(res => {
            if (res !== ""){
                sessionStorage.setItem("jwt", res);
                key = res;
            }
        }, err => {
            error = err;
            focus();
        });
    }

    function focus(){
        document.getElementById(id).focus();
    }
    onMount(focus);
</script>

<div class='popup' on:keydown={handleEnter}>
    <p>Enter the access key:</p>
    <Text bind:id bind:value></Text>
    <Button type="primary" click={doLogin}>Login</Button>
    {#if error}
    <div class="error">{error}</div>
    {/if}
</div>

<style>
    .popup {
        position: fixed;
        top: 0px;
        left: 50%;
        padding: 20px;
        margin-left: -150px;
        margin-top: -10px;
        width: 300px;
        box-shadow: var(--wx-card-shadow);
    }
    .error {
        background: red;
        color: white;
        padding: 4px 5px;
        margin: 5px 5px 0 0;
        border-radius: 4px;
        text-align: center;
    }
</style>