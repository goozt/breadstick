import { v4 } from 'uuid';

export const newUUID = (): string => v4();

export const api = {
	token: import.meta.env.VITE_API_TOKEN,
	server: import.meta.env.VITE_API_SERVER,
	url: import.meta.env.VITE_API_SERVER + (import.meta.env.VITE_API_PATH ?? '/api/v1/')
};
