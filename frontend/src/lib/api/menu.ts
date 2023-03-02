import { useQuery, useMutation } from '@sveltestack/svelte-query';
import { queryConfig } from '$services/api';
import { fetchAPIv1 } from '$services/fetch';
import type { Item, MenuResult, CreateResult } from './menu.d';

const getMenu = async () => {
	return fetchAPIv1('GET', '/menu');
};

const createItem = async (newItem: Item): Promise<CreateResult> => {
	return fetchAPIv1('POST', '/menu', newItem);
};

const parseFormData = (data: FormData, key: string): string => {
	return data.has(key) ? data.get(key)?.toString() ?? '' : '';
};

export default {
	getFormData: (formData: FormData): Item => {
		const item: Item = {
			name: parseFormData(formData, 'name'),
			category: parseFormData(formData, 'category'),
			price: parseFloat(parseFormData(formData, 'price'))
		};
		return item;
	},
	list: () => {
		return useQuery<MenuResult>('menu', getMenu, queryConfig);
	},
	create: () => {
		return useMutation(createItem);
	}
};
