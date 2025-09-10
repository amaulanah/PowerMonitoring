import { writable } from 'svelte/store';
import type { PowerMeterReading } from '$lib/types/PowerMeterReading';

function createRealtimeStore() {
	const { subscribe, set } = writable<PowerMeterReading[]>([]);
	let ws: WebSocket | null = null;

	function connect(url: string, token: string) {
		if (ws) {
			ws.close();
		}

		ws = new WebSocket(`${url}?token=${token}`);

		ws.onopen = () => console.log('WebSocket connection opened');
		ws.onmessage = (event) => {
			try {
				set(JSON.parse(event.data));
			} catch (error) {
				console.error('Error parsing WebSocket message:', error);
			}
		};
		ws.onclose = () => {
			console.log('WebSocket connection closed');
			ws = null;
		};
		ws.onerror = (error) => console.error('WebSocket error:', error);
	}

	function close() {
		if (ws) ws.close();
	}

	return { subscribe, connect, close };
}

export const realtimeStore = createRealtimeStore();