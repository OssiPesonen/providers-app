<script lang="ts">
	import { ModeWatcher, mode } from 'mode-watcher';
	import '$css/globals.css';
	import { Toaster } from "$components/ui/sonner";

	let { children } = $props();
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
	{@render children()}
</div>
