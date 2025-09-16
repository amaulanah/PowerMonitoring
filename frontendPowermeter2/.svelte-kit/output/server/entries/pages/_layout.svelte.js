import { s as subscribe } from "../../chunks/utils2.js";
import { c as create_ssr_component, e as each, a as escape, v as validate_component } from "../../chunks/ssr.js";
import { p as page } from "../../chunks/stores.js";
import { i as isAuthenticated } from "../../chunks/authStore.js";
import "../../chunks/websocketStore.js";
import "@sveltejs/kit/internal";
import "../../chunks/exports.js";
import "../../chunks/utils.js";
import "../../chunks/state.svelte.js";
const Sidebar = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let $page, $$unsubscribe_page;
  $$unsubscribe_page = subscribe(page, (value) => $page = value);
  let { devices = [] } = $$props;
  if ($$props.devices === void 0 && $$bindings.devices && devices !== void 0) $$bindings.devices(devices);
  $$unsubscribe_page();
  return `  <aside class="w-64 bg-white dark:bg-gray-800 shadow-md p-4 flex flex-shrink-0 flex-col h-full"><div class="p-4 border-b dark:border-gray-700" data-svelte-h="svelte-w687sj"><h2 class="text-xl font-bold text-gray-800 dark:text-white">Power Meters</h2></div> <nav class="mt-4 overflow-y-auto"><ul>${each(devices.sort(), (device) => {
    return `<li><a href="${"/pm/" + escape(device, true)}" class="${[
      "block px-4 py-2 text-gray-600 dark:text-gray-300 hover:bg-gray-200 dark:hover:bg-gray-700 transition-colors duration-200",
      ($page.url.pathname.includes(`/pm/${device}`) ? "bg-gray-200" : "") + " " + ($page.url.pathname.includes(`/pm/${device}`) ? "dark:bg-gray-700" : "")
    ].join(" ").trim()}">${escape(device.toUpperCase())}</a> </li>`;
  })}</ul></nav></aside>`;
});
const Layout = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let $isAuthenticated, $$unsubscribe_isAuthenticated;
  $$unsubscribe_isAuthenticated = subscribe(isAuthenticated, (value) => $isAuthenticated = value);
  let { data } = $$props;
  if ($$props.data === void 0 && $$bindings.data && data !== void 0) $$bindings.data(data);
  $$unsubscribe_isAuthenticated();
  return `<div class="flex h-screen"> ${$isAuthenticated ? `${validate_component(Sidebar, "Sidebar").$$render($$result, { devices: data.devices }, {}, {})}` : ``} <main class="flex-1 p-4 md:p-6 lg:p-8 overflow-y-auto">${slots.default ? slots.default({}) : ``}</main></div>`;
});
export {
  Layout as default
};
