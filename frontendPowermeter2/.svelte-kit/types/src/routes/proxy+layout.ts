// @ts-nocheck
import type { LayoutLoad } from './$types';
import { browser } from '$app/environment';
import { goto } from '$app/navigation';
import { get } from 'svelte/store';
import { authToken } from '$lib/stores/authStore';

export const load = async ({ url, fetch }: Parameters<LayoutLoad>[0]) => {
    // Pengecekan ini hanya berjalan di sisi browser
    if (browser) {
        const token = get(authToken);
        const path = url.pathname;

        // ATURAN 1: Jika TIDAK ada token & kita BUKAN di halaman login, paksa ke /login.
        if (!token && path !== '/login') {
            await goto('/login');
        }

        // ATURAN 2: Jika ADA token & kita SEDANG di halaman login, paksa ke halaman utama.
        if (token && path === '/login') {
            await goto('/');
        }
    }

    // Bagian ini tetap berjalan untuk mengambil data devices jika ada token.
    // Ini penting agar sidebar tetap terisi.
    const tokenForFetch = get(authToken);
    if (tokenForFetch) {
        try {
            const response = await fetch('/api/devices', {
                headers: { 'Authorization': `Bearer ${tokenForFetch}` }
            });
            if (response.ok) {
                const devices = await response.json();
                return { devices };
            }
        } catch (error) {
            console.error("Gagal fetch devices di layout:", error);
        }
    }

    // Kembalikan data kosong jika tidak ada token atau fetch gagal
    return { devices: [] };
};