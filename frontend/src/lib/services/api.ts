import { browser } from '$app/environment';
import { env } from '$services/env';

export const queryConfig = {
	enabled: browser && env.token !== undefined
};
