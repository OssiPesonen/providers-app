<script lang="ts">
	import CreateProviderForm from './components/create-provider-form.svelte';
</script>

<div class="w-full h-screen flex items-center justify-center">
	<div class="w-full max-w-3xl">
		<h1 class="tracking-tight">Create a provider</h1>
		<p class="pb-8 text-muted-foreground text-md">Enter your company information below.</p>
		<CreateProviderForm />
	</div>
</div>
