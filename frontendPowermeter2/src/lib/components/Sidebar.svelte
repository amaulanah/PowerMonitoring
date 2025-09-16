<!-- src/lib/components/Sidebar.svelte -->
<script lang="ts">
	import { page } from '$app/stores';
	export let devices: string[] = [];
</script>

<aside class="w-64 bg-white dark:bg-gray-800 shadow-md p-4 flex flex-shrink-0 flex-col h-full">
	<div class="p-4 border-b dark:border-gray-700">
		<h2 class="text-xl font-bold text-gray-800 dark:text-white">Power Meters</h2>
	</div>
	<nav class="mt-4 overflow-y-auto">
		<ul>
			{#each devices.sort() as device (device)}
				<li>
					<a
						href="/pm/{device}"
						class="block px-4 py-2 text-gray-600 dark:text-gray-300 hover:bg-gray-200 dark:hover:bg-gray-700 transition-colors duration-200"
						class:bg-gray-200={$page.url.pathname.includes(`/pm/${device}`)}
						class:dark:bg-gray-700={$page.url.pathname.includes(`/pm/${device}`)}
					>
						{device.toUpperCase()}
					</a>
				</li>
			{/each}
		</ul>
	</nav>
</aside>
