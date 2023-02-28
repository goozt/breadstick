import { newUUID } from '$lib/tools';
import { writable } from 'svelte/store';
import type { ToastItem, ToastItemType } from '$ui-types';
export const toastChannel = writable([] as ToastItem[]);
const toastTimeout = 4000;

export const newToast = (toastType: ToastItemType, toastMessage: string) => {
	toastChannel.update((toasts) => {
		const toast = {
			id: newUUID(),
			type: toastType,
			message: toastMessage
		};
		setTimeout(() => {
			toastChannel.update((toasts) => {
				const items = toasts;
				const index = items.indexOf(toast);
				if (index !== -1) {
					items.splice(index, 1);
				}
				return items;
			});
		}, toastTimeout);
		return [...toasts, toast];
	});
};
