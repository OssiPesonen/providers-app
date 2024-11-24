import { readonly, writable } from 'svelte/store';
import { jwtDecode } from 'jwt-decode';
import type { RpcError } from 'grpc-web';
import { apiClient } from '$lib/api/client';
import { getLocalStorageItem, removeLocalStorageItem, setLocalStorageItem } from '$lib/utils/localStorage.util';

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
		// Attempted to access refresh token that isn't there -> log out
		await logout();
	}

	try {
		const client = apiClient();
		const { response: r } = await client.refreshToken({
			refreshToken
		});

		setLocalStorageItem(accessTokenCacheKey, r.accessToken);
		setLocalStorageItem(refreshTokenCacheKey, r.refreshToken);
	} catch{
		await logout();
	}
};

export const getAccessToken = () => {
	const token: string = getLocalStorageItem(accessTokenCacheKey);
	
	// Token in cache but has expired
	if (token && !isTokenStillValid(token)) {
		refreshToken()
	}

	return token;
};

export const login = async (email: string, password: string) => {
	error.set('');

	try {
		const client = apiClient();
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

export const logout = async () => {
	error.set('');
	const client = apiClient();
	
	const refreshToken = getLocalStorageItem(refreshTokenCacheKey);

	if (refreshToken) {
		try {
			// We don't care about the response here
			await client.revokeRefreshToken({
				refreshToken,
			});
		} catch {
			// Ignore
		}
	}
	
	removeLocalStorageItem(accessTokenCacheKey);
	removeLocalStorageItem(refreshTokenCacheKey);
};

export const register = async (email: string, password: string) => {
	error.set('');
	try {
		const client = apiClient();
		const { status } = await client.registerUser({
			email,
			password,
			username: '',
		});

		return status.code === 'OK';
	} catch (e) {
		// This should only occur with network errors, or internal server errors
		const rpcError = e as RpcError;
		error.set(rpcError.code.toString());
		return false;
	}
};


export const isAuthenticated = () => getAccessToken() !== '';
export const authError = readonly(error);