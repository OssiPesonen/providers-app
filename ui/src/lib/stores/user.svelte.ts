import type { RpcError } from 'grpc-web';
import apiClient from '$lib/api/client';
import { getAccessToken } from './auth.svelte';
import { Empty } from '$lib/proto/google/protobuf/empty';
import type { UserInfo } from '$lib/proto/providers_app_service';

let userInfoError = $state('');
let userInfo = $state<UserInfo | undefined>(undefined);

// Reset state
export const clearUserInfo = () => {
	userInfo = undefined;
}

// Fetch user info. Set 'overwrite' to 'true' to force refresh
export const getUserInfo = async (overwrite: boolean = false) => {
	if (userInfo !== undefined && !overwrite) return;

	try {
		const client = apiClient();
		const token = getAccessToken();
		if (!token) {
			return;
		}

		const { response: r } = await client.getUserInfo(Empty, {
			meta: {
				"Authorization": `Bearer ${token}`
			}
		});

		userInfo = r;
	} catch (e) {
		const rpcError = e as RpcError;
		userInfoError = rpcError.code.toString();
		return false;
	}
};

export const error = {
	get message() {
		return userInfoError
	}
}

export const user = {
	get info() {
		return userInfo
	}
}