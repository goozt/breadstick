const server: string = import.meta.env.VITE_API_SERVER ?? 'https://breadstick-api.goozt.org';

type Environment = {
	server: string;
	api: string;
};

export const env: Environment = {
	server: server,
	api: server + (import.meta.env.VITE_API_PATH ?? '/api/v1')
};
