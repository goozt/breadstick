import preprocess from 'svelte-preprocess';
import adapter from '@sveltejs/adapter-auto';
import { vitePreprocess } from '@sveltejs/kit/vite';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	// Consult https://kit.svelte.dev/docs/integrations#preprocessors
	// for more information about preprocessors
	preprocess: [
		vitePreprocess(),
		preprocess({
			postcss: true
		})
	],

	kit: {
		alias: {
			$store: './src/lib/stores.js',
			$stores: './src/lib/stores/*',
			$data: './src/lib/data',
			'$data/*': './src/lib/data/*',
			$ui: './src/lib/components',
			'$ui/*': './src/lib/components/*',
			'$ui-types': './src/lib/components/types',
			'$ui-types/*': './src/lib/components/types/*',
			$icons: './src/lib/icons',
			'$icons/*': './src/lib/icons/*'
		},
		// adapter-auto only supports some environments, see https://kit.svelte.dev/docs/adapter-auto for a list.
		// If your environment is not supported or you settled on a specific environment, switch out the adapter.
		// See https://kit.svelte.dev/docs/adapters for more information about adapters.
		adapter: adapter()
	}
};

export default config;
