<script lang="ts">
	import { goto } from '$app/navigation';
	import { logout } from '$lib/stores/auth.svelte';
	import { onMount } from 'svelte';

	onMount(async () => {
		await logout();
		goto('/');
	});
</script>
