import { env } from '$services/env';

const fetchData = async (method: string, url: string, data: unknown) => {
	const res = await fetch(url, {
		method: method,
		mode: 'cors',
		headers: {
			Accept: 'application/json',
			Authorization: 'Bearer ' + env.token,
			'Content-Type': 'application/json'
		},
		body: JSON.stringify(data)
	});
	return await res.json();
};

export const fetchServerData = (method: string, path: string, data: unknown = undefined) => {
	return fetchData(method, env.server + path, data);
};

export const fetchAPIv1 = (method: string, path: string, data: unknown = undefined) => {
	return fetchData(method, env.api + path, data);
};
