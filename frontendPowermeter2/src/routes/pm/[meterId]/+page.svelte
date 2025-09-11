<!-- src/routes/pm/[meterId]/+page.svelte -->
<script lang="ts">
	import type { PageData } from './$types';
	import { realtimeStore } from '$lib/stores/websocketStore';
	import RealtimeChart from '$lib/components/RealtimeChart.svelte';
	import MetricCard from '$lib/components/MetricCard.svelte';

	export let data: PageData;

	let liveData: any = {};

	// Subscribe ke store untuk mendapatkan data realtime.
	// Tidak perlu lagi mengelola koneksi di sini.
	realtimeStore.subscribe((message) => {
		if (message && message.length > 0) {
			const found = message.find((d: any) => d.deviceId === data.meterId);
			if (found) {
				liveData = found;
			}
		}
	});
</script>

<div class="space-y-6">
	<h1 class="text-3xl font-bold text-gray-800 dark:text-white">
		Dashboard: {data.meterId.toUpperCase()}
	</h1>

	<!-- Bagian HTML ini tidak perlu diubah -->
	<div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 gap-4">
		<MetricCard title="Active Energy" value={liveData.Active_Energy_Kwh} unit="KWh" />
		<MetricCard title="Current R" value={liveData.Current_A} unit="kW" />
		<MetricCard title="Current S" value={liveData.Current_B} unit="Hz" />
		<MetricCard title="Current T" value={liveData.Current_C} unit="" />
		<MetricCard title="Current N" value={liveData.Current_N} unit="A" />
		<MetricCard title="Current G" value={liveData.Current_G} unit="A" />
		<MetricCard title="Current Average" value={liveData.Current_Avg} unit="A" />
		<MetricCard title="Voltage R" value={liveData.Voltage_AB} unit="A" />
		<MetricCard title="Voltage S" value={liveData.Voltage_BC} unit="V" />
		<MetricCard title="Voltage T" value={liveData.Voltage_CA} unit="V" />
		<MetricCard title="Voltage Average" value={liveData.VoltageL_Avg} unit="V" />
		<MetricCard title="Voltage R To N" value={liveData.Voltage_AN} unit="V" />
		<MetricCard title="Voltage S To N" value={liveData.Voltage_BN} unit="V" />
		<MetricCard title="Voltage T To N" value={liveData.Voltage_CN} unit="V" />
		<MetricCard title="N/A" value={liveData.NA} unit="V" />
		<MetricCard title="Voltage Average TO Neutral" value={liveData.VoltageN_Avg} unit="V" />
		<MetricCard title="Active Power R" value={liveData.Active_Power_A} unit="kW" />
		<MetricCard title="Active Power S" value={liveData.Active_Power_B} unit="kW" />
		<MetricCard title="Active Power T" value={liveData.Active_Power_C} unit="kW" />
		<MetricCard title="Active Power Total" value={liveData.Active_Power_Total} unit="kVAR" />
		<MetricCard title="Reactive Power R" value={liveData.Reactive_Power_A} unit="kVAR" />
		<MetricCard title="Reactive Power S" value={liveData.Reactive_Power_B} unit="kVAR" />
		<MetricCard title="Reactive Power T" value={liveData.Reactive_Power_C} unit="kVAR" />
		<MetricCard title="Reactive Power Total" value={liveData.Reactive_Power_Total} unit="" />
		<MetricCard title="Apparent Power R" value={liveData.Apparent_Power_A} unit="" />
		<MetricCard title="Apparent Power S" value={liveData.Apparent_Power_B} unit="" />
		<MetricCard title="Apparent Power T" value={liveData.Apparent_Power_C} unit="%" />
		<MetricCard title="Apparent Power Total" value={liveData.Apparent_Power_Total} unit="%" />
		<MetricCard title="Power Factor R" value={liveData.Power_Factor_A} unit="%" />
		<MetricCard title="Power Factor S" value={liveData.Power_Factor_B} unit="%" />
		<MetricCard title="Power Factor T" value={liveData.Power_Factor_C} unit="%" />
		<MetricCard title="Power Factor Total" value={liveData.Power_Factor_Total} unit="%" />
		<MetricCard title="Frequency" value={liveData.Frequency} unit="Hz" />
	</div>

	<div class="bg-white dark:bg-gray-800 p-4 rounded-lg shadow">
		<RealtimeChart meterId={data.meterId} />
	</div>
</div>