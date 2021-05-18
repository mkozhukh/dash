<script>
	import { setContext } from "svelte";
	import { writable } from "svelte/store";

	import Prompt from "./Prompt.svelte";
	import Notices from "./Notices.svelte";
	import { uid } from "helpers/common";

	let prompt = null;
	function showPrompt(message, value) {
		prompt = { message, value: value || "" };
		return new Promise((res, rej) => {
			prompt.resolve = v => {
				prompt = null;
				res(v);
			};
			prompt.reject = v => {
				prompt = null;
				rej(v);
			};
		});
	}

	let notices = writable([]);
	function showNotice(msg) {
		msg = { ...msg };
		msg.id = msg.id || uid();
		msg.remove = () =>
			notices.update(data => data.filter(a => a.id !== msg.id));

		if (msg.expire != -1) {
			setTimeout(msg.remove, msg.expire || 5000);
		}
		notices.update(data => [...data, msg]);
	}

	setContext("wx-helpers", {
		showPrompt,
		showNotice,
	});
</script>

<slot />
{#if prompt}
	<Prompt value={prompt.value} ok={prompt.resolve} cancel={prompt.reject}>
		{prompt.message}
	</Prompt>
{/if}
<Notices data={notices} />
