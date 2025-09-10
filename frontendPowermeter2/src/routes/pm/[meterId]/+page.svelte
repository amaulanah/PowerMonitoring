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
		<MetricCard title="Energy" value={liveData.energyKWh} unit="KWh" />
		<MetricCard title="Active Power Total" value={liveData.activePowerTotal} unit="kW" />
		<MetricCard title="Frequency" value={liveData.frequency} unit="Hz" />
		<MetricCard title="PF Total" value={liveData.powerFactorTotal} unit="" />
		<MetricCard title="Current Average" value={liveData.currentAverage} unit="A" />
		<MetricCard title="Current L1" value={liveData.currentL1} unit="A" />
		<MetricCard title="Current L2" value={liveData.currentL2} unit="A" />
		<MetricCard title="Current L3" value={liveData.currentL3} unit="A" />
		<MetricCard title="Voltage L1-L2" value={liveData.voltageL1ToL2} unit="V" />
		<MetricCard title="Voltage L1-L3" value={liveData.voltageL1ToL3} unit="V" />
		<MetricCard title="Voltage L2-L3" value={liveData.voltageL2ToL3} unit="V" />
		<MetricCard title="Avg. Voltage (L-L)" value={liveData.voltage3PhaseAverage} unit="V" />
		<MetricCard title="Voltage L1-N" value={liveData.voltageL1ToN} unit="V" />
		<MetricCard title="Voltage L2-N" value={liveData.voltageL2ToN} unit="V" />
		<MetricCard title="Voltage L3-N" value={liveData.voltageL3ToN} unit="V" />
		<MetricCard title="Avg. Voltage (L-N)" value={liveData.voltage1PhaseAverage} unit="V" />
		<MetricCard title="Active Power L1" value={liveData.activePowerL1} unit="kW" />
		<MetricCard title="Active Power L2" value={liveData.activePowerL2} unit="kW" />
		<MetricCard title="Active Power L3" value={liveData.activePowerL3} unit="kW" />
		<MetricCard title="Reactive Power L1" value={liveData.reactivePowerL1} unit="kVAR" />
		<MetricCard title="Reactive Power L2" value={liveData.reactivePowerL2} unit="kVAR" />
		<MetricCard title="Reactive Power L3" value={liveData.reactivePowerL3} unit="kVAR" />
		<MetricCard title="Reactive Power Total" value={liveData.reactivePowerTotal} unit="kVAR" />
		<MetricCard title="Power Factor L1" value={liveData.powerFactorL1} unit="" />
		<MetricCard title="Power Factor L2" value={liveData.powerFactorL2} unit="" />
		<MetricCard title="Power Factor L3" value={liveData.powerFactorL3} unit="" />
		<MetricCard title="HD Current" value={liveData.harmonicDistortionCurrent} unit="%" />
		<MetricCard title="HD Voltage 3Ph" value={liveData.harmonicDistortionVoltage3Ph} unit="%" />
		<MetricCard title="HD Voltage 1Ph" value={liveData.harmonicDistortionVoltage1Ph} unit="%" />
	</div>

	<div class="bg-white dark:bg-gray-800 p-4 rounded-lg shadow">
		<RealtimeChart meterId={data.meterId} />
	</div>
</div>