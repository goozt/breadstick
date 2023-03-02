import { writable } from 'svelte/store';
import type { NavMenu } from '$types';

export const pageTracker = writable({ name: 'Home', url: '/' } as NavMenu);
