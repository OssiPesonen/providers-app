<script lang="ts">
	import { z } from 'zod';
	import { zod } from 'sveltekit-superforms/adapters';
	import { superForm, defaults } from 'sveltekit-superforms';
	import type { RpcError } from 'grpc-web';
	import * as Form from '$components/ui/form';
	import * as Alert from '$components/ui/alert';
	import { Input } from '$components/ui/input/index.js';
	import { apiClient } from '$lib/api/client';
	import { toast } from 'svelte-sonner';
	import { goto } from '$app/navigation';
	import { Icon } from 'svelte-icons-pack';
	import { LuCircleAlert } from 'svelte-icons-pack/lu';
	import { getAccessToken } from '$lib/stores/auth';

	let error = $state('');
	const formSchema = z.object({
		name: z.string().min(2).max(200),
		region: z.string().min(2).max(200),
		city: z.string().min(2).max(200),
		lineOfBusiness: z.string().min(2).max(200)
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
					name: result.errors.name
				};
			});

			return;
		}

		try {
			const client = apiClient();
			const success = await client.createProvider({
				name: $formData.name,
				lineOfBusiness: $formData.lineOfBusiness,
				region: $formData.region,
				city: $formData.city
			}, {
				meta: {
					"Authorization": `Bearer ${getAccessToken()}`,
				}
			});

			if (success.response.id) {
				toast.success('Congratulations!', {
					description: 'You now have a provider account',
					dismissable: true,
					position: 'top-center'
				});

				// Todo: replace with account page
				goto('/account/info');
			}
		} catch (e) {
			const err = e as RpcError;
			error = err.message;
		}
	}
</script>

<div class="grid gap-6">
	<form use:enhance onsubmit={handleSubmit}>
		<Form.Field {form} name="name" class="mb-2">
			<Form.Control let:attrs>
				<Form.Label>Company / your name</Form.Label>
				<Input {...attrs} bind:value={$formData.name} />
			</Form.Control>
			<Form.FormDescription>
				Enter your name if you are an individual entrepeneur
			</Form.FormDescription>
			<Form.FieldErrors />
		</Form.Field>
		<Form.Field {form} name="lineOfBusiness" class="mb-2">
			<Form.Control let:attrs>
				<Form.Label>Line of business</Form.Label>
				<Input {...attrs} bind:value={$formData.lineOfBusiness} />
			</Form.Control>
			<Form.FormDescription>e.g. Healthcare, Fitness, IT, Accounting</Form.FormDescription>
			<Form.FieldErrors />
		</Form.Field>
		<Form.Field {form} name="region" class="mb-2">
			<Form.Control let:attrs>
				<Form.Label>Region</Form.Label>
				<Input {...attrs} bind:value={$formData.region} />
			</Form.Control>
			<Form.FieldErrors />
		</Form.Field>
		<Form.Field {form} name="city" class="mb-2">
			<Form.Control let:attrs>
				<Form.Label>City</Form.Label>
				<Input {...attrs} bind:value={$formData.city} />
			</Form.Control>
			<Form.FieldErrors />
		</Form.Field>
		<Form.Button class="w-full mt-2">Create</Form.Button>
	</form>
	{#if error !== ''}
		<Alert.Root variant="destructive" class="mt-4">
			<Icon src={LuCircleAlert} className="h-4 w-4" />
			<Alert.Title>Error</Alert.Title>
			<Alert.Description>{error}</Alert.Description>
		</Alert.Root>
	{/if}
</div>
