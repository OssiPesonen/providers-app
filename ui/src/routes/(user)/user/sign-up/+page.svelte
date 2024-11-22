<script lang="ts">
	import UserSignupForm from './components/sign-up-form.svelte';
	import * as Card from '$lib/components/ui/card';
</script>

<div
	class="container relative h-screen flex-col items-center justify-center grid lg:max-w-none grid-cols-1 lg:grid-cols-2 lg:px-0"
>
	<div class="bg-muted relative hidden h-full flex-col p-10 text-white lg:flex dark:border-r">
		<div
			class="absolute inset-0 bg-cover"
			style="
					background-image:
						url(https://images.unsplash.com/photo-1658409646482-e1c6176fecce?q=80&w=2500&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D);"
		></div>

		<div class="relative z-20 flex items-center text-lg font-semibold">Traffic Lights</div>
	</div>
	<div class="lg:p-8 pt-16 md:pt-0">
		<div class="mx-auto flex w-full flex-col justify-center space-y-6 sm:w-[350px]">
			<div class="flex flex-col space-y-4 text-center">
				<h1 class="tracking-tight">Create an account</h1>
				<p class="text-muted-foreground text-sm">Enter your information below.</p>
			</div>
			<UserSignupForm />
		</div>
	</div>
</div>
