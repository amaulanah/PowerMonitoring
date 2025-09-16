import { o as onDestroy, c as create_ssr_component, d as spread, b as add_attribute, f as escape_object, v as validate_component, e as each, a as escape } from "../../../../chunks/ssr.js";
import { r as realtimeStore } from "../../../../chunks/websocketStore.js";
import { g as get_store_value } from "../../../../chunks/utils2.js";
import { Chart as Chart$1, LineController, Title, Tooltip, Legend, LineElement, CategoryScale, LinearScale, PointElement, TimeScale } from "chart.js";
import "chartjs-adapter-date-fns";
import { a as authToken } from "../../../../chunks/authStore.js";
const eventPrefix = /^on/;
const events = [];
Object.keys(globalThis).forEach((key) => {
  if (eventPrefix.test(key)) {
    events.push(key.replace(eventPrefix, ""));
  }
});
function useForwardEvents(getRef) {
  const destructors = [];
  onDestroy(() => {
    while (destructors.length) {
      destructors.pop()();
    }
  });
}
function clean(props2) {
  let { data: data2, type: type2, options: options2, plugins: plugins2, children, $$scope, $$slots, ...rest } = props2;
  return rest;
}
const Chart = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let { type } = $$props;
  let { data = { datasets: [] } } = $$props;
  let { options = {} } = $$props;
  let { plugins = [] } = $$props;
  let { updateMode = void 0 } = $$props;
  let { chart = null } = $$props;
  let canvasRef;
  let props = clean($$props);
  onDestroy(() => {
    if (chart) chart.destroy();
    chart = null;
  });
  useForwardEvents();
  if ($$props.type === void 0 && $$bindings.type && type !== void 0) $$bindings.type(type);
  if ($$props.data === void 0 && $$bindings.data && data !== void 0) $$bindings.data(data);
  if ($$props.options === void 0 && $$bindings.options && options !== void 0) $$bindings.options(options);
  if ($$props.plugins === void 0 && $$bindings.plugins && plugins !== void 0) $$bindings.plugins(plugins);
  if ($$props.updateMode === void 0 && $$bindings.updateMode && updateMode !== void 0) $$bindings.updateMode(updateMode);
  if ($$props.chart === void 0 && $$bindings.chart && chart !== void 0) $$bindings.chart(chart);
  return `<canvas${spread([escape_object(props)], {})}${add_attribute("this", canvasRef, 0)}></canvas>`;
});
const Line = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  Chart$1.register(LineController);
  let { chart = null } = $$props;
  let props;
  let baseChartRef;
  useForwardEvents();
  if ($$props.chart === void 0 && $$bindings.chart && chart !== void 0) $$bindings.chart(chart);
  let $$settled;
  let $$rendered;
  let previous_head = $$result.head;
  do {
    $$settled = true;
    $$result.head = previous_head;
    props = $$props;
    $$rendered = `${validate_component(Chart, "Chart").$$render(
      $$result,
      Object.assign({}, { type: "line" }, props, { this: baseChartRef }, { chart }),
      {
        this: ($$value) => {
          baseChartRef = $$value;
          $$settled = false;
        },
        chart: ($$value) => {
          chart = $$value;
          $$settled = false;
        }
      },
      {}
    )}`;
  } while (!$$settled);
  return $$rendered;
});
const RealtimeChart = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  Chart$1.register(Title, Tooltip, Legend, LineElement, CategoryScale, LinearScale, PointElement, TimeScale);
  let { meterId } = $$props;
  let chartData = { labels: [], datasets: [] };
  let selectedParameter = "Active_Energy_Kwh";
  let selectedInterval = "minute";
  const parameters = [
    "Timestamp",
    "DeviceId",
    "Active_Energy_Kwh",
    "Current_A",
    "Current_B",
    "Current_C",
    "Current_N",
    "Current_G",
    "Current_Avg",
    "Voltage_AB",
    "Voltage_BC",
    "Voltage_CA",
    "VoltageL_Avg",
    "Voltage_AN",
    "Voltage_BN",
    "Voltage_CN",
    "NA",
    "VoltageN_Avg",
    "Active_Power_A",
    "Active_Power_B",
    "Active_Power_C",
    "Active_Power_Total",
    "Reactive_Power_A",
    "Reactive_Power_B",
    "Reactive_Power_C",
    "Reactive_Power_Total",
    "Apparent_Power_A",
    "Apparent_Power_B",
    "Apparent_Power_C",
    "Apparent_Power_Total",
    "Power_Factor_A",
    "Power_Factor_B",
    "Power_Factor_C",
    "Power_Factor_Total",
    "Frequency"
  ];
  const intervals = ["second", "minute", "hour", "day"];
  async function fetchData() {
    const token = get_store_value(authToken);
    if (!token) {
      console.error("tidak ada token, fetch chart tidak ditampilkan");
      return;
    }
    try {
      const authHeaders = { "Authorization": `Bearer ${token}` };
      const res = await fetch(`/api/historical-data?deviceId=${meterId}&parameter=${selectedParameter}&interval=${selectedInterval}`, { headers: authHeaders });
      if (!res.ok) {
        throw new Error(`Gagal mengambil data chart: ${res.statusText}`);
      }
      const data = await res.json();
      chartData = {
        datasets: [
          {
            label: `${selectedParameter} (${meterId})`,
            data: data.map((d) => ({ x: new Date(d.time), y: d.value })),
            borderColor: "rgb(75, 192, 192)",
            tension: 0.1,
            pointRadius: 2
          }
        ]
      };
    } catch (error) {
      console.error("Gagal fetch data chart:", error);
      chartData = { labels: [], datasets: [] };
    }
  }
  if ($$props.meterId === void 0 && $$bindings.meterId && meterId !== void 0) $$bindings.meterId(meterId);
  {
    {
      fetchData();
    }
  }
  return `  <div class="space-y-4"><div class="flex flex-wrap items-center gap-4"><div><label for="parameter" class="block text-sm font-medium text-gray-700 dark:text-gray-300" data-svelte-h="svelte-1tbjpvv">Parameter</label> <select id="parameter" class="mt-1 block w-full pl-3 pr-10 py-2 text-base border-gray-300 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm rounded-md dark:bg-gray-700 dark:border-gray-600 dark:text-white">${each(parameters, (param) => {
    return `<option${add_attribute("value", param, 0)}>${escape(param)}</option>`;
  })}</select></div> <div><label for="interval" class="block text-sm font-medium text-gray-700 dark:text-gray-300" data-svelte-h="svelte-17y92vu">Interval Agregasi</label> <select id="interval" class="mt-1 block w-full pl-3 pr-10 py-2 text-base border-gray-300 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm rounded-md dark:bg-gray-700 dark:border-gray-600 dark:text-white">${each(intervals, (interval) => {
    return `<option${add_attribute("value", interval, 0)}>${escape(interval.charAt(0).toUpperCase() + interval.slice(1))}</option>`;
  })}</select></div></div> <div class="h-96 relative">${validate_component(Line, "Line").$$render(
    $$result,
    {
      data: chartData,
      options: {
        responsive: true,
        maintainAspectRatio: false,
        scales: {
          x: {
            type: "time",
            time: {
              unit: "hour"
            }
          }
        }
      }
    },
    {},
    {}
  )}</div></div>`;
});
const MetricCard = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let { title } = $$props;
  let { value } = $$props;
  let { unit } = $$props;
  if ($$props.title === void 0 && $$bindings.title && title !== void 0) $$bindings.title(title);
  if ($$props.value === void 0 && $$bindings.value && value !== void 0) $$bindings.value(value);
  if ($$props.unit === void 0 && $$bindings.unit && unit !== void 0) $$bindings.unit(unit);
  return `  <div class="bg-white dark:bg-gray-800 p-4 rounded-lg shadow text-center"><h3 class="text-sm font-medium text-gray-500 dark:text-gray-400 truncate">${escape(title)}</h3> <p class="mt-1 text-3xl font-semibold text-gray-900 dark:text-white">${escape(value !== void 0 ? value.toFixed(2) : "...")} <span class="text-lg font-medium text-gray-500 dark:text-gray-400">${escape(unit)}</span></p></div>`;
});
const Page = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let { data } = $$props;
  let liveData = {};
  console.log(`Halaman untuk meter ${data.meterId} sedang dieksekusi.`);
  realtimeStore.subscribe((allReadings) => {
    if (allReadings && allReadings.length > 0) {
      console.log("Semua data realtime yang diterima:", allReadings);
      const found = allReadings.find((d) => d.deviceId === data.meterId);
      if (found) {
        console.log(`Data spesifik ditemukan untuk ${data.meterId}:`, found);
        liveData = found;
      }
    }
  });
  if ($$props.data === void 0 && $$bindings.data && data !== void 0) $$bindings.data(data);
  return `<div class="space-y-6"><h1 class="text-3xl font-bold text-gray-800 dark:text-white">Dashboard: ${escape(data.meterId.toUpperCase())}</h1> <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 gap-4">${validate_component(MetricCard, "MetricCard").$$render(
    $$result,
    {
      title: "Active Energy",
      value: liveData.Active_Energy_Kwh,
      unit: "kWh"
    },
    {},
    {}
  )} ${validate_component(MetricCard, "MetricCard").$$render(
    $$result,
    {
      title: "Current R",
      value: liveData.Current_A,
      unit: "A"
    },
    {},
    {}
  )} ${validate_component(MetricCard, "MetricCard").$$render(
    $$result,
    {
      title: "Current S",
      value: liveData.Current_B,
      unit: "A"
    },
    {},
    {}
  )} ${validate_component(MetricCard, "MetricCard").$$render(
    $$result,
    {
      title: "Current T",
      value: liveData.Current_C,
      unit: "A"
    },
    {},
    {}
  )} ${validate_component(MetricCard, "MetricCard").$$render(
    $$result,
    {
      title: "Current N",
      value: liveData.Current_N,
      unit: "A"
    },
    {},
    {}
  )} ${validate_component(MetricCard, "MetricCard").$$render(
    $$result,
    {
      title: "Current G",
      value: liveData.Current_G,
      unit: "A"
    },
    {},
    {}
  )} ${validate_component(MetricCard, "MetricCard").$$render(
    $$result,
    {
      title: "Current Average",
      value: liveData.Current_Avg,
      unit: "A"
    },
    {},
    {}
  )} ${validate_component(MetricCard, "MetricCard").$$render(
    $$result,
    {
      title: "Voltage R",
      value: liveData.Voltage_AN,
      unit: "V"
    },
    {},
    {}
  )} ${validate_component(MetricCard, "MetricCard").$$render(
    $$result,
    {
      title: "Voltage S",
      value: liveData.Voltage_BN,
      unit: "V"
    },
    {},
    {}
  )} ${validate_component(MetricCard, "MetricCard").$$render(
    $$result,
    {
      title: "Voltage T",
      value: liveData.Voltage_CN,
      unit: "V"
    },
    {},
    {}
  )} ${validate_component(MetricCard, "MetricCard").$$render(
    $$result,
    {
      title: "Voltage Average",
      value: liveData.VoltageN_Avg,
      unit: "V"
    },
    {},
    {}
  )} ${validate_component(MetricCard, "MetricCard").$$render(
    $$result,
    {
      title: "Voltage R To N",
      value: liveData.Voltage_AN,
      unit: "V"
    },
    {},
    {}
  )} ${validate_component(MetricCard, "MetricCard").$$render(
    $$result,
    {
      title: "Voltage S To N",
      value: liveData.Voltage_BN,
      unit: "V"
    },
    {},
    {}
  )} ${validate_component(MetricCard, "MetricCard").$$render(
    $$result,
    {
      title: "Voltage T To N",
      value: liveData.Voltage_CN,
      unit: "V"
    },
    {},
    {}
  )} ${validate_component(MetricCard, "MetricCard").$$render(
    $$result,
    {
      title: "N/A",
      value: liveData.NA,
      unit: "V"
    },
    {},
    {}
  )} ${validate_component(MetricCard, "MetricCard").$$render(
    $$result,
    {
      title: "Voltage Average TO Neutral",
      value: liveData.VoltageN_Avg,
      unit: "V"
    },
    {},
    {}
  )} ${validate_component(MetricCard, "MetricCard").$$render(
    $$result,
    {
      title: "Active Power R",
      value: liveData.Active_Power_A,
      unit: "kW"
    },
    {},
    {}
  )} ${validate_component(MetricCard, "MetricCard").$$render(
    $$result,
    {
      title: "Active Power S",
      value: liveData.Active_Power_B,
      unit: "kW"
    },
    {},
    {}
  )} ${validate_component(MetricCard, "MetricCard").$$render(
    $$result,
    {
      title: "Active Power T",
      value: liveData.Active_Power_C,
      unit: "kW"
    },
    {},
    {}
  )} ${validate_component(MetricCard, "MetricCard").$$render(
    $$result,
    {
      title: "Active Power Total",
      value: liveData.Active_Power_Total,
      unit: "kW"
    },
    {},
    {}
  )} ${validate_component(MetricCard, "MetricCard").$$render(
    $$result,
    {
      title: "Reactive Power R",
      value: liveData.Reactive_Power_A,
      unit: "kVAR"
    },
    {},
    {}
  )} ${validate_component(MetricCard, "MetricCard").$$render(
    $$result,
    {
      title: "Reactive Power S",
      value: liveData.Reactive_Power_B,
      unit: "kVAR"
    },
    {},
    {}
  )} ${validate_component(MetricCard, "MetricCard").$$render(
    $$result,
    {
      title: "Reactive Power T",
      value: liveData.Reactive_Power_C,
      unit: "kVAR"
    },
    {},
    {}
  )} ${validate_component(MetricCard, "MetricCard").$$render(
    $$result,
    {
      title: "Reactive Power Total",
      value: liveData.Reactive_Power_Total,
      unit: "kVAR"
    },
    {},
    {}
  )} ${validate_component(MetricCard, "MetricCard").$$render(
    $$result,
    {
      title: "Apparent Power R",
      value: liveData.Apparent_Power_A,
      unit: "VA"
    },
    {},
    {}
  )} ${validate_component(MetricCard, "MetricCard").$$render(
    $$result,
    {
      title: "Apparent Power S",
      value: liveData.Apparent_Power_B,
      unit: "VA"
    },
    {},
    {}
  )} ${validate_component(MetricCard, "MetricCard").$$render(
    $$result,
    {
      title: "Apparent Power T",
      value: liveData.Apparent_Power_C,
      unit: "VA"
    },
    {},
    {}
  )} ${validate_component(MetricCard, "MetricCard").$$render(
    $$result,
    {
      title: "Apparent Power Total",
      value: liveData.Apparent_Power_Total,
      unit: "VA"
    },
    {},
    {}
  )} ${validate_component(MetricCard, "MetricCard").$$render(
    $$result,
    {
      title: "Power Factor R",
      value: liveData.Power_Factor_A,
      unit: ""
    },
    {},
    {}
  )} ${validate_component(MetricCard, "MetricCard").$$render(
    $$result,
    {
      title: "Power Factor S",
      value: liveData.Power_Factor_B,
      unit: ""
    },
    {},
    {}
  )} ${validate_component(MetricCard, "MetricCard").$$render(
    $$result,
    {
      title: "Power Factor T",
      value: liveData.Power_Factor_C,
      unit: ""
    },
    {},
    {}
  )} ${validate_component(MetricCard, "MetricCard").$$render(
    $$result,
    {
      title: "Power Factor Total",
      value: liveData.Power_Factor_Total,
      unit: ""
    },
    {},
    {}
  )} ${validate_component(MetricCard, "MetricCard").$$render(
    $$result,
    {
      title: "Frequency",
      value: liveData.Frequency,
      unit: "Hz"
    },
    {},
    {}
  )}</div> <div class="bg-white dark:bg-gray-800 p-4 rounded-lg shadow">${validate_component(RealtimeChart, "RealtimeChart").$$render($$result, { meterId: data.meterId }, {}, {})}</div></div>`;
});
export {
  Page as default
};
