import { writable } from 'svelte/store';
import { defaultItems } from '$data/MenuItems';
import type { NavMenu, MenuItem, ToastItem } from '$ui-types';

export const pageTracker = writable({ name: 'Home', url: '/' } as NavMenu);
export const itemCollection = writable(defaultItems as MenuItem[]);
export const toastChannel = writable([] as ToastItem[]);
