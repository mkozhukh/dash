import resolve from "@rollup/plugin-node-resolve";
import commonjs from "@rollup/plugin-commonjs";

import svelte from "rollup-plugin-svelte";
import { terser } from "rollup-plugin-terser";
import livereload from "rollup-plugin-livereload";
import visualizer from "rollup-plugin-visualizer";

import pkg from "./package.json";
const name = pkg.productTag;
const watch = !!process.env.ROLLUP_WATCH;

const onwarn = (warning, onwarn) => {
	return (
		warning.code === "a11y-no-onchange" ||
		warning.code === "a11y-autofocus" ||
		warning.code === "a11y-missing-attribute" ||
		onwarn(warning)
	);
};

function config(mode, start, format, production, full, limited) {
	return {
		input: start,
		watch: {
			clearScreen: false,
			include: ["src/**/*", "demos/**/*", "../wx/**/*", "../helpers/**/*"],
		},
		output: {
			sourcemap: !limited,
			name: "app",
			format: format,
			file: `public/${mode}/${name}.js`,
		},
		plugins: [
			!production && livereload({ port: 35803 }),
			resolve(),
			svelte({
				// enable run-time checks when not in production
				dev: !production,
				// we'll extract any component CSS out into
				// a separate file — better for performance
				css: css => {
					css.write(`${name}.css`, !limited);
				},
				onwarn,
			}),

			// If you have external dependencies installed from
			// npm, you'll most likely need these plugins. In
			// some cases you'll need additional configuration —
			// consult the documentation for details:
			// https://github.com/rollup/rollup-plugin-commonjs
			commonjs(),

			// If we're building for production (npm run build
			// instead of npm run dev), minify
			production && terser(),
			production &&
				!limited &&
				visualizer({ filename: "public/stats.html", sourcemap: true }),
		],
	};
}

export default function (cli) {
	let out = [];
	const production = !!cli["config-production"];
	const limited = !!cli["config-limited"];
	if (!!cli["config-demos"])
		out.push(
			config("demos", "demos/index.js", "iife", production, true, limited)
		);
	else out.push(config("dist", "src/index.js", "iife", production, false, limited));

	return out;
}
