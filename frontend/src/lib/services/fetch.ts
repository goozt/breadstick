import { env } from '$services/env';

const fetchData = async (url: string) => {
	const res = await fetch(url, {
		mode: 'cors',
		headers: {
			Accept: 'application/json',
			Authorization: 'Bearer ' + env.token
		}
	});
	return await res.json();
};

export const fetchServerData = (path: string) => {
	return fetchData(env.server + path);
};

export const fetchAPIv1 = (path: string) => {
	return fetchData(env.api + path);
};
