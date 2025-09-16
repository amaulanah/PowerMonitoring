import { w as writable } from "./index.js";
import "jwt-decode";
const initialToken = null;
const authToken = writable(initialToken);
const isAuthenticated = writable(false);
authToken.subscribe((token) => {
});
export {
  authToken as a,
  isAuthenticated as i
};
