export function locate(el) {
	let node = !el.tagName && el.target ? el.target : el;

	while (node) {
		if (node.getAttribute) {
			const id = node.getAttribute("data-id");
			if (id) return node;
		}

		node = node.parentNode;
	}

	return null;
}

export function locateID(el) {
	const node = locate(el);
	if (node) return node.getAttribute("data-id") * 1;

	return null;
}
