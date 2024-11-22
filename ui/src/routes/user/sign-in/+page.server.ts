import type { Actions, PageServerLoad } from './$types.js';
import { fail, superValidate } from 'sveltekit-superforms';
import { formSchema } from './schema';
import { zod } from 'sveltekit-superforms/adapters';
import { apiClient } from '$lib/api/client.js';

export const load: PageServerLoad = async () => {
	return {
		form: await superValidate(zod(formSchema))
	};
};

export const actions: Actions = {
	default: async (event) => {
		const form = await superValidate(event, zod(formSchema));

		if (!form.valid) {
			return fail(400, { form });
		}

		try {
			const client = apiClient();
			const { response } = await client.getToken({
				email: form.data.email,
				password: form.data.password
			});

			return { 
                tokens: {
                    accessToken: response.accessToken,
                    refreshToken: response.refreshToken,
                    exp: response.exp,
                },
                form
            }
		} catch {
			return fail(400, { form });
		}
	}
};
