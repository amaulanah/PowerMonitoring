<script lang="ts">
	import { authToken } from '$lib/stores/authStore';
	import { goto } from '$app/navigation';
    
    // 1. IMPORT onMount DAN get
    import { onMount } from 'svelte';
    import { get } from 'svelte/store';

	let username = '';
	let password = '';
	let error = '';

	async function handleLogin() {
		error = '';
		try {
			const response = await fetch('/api/login', {
				method: 'POST',
				headers: { 'Content-Type': 'application/json' },
				body: JSON.stringify({ username, password })
			});
			
			if (!response.ok) {
				throw new Error('Username atau password salah');
			}
			
			const data = await response.json();
			authToken.set(data.token);
			await goto('/');
			
		} catch (err: any) {
			error = err.message;
		}
	}
//	async function handleLogin() {
//		error = '';
//		try {
//			const response = await fetch('/api/login', {
//				method: 'POST',
//				headers: { 'Content-Type': 'application/json' },
//				body: JSON.stringify({ username, password })
//			});
//			
//			if (!response.ok) {
//				throw new Error('Username atau password salah');
//			}
//
//			const data = await response.json();
//			authToken.set(data.token);
//		goto('/');
//			} catch (err: any) {
//				error = err.message;
//			}
//		}


//    // 2. TAMBAHKAN BLOK onMount INI
//    // Blok ini akan berjalan saat halaman login pertama kali dibuka
//    onMount(() => {
//        // Cek apakah sudah ada token di dalam store
//        if (get(authToken)) {
//            console.log("Login page: Sudah ada token, redirect ke halaman utama...");
//            goto('/'); // Jika ada, langsung arahkan ke dashboard
//        }
//    });
//
//	async function handleLogin() {
//		error = '';
//		try {
//			const response = await fetch('/api/login', {
//				method: 'POST',
//				headers: { 'Content-Type': 'application/json' },
//				body: JSON.stringify({ username, password })
//			});
//
//			if (!response.ok) {
//				throw new Error('Username atau password salah');
//			}
//
//			const data = await response.json();
//			const token = data.token;
//			
//          // Baris-baris ini sudah benar
//			window.localStorage.setItem('jwt_token', token);
//			authToken.set(token);
//
//			goto('/'); // Arahkan ke halaman utama setelah login
//		} catch (err: any) {
//			error = err.message;
//		}
//	}
</script>

<div class="flex items-center justify-center h-screen">
	<div class="w-full max-w-md p-8 space-y-6 bg-white rounded-lg shadow-md">
		<h1 class="text-2xl font-bold text-center">Login</h1>
		<form on:submit|preventDefault={handleLogin} class="space-y-6">
			<div>
				<label for="username" class="block text-sm font-medium">Username</label>
				<input type="text" id="username" bind:value={username} class="w-full px-3 py-2 mt-1 border rounded-md" required />
			</div>
			<div>
				<label for="password" class="block text-sm font-medium">Password</label>
				<input type="password" id="password" bind:value={password} class="w-full px-3 py-2 mt-1 border rounded-md" required />
			</div>
			{#if error}
				<p class="text-sm text-red-600">{error}</p>
			{/if}
			<button type="submit" class="w-full px-4 py-2 font-bold text-white bg-blue-600 rounded-md hover:bg-blue-700">
				Login
			</button>
		</form>
	</div>
</div>
