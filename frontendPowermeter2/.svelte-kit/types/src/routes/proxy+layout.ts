// @ts-nocheck
// src/routes/+layout.ts
import type { LayoutLoad } from './$types';

export const load = async ({ fetch }: Parameters<LayoutLoad>[0]) => {
	try {
		const response = await fetch('http://localhost:1234/api/devices');
		if (!response.ok) {
			return { devices: [] };
		}
		const devices: string[] = await response.json();
		return { devices };
	} catch (error) {
		console.error('Gagal mengambil daftar perangkat:', error);
		return { devices: [] };
	}
};