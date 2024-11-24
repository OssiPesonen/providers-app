import { isAuthenticated } from '$lib/stores/auth.svelte';
import { getUserInfo } from '$lib/stores/user.svelte';
import type { LayoutLoad } from './$types';

export const load: LayoutLoad = () => {
	if (isAuthenticated()) {
		getUserInfo();
	}
};