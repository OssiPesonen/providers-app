<script lang="ts">
	import Spinner from '$lib/components/ui/spinner/spinner.svelte';
	import { onMount } from 'svelte';
	import UserAuthForm from './components/sign-in-form.svelte';
	import { isAuthenticated } from '$lib/stores/auth';
	import { goto } from '$app/navigation';
	import { toast } from 'svelte-sonner';

	// Assume users are authenticated
	let authenticated = $state(true);

	onMount(async () => {
		if (await isAuthenticated()) {
			goto('/');
		} else {
			authenticated = false;
		}
	});
</script>

{#if !authenticated }
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
			<h3 class="relative z-20 flex items-center text-xl font-black">Traffic Lights</h3>
		</div>
		<div class="lg:p-8 pt-16 md:pt-0">
			<div class="mx-auto flex w-full flex-col justify-center space-y-6 max-w-lg">
				<div class="flex flex-col space-y-4 text-center">
					<h1 class="tracking-tight">Sign in</h1>
					<p class="text-muted-foreground text-sm">Enter your email and password to sign in.</p>
				</div>
				<UserAuthForm />
			</div>
		</div>
	</div>
{:else}
	<Spinner fullpage />
{/if}
