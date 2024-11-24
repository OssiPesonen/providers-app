import { isAuthenticated } from "$lib/stores/auth.svelte";
import type { PageLoad } from "../$types";

export const load: PageLoad = () => {
	return {
        isUserLoggedIn: isAuthenticated()
    }
}