import { writable } from "svelte/store";
import { locate, locateID } from "../locate";

export function toggleID(value) {
	const { subscribe, set } = writable(value);
	return {
		subscribe,
		toggle: id => set((value = value && value == id ? null : id * 1)),
		off: id => {
			if (!id || value == id) set((value = null));
		},
	};
}

export function click(node, handlers) {
	let lastClick = null;

	function handleClick(ev) {
		const node = locate(ev);
		if (!node) return;

		let test = ev.target;
		while (test != node) {
			let action = test.dataset ? test.dataset.action : null;
			if (action) {
				if (handlers[action]) handlers[action](node.dataset.id * 1);
				lastClick = new Date();
				break;
			}
			test = test.parentNode;
		}
		if (handlers.click) handlers.click(node.dataset.id * 1);
	}
	function handleDblClick(ev) {
		if (lastClick && new Date() - lastClick < 200) return;
		const id = locateID(ev);
		if (id && handlers.dblclick) handlers.dblclick(id);
	}
	node.addEventListener("click", handleClick);
	node.addEventListener("dblclick", handleDblClick);
}
