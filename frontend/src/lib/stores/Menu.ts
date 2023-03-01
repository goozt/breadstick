import { newUUID } from '$services/uuid';
import { writable } from 'svelte/store';
import { defaultItems } from '$data/MenuItems';
import type { MenuItem } from '$ui-types';
import { newToast } from '$stores/Toast';

export const itemCollection = writable(defaultItems as MenuItem[]);

const generateItem = (): MenuItem => {
	const length = defaultItems.length;
	const number = Math.floor(Math.random() * length);
	const randItem = defaultItems[number];
	return {
		id: newUUID(),
		name: 'Item ' + (number + 1),
		price: Math.floor(Math.random() * 1000),
		imageurl: randItem.imageurl
	};
};

export const addMenuItem = () => {
	itemCollection.update((collection) => {
		const item = generateItem();
		newToast('green', 'Added new item to menu.');
		return [...collection, item];
	});
};

export const removeMenuItem = (event: Event) => {
	const e = event.target as Element;
	const id = e.getAttribute('data-id');
	itemCollection.update((collection) => {
		if (id) {
			const item = collection.find((item) => item.id === id);
			if (item) {
				const items = collection;
				const index = items.indexOf(item);
				if (index !== -1) {
					items.splice(index, 1);
				}
				newToast('green', 'Removed ' + item.name + ' from menu');
				return items;
			}
			newToast('red', 'Item not found');
			return collection;
		}
		newToast('red', 'Internal error: unable to get item id');
		return collection;
	});
};
