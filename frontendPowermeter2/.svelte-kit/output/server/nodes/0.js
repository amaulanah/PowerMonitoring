import * as universal from '../entries/pages/_layout.ts.js';

export const index = 0;
let component_cache;
export const component = async () => component_cache ??= (await import('../entries/pages/_layout.svelte.js')).default;
export { universal };
export const universal_id = "src/routes/+layout.ts";
export const imports = ["_app/immutable/nodes/0.9sW5dmym.js","_app/immutable/chunks/NQmCZVoA.js","_app/immutable/chunks/Euv9bMr_.js","_app/immutable/chunks/C__Wf3JA.js","_app/immutable/chunks/DK1tR62X.js","_app/immutable/chunks/IHki7fMi.js","_app/immutable/chunks/XT8gz2Jn.js","_app/immutable/chunks/DqI0N4dn.js"];
export const stylesheets = ["_app/immutable/assets/0.CGL5fLTN.css"];
export const fonts = [];
