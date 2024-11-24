<script>
	import { Icon } from 'svelte-icons-pack';
	import { toggleMode, mode } from 'mode-watcher';
	import { LuLogOut, LuMoon } from 'svelte-icons-pack/lu';
	import { Switch } from '$components/ui/switch';
	import { buttonVariants } from '$components/ui/button';
	import { isAuthenticated } from '$lib/stores/auth.svelte';
	import { user } from '$lib/stores/user.svelte';

	let isAuthed = isAuthenticated();
</script>

<nav
	class="fixed w-full top-0 z-50 dark:bg-slate-950 bg-white backdrop-filter backdrop-blur-lg bg-opacity-30 pl-4 pr-4"
>
	<div class="flex gap-4 items-center h-16 max-w-[90%] m-auto">
		<div class="flex-1">
			<a href="/" class="p-2 font-bold hover:opacity-75 transition-opacity"> Frontpage </a>
		</div>
		<div class="flex-1 flex gap-4 justify-end">
			<div class="flex align-middle items-center mr-6">
				<Switch id="dark-mode" checked={$mode === 'dark'} onCheckedChange={toggleMode} />
				<Icon src={LuMoon} className="h-4 w-4 ml-2" />
			</div>
			{#if !isAuthed}
				<a href="/user/sign-in" class={buttonVariants({ variant: 'ghost' })}> Sign in </a>
				<a href="/user/sign-up" class={buttonVariants({ variant: 'outline' })}> Sign up </a>
			{:else}
				{#if user && user.info}
					<a href="/account/info" class={buttonVariants({ variant: 'outline' })}> My Account </a>
				{/if}
				<a href="/user/logout" class={buttonVariants({ variant: 'outline' })}>
					<Icon src={LuLogOut} className="h-4 w-4 mr-2" />
					Sign out
				</a>
			{/if}
		</div>
	</div>
</nav>
