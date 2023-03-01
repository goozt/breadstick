import { api } from '$lib/tools';

export const fetchHealth = async () => {
	return fetch(api.server + '/health', {
		mode: 'cors',
		headers: {
			Accept: 'application/json',
			Authorization: 'Bearer ' + api.token,
			'Access-Control-Allow-Origin': '*'
		}
	})
		.then((res) => res.json())
		.catch(() => {
			return { status: 'error' };
		});
};
