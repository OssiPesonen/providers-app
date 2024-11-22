<script lang="ts">
	import { z } from 'zod';
	import type { PageData } from '../$types'
	import { goto } from '$app/navigation';
	import { Icon } from 'svelte-icons-pack';
	import { LuCircleAlert } from 'svelte-icons-pack/lu';
	import { zod } from 'sveltekit-superforms/adapters';
	import { superForm, defaults } from 'sveltekit-superforms';
	import * as Form from '$lib/components/ui/form';
	import * as Alert from '$lib/components/ui/alert';
	import { Input } from '$lib/components/ui/input/index.js';
	import { login, authError } from '$lib/stores/auth';

	const formSchema = z.object({
		email: z.string().email(),
		password: z
			.string()
			.min(8, { message: 'Must be 8 or more characters long' })
			.max(50, { message: 'Must be 50 characters or less' })
	});

	const form = superForm(defaults(zod(formSchema)), {
		SPA: true,
		validators: zod(formSchema),
		resetForm: false
	});

	const { form: formData, enhance, validateForm, errors } = form;

	async function handleSubmit(event: Event) {
		event.preventDefault();
		const result = await validateForm();

		if (!result.valid) {
			errors.update((v) => {
				return {
					...v,
					email: result.errors.email,
					password: result.errors.password
				};
			});

			return;
		}

		const success = await login($formData.email, $formData.password);

		// Only track success as error is already being subscribed to
		if (success) {
			return await goto('/');
		}
	}
</script>

<div class="grid gap-6">
	<form method="POST" use:enhance onsubmit={handleSubmit}>
		<Form.Field {form} name="email" class="mb-2">
			<Form.Control let:attrs>
				<Form.Label>Username</Form.Label>
				<Input {...attrs} bind:value={$formData.email} />
			</Form.Control>
			<Form.FieldErrors />
		</Form.Field>
		<Form.Field {form} name="password">
			<Form.Control let:attrs>
				<Form.Label>Password</Form.Label>
				<Input {...attrs} type="password" bind:value={$formData.password} />
			</Form.Control>
			<Form.FieldErrors />
		</Form.Field>
		<Form.Button class="w-full mt-2">Submit</Form.Button>
		{#if $authError === 'UNAUTHENTICATED'}
			<Alert.Root variant="destructive" class="mt-4">
				<Icon src={LuCircleAlert} className="h-4 w-4" />
				<Alert.Title>Error</Alert.Title>
				<Alert.Description>Invalid email or password. Please try again.</Alert.Description>
			</Alert.Root>
		{/if}
	</form>
</div>
