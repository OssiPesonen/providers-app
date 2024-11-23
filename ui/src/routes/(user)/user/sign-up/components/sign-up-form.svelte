<script lang="ts">
	import { z } from 'zod';
	import { goto } from '$app/navigation';
	import { zod } from 'sveltekit-superforms/adapters';
	import { superForm, defaults } from 'sveltekit-superforms';
	import * as Form from '$components/ui/form';
	import * as Alert from '$components/ui/alert';
	import { Input } from '$components/ui/input/index.js';
	import { buttonVariants } from '$components/ui/button';
	import { authError, register } from '$lib/stores/auth';
	import { LuCircleAlert } from 'svelte-icons-pack/lu';
	import { Icon } from 'svelte-icons-pack';

	const formSchema = z
		.object({
			email: z.string().email(),
			password: z
				.string()
				.min(8, { message: 'Must be 8 or more characters long' })
				.max(50, { message: 'Must be 50 characters or less' }),
			confirmPassword: z
				.string()
				.min(8, { message: 'Must be 8 or more characters long' })
				.max(50, { message: 'Must be 50 characters or less' })
		})
		.superRefine(({ confirmPassword, password }, ctx) => {
			if (confirmPassword !== password) {
				ctx.addIssue({
					code: 'custom',
					message: 'The passwords did not match',
					path: ['confirmPassword']
				});
			}
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

		const success = await register($formData.email, $formData.password);

		// Only track success as error is already being subscribed to
		if (success) {
			return await goto('/user/sign-in?registered=1');
		}
	}
</script>

<div class="grid gap-6">
	<form use:enhance onsubmit={handleSubmit}>
		<Form.Field {form} name="email" class="mb-2">
			<Form.Control let:attrs>
				<Form.Label>Email address</Form.Label>
				<Input {...attrs} bind:value={$formData.email} />
			</Form.Control>
			<Form.FieldErrors class="font-normal text-xs" />
		</Form.Field>
		<Form.Field {form} name="password">
			<Form.Control let:attrs>
				<Form.Label>Password</Form.Label>
				<Input {...attrs} type="password" bind:value={$formData.password} />
			</Form.Control>
			<Form.FieldErrors class="font-normal text-xs" />
		</Form.Field>
		<Form.Field {form} name="confirmPassword" class="mb-4">
			<Form.Control let:attrs>
				<Form.Label>Password again</Form.Label>
				<Input {...attrs} type="password" bind:value={$formData.confirmPassword} />
			</Form.Control>
			<Form.FieldErrors class="font-normal text-xs" />
		</Form.Field>
		<Form.Button class="w-full mt-2">Create account</Form.Button>
	</form>
	{#if $authError === 'INTERNAL'}
		<Alert.Root variant="destructive" class="mt-4">
			<Icon src={LuCircleAlert} className="h-4 w-4" />
			<Alert.Title>Error</Alert.Title>
			<Alert.Description>Something went wrong when creating your account. Please try again later.</Alert.Description>
		</Alert.Root>
	{/if}
	<div class="relative">
		<div class="absolute inset-0 flex items-center">
			<span class="w-full border-t"></span>
		</div>
		<div class="relative flex justify-center text-xs uppercase">
			<span class="bg-background text-muted-foreground px-2"> Or </span>
		</div>
	</div>
	<a href="/user/sign-in" class={buttonVariants({ variant: 'outline' })}> Sign in </a>
	<a href="/" class={buttonVariants({ variant: 'ghost' })}> Back to Frontpage </a>
</div>
