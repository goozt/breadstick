import { useQuery, useMutation } from '@sveltestack/svelte-query';
import { queryConfig } from '$services/api';
import { fetchAPIv1 } from '$services/fetch';
import type {
	Item,
	MenuResult,
	CreateResult,
	ItemResult,
	EditableItem,
	DeleteResult,
	DeleteMenuResult
} from '$types/menu';

//  Get menu items

export const getMenu = () => {
	return useQuery<MenuResult>('menu', () => fetchAPIv1('GET', '/menu'), queryConfig);
};

// Create new menu items

export const createItem = () => {
	return useMutation(
		async (newItem: Item): Promise<CreateResult> => fetchAPIv1('POST', '/menu', newItem)
	);
};

// Get menu item with id

export const getItem = (id: string) => {
	return useQuery(
		`menu/${id}`,
		(): Promise<ItemResult> => fetchAPIv1('GET', `/menu/${id}`),
		queryConfig
	);
};

// Update menu item with id

export const updateItem = (id: string, newItem: EditableItem) => {
	return useMutation((): Promise<ItemResult> => fetchAPIv1('PUT', `/menu/${id}`, newItem));
};

// Delete menu item with id

export const deleteItem = (id: string) => {
	return useMutation((): Promise<DeleteResult> => fetchAPIv1('DELETE', `/menu/${id}`));
};

// Delete all menu items

export const deleteMenu = () => {
	return useMutation((): Promise<DeleteMenuResult> => fetchAPIv1('DELETE', '/menu'));
};
