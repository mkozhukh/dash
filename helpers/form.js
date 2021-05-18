import { writable } from "svelte/store";

export function form(v, changes, config) {
	let ready = false;
	let timer = null;
	const store = writable(v);

	store.reset = function (v) {
		ready = false;
		store.set(v);
	};
	store.subscribe(v => {
		if (ready) {
			if (v) {
				if (!config || !config.debounce) changes(v);
				else {
					clearTimeout(timer);
					timer = setTimeout(() => changes(v), config.debounce);
				}
			}
		} else ready = true;
	});

	return store;
}
