<script lang="ts">
	import { Button } from '$lib/components/ui/button/index.js';
	import type { ActionData, PageData } from './$types';
	import UserAuthForm from './components/sign-in-form.svelte';
	let { data, form }: { data: PageData, form: ActionData } = $props();
</script>
<div
	class="container relative hidden h-screen flex-col items-center justify-center md:grid lg:max-w-none lg:grid-cols-2 lg:px-0"
>
	<Button
		href="/user/sign-up"
		variant="ghost"
		class="absolute right-4 top-4 md:right-8 md:top-8 hidden"
	>
		Sign up
	</Button>
	<div class="bg-muted relative hidden h-full flex-col p-10 text-white lg:flex dark:border-r">
		<div
			class="absolute inset-0 bg-cover"
			style="
				background-image:
					url(https://images.unsplash.com/photo-1533134486753-c833f0ed4866?q=80&w=3540&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D);"
		></div>

		<div class="relative z-20 flex items-center text-lg font-semibold">
			Traffic Lights
		</div>
	</div>
	<div class="lg:p-8">
		<div class="mx-auto flex w-full flex-col justify-center space-y-6 sm:w-[350px]">
			<div class="flex flex-col space-y-4 text-center">
				<h1 class="tracking-tight">Sign in</h1>
				<p class="text-muted-foreground text-sm">Enter your email and password to sign in</p>
			</div>
			<UserAuthForm data={data.form} actionData={form} />
		</div>
	</div>
</div>
