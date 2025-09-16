import { c as create_ssr_component, b as add_attribute } from "../../../chunks/ssr.js";
import "../../../chunks/authStore.js";
import "@sveltejs/kit/internal";
import "../../../chunks/exports.js";
import "../../../chunks/utils.js";
import "../../../chunks/state.svelte.js";
const Page = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let username = "";
  let password = "";
  return `<div class="flex items-center justify-center h-screen"><div class="w-full max-w-md p-8 space-y-6 bg-white rounded-lg shadow-md"><h1 class="text-2xl font-bold text-center" data-svelte-h="svelte-16fcape">Login</h1> <form class="space-y-6"><div><label for="username" class="block text-sm font-medium" data-svelte-h="svelte-8yf409">Username</label> <input type="text" id="username" class="w-full px-3 py-2 mt-1 border rounded-md" required${add_attribute("value", username, 0)}></div> <div><label for="password" class="block text-sm font-medium" data-svelte-h="svelte-1ck1rtt">Password</label> <input type="password" id="password" class="w-full px-3 py-2 mt-1 border rounded-md" required${add_attribute("value", password, 0)}></div> ${``} <button type="submit" class="w-full px-4 py-2 font-bold text-white bg-blue-600 rounded-md hover:bg-blue-700" data-svelte-h="svelte-zittyw">Login</button></form></div></div>`;
});
export {
  Page as default
};
