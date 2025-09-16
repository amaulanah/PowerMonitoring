import { g as get_store_value } from "../../chunks/utils2.js";
import "@sveltejs/kit/internal";
import "../../chunks/exports.js";
import "../../chunks/utils.js";
import "../../chunks/state.svelte.js";
import { a as authToken } from "../../chunks/authStore.js";
const load = async ({ url, fetch }) => {
  const tokenForFetch = get_store_value(authToken);
  if (tokenForFetch) {
    try {
      const response = await fetch("/api/devices", {
        headers: { "Authorization": `Bearer ${tokenForFetch}` }
      });
      if (response.ok) {
        const devices = await response.json();
        return { devices };
      }
    } catch (error) {
      console.error("Gagal fetch devices di layout:", error);
    }
  }
  return { devices: [] };
};
export {
  load
};
