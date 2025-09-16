// @ts-nocheck
// src/routes/pm/[meterId]/+page.ts
import type { PageLoad } from './$types';

export const load = ({ params }: Parameters<PageLoad>[0]) => {
	return {
		meterId: params.meterId
	};
};