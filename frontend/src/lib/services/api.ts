import { browser } from '$app/environment';
import { QueryClient } from '@sveltestack/svelte-query';

export const queryConfig = {
	enabled: browser
};

export const queryClient = new QueryClient();
