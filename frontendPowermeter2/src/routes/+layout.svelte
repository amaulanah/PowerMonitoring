<script lang="ts">
	import '../app.css';
	import Sidebar from '$lib/components/Sidebar.svelte';
	import type { LayoutData } from './$types';

	export let data: LayoutData;
</script>

<div class="flex h-screen">
	<Sidebar devices={data.devices} />
	<main class="flex-1 p-4 md:p-6 lg:p-8 overflow-y-auto">
		<slot />
	</main>
</div>