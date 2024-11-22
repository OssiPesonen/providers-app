<script lang="ts">
	import { LuMoon } from 'svelte-icons-pack/lu';
	import { ModeWatcher, toggleMode, mode } from 'mode-watcher';
	import { Icon } from 'svelte-icons-pack';
	import { Switch } from '$lib/components/ui/switch';
	import '../css/globals.css';

	let { children } = $props();
</script>

<div id="root">
	<ModeWatcher />
	<div class="fixed top-5 right-5 z-50 flex align-middle items-center">
		<Switch id="dark-mode" checked={$mode === 'dark'} onCheckedChange={toggleMode} />
		<Icon src={LuMoon} className="h-4 w-4 ml-2" />
	</div>
	{@render children()}
</div>
