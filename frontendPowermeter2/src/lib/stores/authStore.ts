// src/lib/stores/authStore.ts
import { writable } from 'svelte/store';
import { browser } from '$app/environment';
import { jwtDecode } from 'jwt-decode';

// Store untuk menyimpan token JWT
const initialToken = browser ? window.localStorage.getItem('jwt_token') : null;
export const authToken = writable<string | null>(initialToken);
export const isAuthenticated = writable<boolean>(!!initialToken);

let expirationTimer: ReturnType<typeof setTimeout>;

// Fungsi terpusat untuk logout yang bisa dipanggil dari mana saja
export function logout() {
	authToken.set(null);
}

// Langganan ini akan menangani sinkronisasi localStorage DAN memeriksa kedaluwarsa
authToken.subscribe((token) => {
	if (browser) {
		// Hapus timer lama setiap kali token berubah (misalnya saat login/logout)
		if (expirationTimer) {
			clearTimeout(expirationTimer);
		}

		if (token) {
			// Simpan token ke localStorage
			window.localStorage.setItem('jwt_token', token);
			isAuthenticated.set(true);

			try {
				const decoded: { exp: number } = jwtDecode(token);
				const expirationTime = decoded.exp * 1000; // Konversi ke milidetik
				const currentTime = Date.now();
				const timeUntilExpiration = expirationTime - currentTime;

				if (timeUntilExpiration <= 0) {
					console.warn("Token sudah kedaluwarsa saat dimuat. Logout.");
					logout();
				} else {
					// Jadwalkan logout otomatis tepat saat token kedaluwarsa
					expirationTimer = setTimeout(() => {
						console.log("Sesi berakhir. Logout otomatis.");
						logout();
					}, timeUntilExpiration);
				}
			} catch (error) {
				console.error("Gagal mendekode token. Logout.", error);
				logout();
			}
		} else {
			// Jika tidak ada token, hapus dari localStorage
			window.localStorage.removeItem('jwt_token');
			isAuthenticated.set(false);
		}
	}
});
// // src/lib/stores/authStore.ts
// import { writable } from 'svelte/store';
// import { browser } from '$app/environment';

// // Store untuk menyimpan token JWT
// const initialToken = browser ? window.localStorage.getItem('jwt_token') : null;
// export const authToken = writable<string | null>(initialToken);

// authToken.subscribe((value) => {
// 	if (browser) {
// 		if (value) {
// 			window.localStorage.setItem('jwt_token', value);
// 		} else {
// 			window.localStorage.removeItem('jwt_token');
// 		}
// 	}
// });

// // Store untuk status login
// export const isAuthenticated = writable<boolean>(!!initialToken);
// authToken.subscribe(token => isAuthenticated.set(!!token));