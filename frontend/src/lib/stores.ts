import { writable } from 'svelte/store';
import { items } from '$data/MenuItems';
import type { NavMenu, MenuItem } from '$ui-types';

export const pageTracker = writable({ name: 'Home', url: '/' } as NavMenu);
export const itemCollection = writable(items as MenuItem[]);
