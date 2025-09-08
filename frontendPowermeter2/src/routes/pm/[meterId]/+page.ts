// src/routes/pm/[meterId]/+page.ts
import type { PageLoad } from './$types';

export const load: PageLoad = ({ params }) => {
	return {
		meterId: params.meterId
	};
};