<script lang="ts">
	import { onMount } from 'svelte';
	import { Toaster } from "$components/ui/sonner";
	import { LuLogOut, LuMoon } from 'svelte-icons-pack/lu';
	import { ModeWatcher, toggleMode, mode } from 'mode-watcher';
	import { Icon } from 'svelte-icons-pack';
	import { Switch } from '$components/ui/switch';
	import '$css/globals.css';
	import { buttonVariants } from '$components/ui/button';
	import { isAuthenticated } from '$lib/stores/auth';

	let { children } = $props();
	let isAuthed = $state(false);

	onMount(async () => {
		isAuthed = await isAuthenticated();
	});
</script>

<div id="root">
	<Toaster />
	<ModeWatcher />
	{#if $mode === 'dark'}
		<div class="fixed top-0 left-0 bottom-0 right-0 -z-30 h-full w-full bg-slate-950">
			<div
				class="absolute bottom-0 left-[-25%] right-0 top-[20%] h-[1500px] w-[1500px] rounded-full bg-[radial-gradient(circle_farthest-side,rgba(108,40,217,.1),rgba(255,255,255,0))]"
			></div>
			<div
				class="absolute bottom-0 right-[-20%] top-[-10%] h-[750px] w-[750px] rounded-full bg-[radial-gradient(circle_farthest-side,rgba(108,40,217,.1),rgba(255,255,255,0))]"
			></div>
		</div>
	{:else}
		<div
			class="absolute top-0 z-[-2] h-screen w-screen bg-white bg-[radial-gradient(ellipse_80%_80%_at_50%_-20%,rgba(120,119,198,0.3),rgba(255,255,255,0))]"
		></div>
	{/if}
	<nav
		class="fixed w-full top-0 z-50 dark:bg-slate-950 bg-white backdrop-filter backdrop-blur-lg bg-opacity-30 pl-4 pr-4"
	>
		<div class="flex gap-4 items-center justify-end h-16">
			<div class=" flex align-middle items-center">
				<Switch id="dark-mode" checked={$mode === 'dark'} onCheckedChange={toggleMode} />
				<Icon src={LuMoon} className="h-4 w-4 ml-2" />
			</div>
			{#if !isAuthed}
				<a href="/user/sign-in" class={buttonVariants({ variant: 'ghost' })}> Sign in </a>
				<a href="/user/sign-up" class={buttonVariants({ variant: 'outline' })}> Sign up </a>
			{:else}
				<a href="/user/logout" class={buttonVariants({ variant: 'outline' })}>
					<Icon src={LuLogOut} className="h-4 w-4 mr-2" />
					Sign out
				</a>
			{/if}
		</div>
	</nav>
	{@render children()}
</div>
