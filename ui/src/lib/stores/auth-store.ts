import { readable } from 'svelte/store';
import { jwtDecode } from 'jwt-decode';
import { apiClient } from '$lib/api/client';
import { getLocalStorageItem, setLocalStorageItem } from '$lib/utils/localStorage.util';

let error = '';

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
	error = '';
	const refreshToken = getLocalStorageItem(refreshTokenCacheKey);
	if (!refreshToken) {
		// Todo: Write a custom error
		throw new Error('Refresh token missing');
	}

	try {
	const client = apiClient();
	await client.getRefreshToken({
		refreshToken,
	});
} catch(error) {
	console.error(error);
}
};

export const getAccessToken = () => {
	const token: null | string = getLocalStorageItem(accessTokenCacheKey);
	// Token in cache but has expired
	if (token && !isTokenStillValid(token)) {
		// Attempt a refresh
		refreshToken();
		return getLocalStorageItem(accessTokenCacheKey);
	}

	return token;
};

export const login = async (email: string, password: string) => {
	error = '';
	const client = apiClient();

	try {
		const { response: r } = await client.getToken({
			email,
			password
		});

		setLocalStorageItem(accessTokenCacheKey, r.accessToken);
		setLocalStorageItem(refreshTokenCacheKey, r.refreshToken);
	} catch (e) {
		console.error(e);
	}
};

export const isAuthenticated = (): boolean => getAccessToken() !== '';
export const authError = readable(error, function (set) {
	set(error);
});
