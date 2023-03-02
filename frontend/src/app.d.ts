// See https://kit.svelte.dev/docs/types#app
// for information about these interfaces
declare global {
	namespace App {
		// interface Error {}
		// interface Locals {}
		// interface PageData {}
		// interface Platform {}
		interface ImportMetaEnv {
			VITE_API_TOKEN: string;
			VITE_API_SERVER: string;
			VITE_API_PATH: string;
		}
	}
}

export {};
