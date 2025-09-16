<script lang="ts">
	import '../app.css';
	import Sidebar from '$lib/components/Sidebar.svelte';
	import { authToken, isAuthenticated } from '$lib/stores/authStore';
	import { realtimeStore } from '$lib/stores/websocketStore';
	import type { LayoutData } from './$types';
	import { onMount } from 'svelte';
	import { page } from '$app/stores'; // <-- 1. Import 'page' store
	import { goto } from '$app/navigation'; // <-- 2. Import 'goto' untuk redirect
//	import { authToken } from '$lib/stores/authStore';
	import { invalidateAll } from '$app/navigation';
	
	export let data: LayoutData;
	
	onMount(() => {
		const unsubscribe = authToken.subscribe(token => {
			if(token) {
				const wsProtocol = location.protocol === 'https:' ? 'wss:' : 'ws:';
				const wsUrl = `${wsProtocol}//${location.host}/ws`;
				realtimeStore.connect(wsUrl,token);
			}else{
				realtimeStore.close();
			}
		});
		return () => unsubscribe();
	});

//	export let data: LayoutData;
	
//	onMount(() => {
//		console.log("layout.svelte mounted");
//		const unsubscribeAuth = authToken.subscribe(token => {
//			if (!token && $page.url.pathname !== '/login') {
//				console.log("Layout: Token tidak ada, redirect ke /login");
//				goto('/login');
//				return;
//			}
//			if (token && $page.url.pathname === '/login') {
//				console.log("Layout: mengarahkan ke /");
//				goto('/');
//			}
//			if (token) {
//				invalidateAll();
//				const wsProtocol = location.protocol === 'https:' ? 'wss:' : 'ws:';
//				const wsUrl = `${wsProtocol}//${location.host}/ws`;
//				realtimeStore.connect(wsUrl, token);
//			} else {
//				realtimeStore.close();
//			}
//		});
//		return () => {
//			unsubscribeAuth();
//		};
//	});
//		const unsubscribeAuth = authToken.subscribe(token => {
//			console.log("status token berubah:", token ? "token available" : "no token");
//			if (token) {
//				invalidateAll();
//				const wsProtocol = location.protocol === 'https:' ? 'wss:' : 'ws:';
//				const wsUrl = `${wsProtocol}//${location.host}/ws`;
//				console.log(`Mencoba menghubungkan Websocker ke ${wsUrl}`);
//				realtimeStore.connect(wsUrl, token);
//			} else {
//				console.log("Token not available, closing websocket");
//				realtimeStore.close();
//			}
//		});
//		return () => {
//			unsubscribeAuth();
//		};
//		const unsubscribeAuth = authToken.subscribe(token => {
//			// =======================================================
//			// BLOK LOGIKA BARU UNTUK PROTEKSI HALAMAN
//			// =======================================================
//			console.log("Status token berubah:", token ? "Token Available" : "No Token")
//			if (!token && $page.url.pathname !== '/login') {
//				// Jika tidak ada token DAN pengguna tidak sedang di halaman login,
//				// paksa redirect ke halaman login.
//				console.log("Layout: Tidak ada token, mengarahkan ke halaman login...");
//				goto('/login');
//			} 
//			// =======================================================
//			else if (token) {
//				console.log("Layout: Token ditemukan, memulai koneksi WebSocket...");
//				//realtimeStore.connect('ws://localhost:1234/ws', token);
//				const wsProtocol = location.protocol === 'https:' ? 'wss:' : 'ws:'
//				const wsUrl = `${wsProtocol}//${location.host}/ws`;
//				console.log(`Menghubungkan ke Websocket di: ${wsUrl}`);
//			} else {
//				console.log("Layout: Tidak ada token, koneksi WebSocket ditutup.");
//				realtimeStore.close();
//			}
//		});
//
//		return () => {
//			unsubscribeAuth();
//		};
//	});
</script>

<div class="flex h-screen">
	<!-- Tampilkan sidebar hanya jika pengguna sudah login -->
	{#if $isAuthenticated}
		<Sidebar devices={data.devices} />
	{/if}

	<main class="flex-1 p-4 md:p-6 lg:p-8 overflow-y-auto">
		<slot />
	</main>
</div>
