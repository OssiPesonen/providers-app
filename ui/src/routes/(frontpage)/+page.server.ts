import type { PageServerLoad } from './$types';
import { listProviders } from '$lib/stores/providers.svelte';

export const load: PageServerLoad = async () => {
    const providers = await listProviders();

	return {
		providers: JSON.parse(JSON.stringify(providers)),
	};
};