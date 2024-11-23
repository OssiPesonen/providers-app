import type { PageLoad } from './$types';
import { Empty } from '$lib/proto/google/protobuf/empty';
import { apiClient } from '$lib/api/client';

export type Provider = {
    id: number;
    name: string;
    region: string;
    city: string;
    lineOfBusiness: string;
};

export const load: PageLoad = async (): Promise<{providers: Provider[]}> => {
    let providers: Provider[] = [];
    const client = apiClient();

    try {
        const { response } = await client.listProviders(Empty);
        providers = response.providers;
    } catch(error) {
        // Todo: Define the interface for this and return
        console.error(error);
    }

	return {
		providers,
	};
};