<!-- src/routes/+layout.svelte -->
<script lang="ts">
	import '../app.css';
	import Sidebar from '$lib/components/Sidebar.svelte';
	import { authToken, isAuthenticated } from '$lib/stores/authStore';
	import { realtimeStore } from '$lib/stores/websocketStore';
	import type { LayoutData } from './$types';
	import { onMount } from 'svelte';
	import { page } from '$app/stores'; // <-- 1. Import 'page' store
	import { goto } from '$app/navigation'; // <-- 2. Import 'goto' untuk redirect

	export let data: LayoutData;
	
	onMount(() => {
		const unsubscribeAuth = authToken.subscribe(token => {
			// =======================================================
			// BLOK LOGIKA BARU UNTUK PROTEKSI HALAMAN
			// =======================================================
			if (!token && $page.url.pathname !== '/login') {
				// Jika tidak ada token DAN pengguna tidak sedang di halaman login,
				// paksa redirect ke halaman login.
				console.log("Layout: Tidak ada token, mengarahkan ke halaman login...");
				goto('/login');
			} 
			// =======================================================
			else if (token) {
				console.log("Layout: Token ditemukan, memulai koneksi WebSocket...");
				realtimeStore.connect('ws://localhost:1234/ws', token);
			} else {
				console.log("Layout: Tidak ada token, koneksi WebSocket ditutup.");
				realtimeStore.close();
			}
		});

		return () => {
			unsubscribeAuth();
		};
	});
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