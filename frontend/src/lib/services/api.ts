import { browser } from '$app/environment';
import { env } from '$services/env';
import { QueryClient } from '@sveltestack/svelte-query';

export const queryConfig = {
	enabled: browser && env.token !== undefined
};

export const queryClient = new QueryClient();
