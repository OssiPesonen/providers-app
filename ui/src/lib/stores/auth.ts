import { readonly, writable } from 'svelte/store';
import { jwtDecode } from 'jwt-decode';
import { apiClient } from '$lib/api/client';
import { getLocalStorageItem, setLocalStorageItem } from '$lib/utils/localStorage.util';
import type { RpcError } from 'grpc-web';

const error = writable('');

const accessTokenCacheKey = 'access-token';
const refreshTokenCacheKey = 'refresh-token';

const isTokenStillValid = (token: string) => {
	const payload = jwtDecode(token);
	const tokenExp = payload.exp ?? 0;
	const now = new Date().getTime() / 1000;
	const expired = tokenExp < now;
	return !expired;
};

const refreshToken = async () => {
	error.set('');
	const refreshToken = getLocalStorageItem(refreshTokenCacheKey);
	if (!refreshToken) {
		// Todo: Write a custom error
		throw new Error('Refresh token missing');
	}
};

export const getAccessToken = async () => {
	const token: null | string = getLocalStorageItem(accessTokenCacheKey);
	// Token in cache but has expired
	if (token && !isTokenStillValid(token)) {
		// Attempt a refresh
		await refreshToken();
		return getLocalStorageItem(accessTokenCacheKey);
	}

	return token;
};

export const login = async (email: string, password: string) => {
	error.set('');
	const client = apiClient();

	try {
		const { response: r } = await client.getToken({
			email,
			password
		});

		setLocalStorageItem(accessTokenCacheKey, r.accessToken);
		setLocalStorageItem(refreshTokenCacheKey, r.refreshToken);
		return true;
	} catch (e) {
		const rpcError = e as RpcError;
		error.set(rpcError.code.toString());
		return false;
	}
};

export const isAuthenticated = async (): Promise<boolean> => await getAccessToken() !== '';
export const authError = readonly(error);