<!-- src/lib/components/RealtimeChart.svelte -->
<script lang="ts">
	import { Line } from 'svelte-chartjs';
	import { Chart, Title, Tooltip, Legend, LineElement, CategoryScale, LinearScale, PointElement, TimeScale } from 'chart.js';
	import 'chartjs-adapter-date-fns';
	import { onMount } from 'svelte';

	Chart.register(Title, Tooltip, Legend, LineElement, CategoryScale, LinearScale, PointElement, TimeScale);

	export let meterId: string;

	let chartData: any = { labels: [], datasets: [] };
	let selectedParameter = 'ActivePowerTotal';
	let selectedInterval = 'minute';

	const parameters = [
		"EnergyKWh", "ActivePowerTotal", "Frequency", "CurrentAverage", "Voltage3PhaseAverage", "PowerFactorTotal"
	];
	const intervals = ['second', 'minute', 'hour', 'day'];

	async function fetchData() {
		try {
			const res = await fetch(`http://localhost:1234/api/historical-data?deviceId=${meterId}&parameter=${selectedParameter}&interval=${selectedInterval}`);
			const data: { time: string; value: number }[] = await res.json();
			
			chartData = {
				datasets: [{
					label: `${selectedParameter} (${meterId})`,
					data: data.map(d => ({ x: new Date(d.time), y: d.value })),
					borderColor: 'rgb(75, 192, 192)',
					tension: 0.1,
					pointRadius: 2
				}]
			};
		} catch (error) {
			console.error("Gagal fetch data chart:", error);
		}
	}

	onMount(fetchData);

	$: { selectedParameter, selectedInterval; fetchData(); }
</script>

<div class="space-y-4">
	<div class="flex flex-wrap items-center gap-4">
		<div>
			<label for="parameter" class="block text-sm font-medium text-gray-700 dark:text-gray-300">Parameter</label>
			<select id="parameter" bind:value={selectedParameter} class="mt-1 block w-full pl-3 pr-10 py-2 text-base border-gray-300 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm rounded-md dark:bg-gray-700 dark:border-gray-600 dark:text-white">
				{#each parameters as param}
					<option value={param}>{param}</option>
				{/each}
			</select>
		</div>
		<div>
			<label for="interval" class="block text-sm font-medium text-gray-700 dark:text-gray-300">Interval Agregasi</label>
			<select id="interval" bind:value={selectedInterval} class="mt-1 block w-full pl-3 pr-10 py-2 text-base border-gray-300 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm rounded-md dark:bg-gray-700 dark:border-gray-600 dark:text-white">
				{#each intervals as interval}
					<option value={interval}>{interval.charAt(0).toUpperCase() + interval.slice(1)}</option>
				{/each}
			</select>
		</div>
	</div>
	<div class="h-96 relative">
		<Line data={chartData} options={{ 
			responsive: true, 
			maintainAspectRatio: false,
			scales: {
				x: {
					type: 'time',
					time: {
						unit: selectedInterval === 'day' ? 'day' : 'hour'
					}
				}
			}
		}} />
	</div>
</div>