<script lang="ts">
	import * as Card from '$components/ui/card';
	import apiClient from '$lib/api/client';
	import { Input } from '$components/ui/input';
	import { getAccessToken } from '$lib/stores/auth.svelte';
	import { Provider } from '$lib/proto/providers_app_service';
	import { Icon } from 'svelte-icons-pack';
	import { LuLoaderCircle, LuSearch } from 'svelte-icons-pack/lu';

	let providerData = $state<Provider[]>([]);
	let inputVal = $state('');
	let searching = $state(false);
	let timer: number;

	const debounce = (v: string) => {
		clearTimeout(timer);
		timer = setTimeout(() => {
			if (v !== '' && v.length > 2) {
				searchProviders(v);
			}
		}, 750);
	};

	$effect(() => {
		if (inputVal !== '' && inputVal.length > 2) {
			searching = true;
			debounce(inputVal);
		}
	});

	async function searchProviders(searchInput: string) {
		const { response } = await apiClient().searchProviders(
			{ searchWords: searchInput },
			{
				meta: {
					Authorization: `Bearer ${getAccessToken()}`
				}
			}
		);

		if (response) {
			providerData = response.providers;
		}

		searching = false;
	}
</script>

<div class="page">
	<div class="w-screen max-w-5xl">
		<h1 class="text-center font-black text-4xl md:text-7xl mb-8 pt-16">providers.app</h1>
		<p class="md:text-xl font- text-center pb-4 dark:text-slate-400 text-slate-700">
			Search for service providers in your area!
		</p>
		<p class="text-sm dark:text-slate-500 text-slate-600 text-center pb-8">
			Find service providers near you, and see their availability. Consultants, teachers,
			instructors, coaches, you name it!
		</p>
		<div class="space-x-4 max-w-xl m-auto mb-10">
			<div class="relative rounded-full pl-4 pr-4">
				{#if searching}
					<Icon src={LuLoaderCircle} className="flex-1 z-10 absolute left-7 top-5 animate-spin" />
				{:else}
					<Icon src={LuSearch} className="flex-1 z-10 absolute left-7 top-5" />
				{/if}
				<Input
					bind:value={inputVal}
					placeholder="Name, city, region, line of business..."
					class="max-w h-12 md:h-14 pl-12 text-md border-0"
				/>
			</div>
		</div>
		<div class="relative">
			<div class="absolute inset-0 flex items-center">
				<span class="w-full border-t"></span>
			</div>
		</div>
		{#if searching}{/if}
		{#if providerData.length > 0}
			<div class="grid grid-cols-1 md:grid-cols-3 gap-2 mt-8 pt-10">
				{#each providerData as { name, city, region, lineOfBusiness }}
					<Card.Root
						class="hover:ring hover:ring-ring hover:outline-none transition-all cursor-pointer"
					>
						<Card.Header class="pl-4 pr-4 pt-4">
							<Card.Title class="mb-0 pb-2 text-base" tag="h3">{name}</Card.Title>
							<Card.Description>{lineOfBusiness}</Card.Description>
						</Card.Header>
						<Card.Content class="p-4">
							<p class="text-sm dark:text-purple-400 font-medium">{city}, {region}</p>
						</Card.Content>
					</Card.Root>
				{/each}
			</div>
		{/if}
	</div>
</div>

<style scoped>
	.page {
		--gray-rgb: 0, 0, 0;
		--gray-alpha-200: rgba(var(--gray-rgb), 0.08);
		--gray-alpha-100: rgba(var(--gray-rgb), 0.05);

		--button-primary-hover: #383838;
		--button-secondary-hover: #f2f2f2;
		display: flex;
		justify-content: center;
		align-items: center;
		flex-wrap: wrap;
		width: 100%;
		height: auto;
		min-height: 100vh;
	}

	.page {
		padding: 32px;
		padding-bottom: 80px;
	}
</style>
