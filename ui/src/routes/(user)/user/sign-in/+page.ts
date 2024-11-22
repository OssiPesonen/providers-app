import { isAuthenticated } from '$lib/stores/auth';
import { redirect } from '@sveltejs/kit';
import type { PageLoad } from './$types';

export const load: PageLoad = async() => {
    if (await isAuthenticated()) {
        redirect(302, '/');
    }
};