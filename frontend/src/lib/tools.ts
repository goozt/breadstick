const server = import.meta.env.VITE_API_SERVER ?? 'https://breadstick-api.goozt.org';

export const api = {
	token: import.meta.env.VITE_API_TOKEN,
	server: server,
	url: server + (import.meta.env.VITE_API_PATH ?? '/api/v1/')
};
