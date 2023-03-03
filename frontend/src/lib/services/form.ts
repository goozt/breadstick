import type { Item } from '$types/menu';

const parseFormData = (data: FormData, key: string): string => {
	return data.has(key) ? data.get(key)?.toString() ?? '' : '';
};

export const getFormData = (formData: FormData): Item => {
	const item: Item = {
		name: parseFormData(formData, 'name'),
		category: parseFormData(formData, 'category'),
		price: parseFloat(parseFormData(formData, 'price'))
	};
	return item;
};
