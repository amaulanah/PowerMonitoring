<!-- src/lib/components/MetricCard.svelte -->
<script lang="ts">
	export let title: string;
	export let value: number | undefined;
	export let unit: string;
</script>

<div class="bg-white dark:bg-gray-800 p-4 rounded-lg shadow text-center">
	<h3 class="text-sm font-medium text-gray-500 dark:text-gray-400 truncate">{title}</h3>
	<p class="mt-1 text-3xl font-semibold text-gray-900 dark:text-white">
		{value !== undefined ? value.toFixed(2) : '...'}
		<span class="text-lg font-medium text-gray-500 dark:text-gray-400">{unit}</span>
	</p>
</div>