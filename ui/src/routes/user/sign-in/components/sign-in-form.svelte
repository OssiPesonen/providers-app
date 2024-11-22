<script lang="ts">
	import { zodClient } from 'sveltekit-superforms/adapters';
	import { Icon } from 'svelte-icons-pack';
	import { LuCircleAlert } from "svelte-icons-pack/lu";
	import { applyAction, deserialize } from '$app/forms';
	import type { ActionResult } from '@sveltejs/kit';
	import { type SuperValidated, type Infer, superForm } from 'sveltekit-superforms';
	import * as Form from '$lib/components/ui/form';
	import { Input } from '$lib/components/ui/input/index.js';
	import * as Alert from '$lib/components/ui/alert/index.js';
	import type { ActionData } from '../$types';
	
	import { formSchema, type FormSchema } from '../schema';

	interface Props {
		data: SuperValidated<Infer<FormSchema>>;
		actionData: ActionData;
	}

	const { data, actionData }: Props = $props();
	const form = superForm(data, {
		validators: zodClient(formSchema)
	});

	const { form: formData, enhance } = form;
	
	async function handleSubmit(event: Event &{ currentTarget: EventTarget & HTMLFormElement}) {
		event.preventDefault();
		const data = new FormData(event.currentTarget);

		const response = await fetch(event.currentTarget.action, {
			method: 'POST',
			body: data
		});

		const result: ActionResult<{ tokens: { accessToken: string, refreshToken: string }}> = deserialize(await response.text());

		if (result.type === 'success') {
			const { refreshToken, accessToken } = result.data
		}

		applyAction(result);
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
		<Form.Button class="w-full mt-2" >Submit</Form.Button>

		{#if actionData?.success === false}
			<Alert.Root variant="destructive" class="mt-4 bg-red-700/10">
				<Icon src={LuCircleAlert} />
				<Alert.Title>Invalid credentials</Alert.Title>
				<Alert.Description>Invalid email or password. Please try again</Alert.Description>
			</Alert.Root>
		{/if}
	</form>
</div>
