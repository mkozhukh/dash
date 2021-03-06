module.exports = {
	env: {
		browser: true,
		es6: true,
	},
	extends: ["eslint:recommended", "prettier"],
	parserOptions: {
		ecmaVersion: 2018,
		sourceType: "module",
	},
	plugins: ["svelte3"],
	overrides: [
		{
			files: ["**/*.svelte"],
			processor: "svelte3/svelte3",
		},
	],
	rules: {
		"no-bitwise": ["error"],
	},
};
