<!-- src/routes/pm/[meterId]/+page.svelte -->
<script lang="ts">
	import type { PageData } from './$types';
	import { realtimeStore } from '$lib/stores/websocketStore';
	import { authToken } from '$lib/stores/authStore';
	import { onMount } from 'svelte';
	import RealtimeChart from '$lib/components/RealtimeChart.svelte';
	import MetricCard from '$lib/components/MetricCard.svelte';

	export let data: PageData;

	let liveData: any = {};
	let token: string | null;
	authToken.subscribe(value => token = value);

	// Subscribe ke store untuk mendapatkan data realtime
	const unsubscribe = realtimeStore.subscribe((message) => {
		if (message && message.length > 0) {
			const found = message.find((d: any) => d.deviceId === data.meterId);
			if (found) {
				liveData = found;
			}
		}
	});

	onMount(() => {
		if (token) {
			realtimeStore.connect('ws://localhost:1234/ws', token);
		}
		
		return () => {
			unsubscribe();
			realtimeStore.close();
		};
	});
</script>

<div class="space-y-6">
	<h1 class="text-3xl font-bold text-gray-800 dark:text-white">Dashboard: {data.meterId.toUpperCase()}</h1>

	<div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
		<MetricCard title="Active Power Total" value={liveData.activePowerTotal} unit="kW" />
		<MetricCard title="Frequency" value={liveData.frequency} unit="Hz" />
		<MetricCard title="Avg. Voltage (L-N)" value={liveData.voltage1PhaseAverage} unit="V" />
		<MetricCard title="PF Total" value={liveData.powerFactorTotal} unit="" />
	</div>

	<div class="bg-white dark:bg-gray-800 p-4 rounded-lg shadow">
		<RealtimeChart meterId={data.meterId} />
	</div>
</div>
