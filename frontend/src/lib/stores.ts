import { writable } from 'svelte/store';
import { items } from '$data/MenuItems';

export const pageTracker = writable('home');
export const itemCollection = writable(items);
