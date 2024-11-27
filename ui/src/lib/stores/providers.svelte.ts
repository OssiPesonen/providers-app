import type { RpcError } from 'grpc-web';
import apiClient from '$lib/api/client';
import { Empty } from '$lib/proto/google/protobuf/empty';
import { Provider } from '$lib/proto/providers_app_service';

let providers = $state<Provider[]>([]);
let fetched = $state(0);
let providerErrors = $state('');

// 5 min
const CACHE_TIME = 5 * 60;

// Retrieve list of providers. This is cached in memory.
// The list will only be refreshed after 5 minutes.
export const listProviders = async (overwrite: boolean = false) => {
	const cacheValid = providers.length > 0 && (Date.now() / 1000) < (fetched + CACHE_TIME);
	if (cacheValid && !overwrite) return providers;

	try {
		const client = apiClient();
		const { response } = await client.listProviders(Empty);
		providers = response.providers;
		// Update fetched timestamp for cache
		fetched = Date.now() / 1000;
		return providers
	} catch (e) {
		const rpcError = e as RpcError;
		providerErrors = rpcError.code.toString();
		return false;
	}
};

export const error = {
	get message() {
		return providerErrors
	}
}