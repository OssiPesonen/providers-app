<script lang="ts">
    import { page } from '$app/stores';
	import { z } from 'zod';
	import { goto } from '$app/navigation';
	import { Icon } from 'svelte-icons-pack';
	import { LuCircleAlert, LuCircleCheck } from 'svelte-icons-pack/lu';
	import { zod } from 'sveltekit-superforms/adapters';
	import { superForm, defaults } from 'sveltekit-superforms';
	import * as Form from '$components/ui/form';
	import * as Alert from '$components/ui/alert';
	import { Input } from '$components/ui/input/index.js';
	import { login, authError } from '$lib/stores/auth';
	import { buttonVariants } from '$components/ui/button';
	
	import { toast } from "svelte-sonner";

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

	const hasRegistered = $page.url.searchParams.has('registered');
	if (hasRegistered) {
		toast.success('Congratulations', {
			position: 'top-center',
			description: 'You have successfully created an account. You may now log in.',
			dismissable: true,
		});
	}

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
	<form use:enhance onsubmit={handleSubmit}>
		<Form.Field {form} name="email" class="mb-2">
			<Form.Control let:attrs>
				<Form.Label>Email address</Form.Label>
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
		<Form.Button class="w-full mt-2">Sign in</Form.Button>
	</form>
	{#if $authError === 'UNAUTHENTICATED'}
		<Alert.Root variant="destructive" class="mt-4">
			<Icon src={LuCircleAlert} className="h-4 w-4" />
			<Alert.Title>Error</Alert.Title>
			<Alert.Description>Invalid email or password. Please try again.</Alert.Description>
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
	<a href="/user/sign-up" class={buttonVariants({ variant: 'outline' })}> Sign up </a>
	<a href="/" class={buttonVariants({ variant: 'ghost' })}> Back to Frontpage </a>
</div>
