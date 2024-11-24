<script lang="ts">
	import { goto } from '$app/navigation';
	import { logout } from '$lib/stores/auth.svelte';
	import { onMount } from 'svelte';
	import { clearUserInfo } from '$lib/stores/user.svelte';

	onMount(async () => {
		clearUserInfo();
		await logout();
		goto('/');
	});
</script>
