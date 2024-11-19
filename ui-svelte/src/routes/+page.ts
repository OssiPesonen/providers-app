import type { PageLoad } from './$types';
import { Empty } from '$lib/proto/google/protobuf/empty';
import { TrafficLightsServiceClient  } from '$lib/proto/traffic_lights_service.client';
import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport';

export type Provider = {
    id: number;
    name: string;
    region: string;
    city: string;
    lineOfBusiness: string;
};

export const load: PageLoad = async (): Promise<{providers: Provider[]}> => {
    let providers: Provider[] = [];
    const transport = new GrpcWebFetchTransport({
        baseUrl: "http://localhost:8080"
    });
    const client = new TrafficLightsServiceClient(transport);
    
    try {
        const { response } = await client.listProviders(Empty);
        providers = response.providers;
    } catch(err) {
        console.error(err)
    }

	return {
		providers,
	};
};