// src/lib/stores/authStore.ts
import { writable } from 'svelte/store';
import { browser } from '$app/environment';

// Store untuk menyimpan token JWT
const initialToken = browser ? window.localStorage.getItem('jwt_token') : null;
export const authToken = writable<string | null>(initialToken);

authToken.subscribe((value) => {
	if (browser) {
		if (value) {
			window.localStorage.setItem('jwt_token', value);
		} else {
			window.localStorage.removeItem('jwt_token');
		}
	}
});

// Store untuk status login
export const isAuthenticated = writable<boolean>(!!initialToken);
authToken.subscribe(token => isAuthenticated.set(!!token));