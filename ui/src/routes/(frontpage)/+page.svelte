<script lang="ts">
	import type { PageData } from '../$types';
	import * as Card from '$components/ui/card';
	import Input from '$components/ui/input/input.svelte';
	import { Button } from '$components/ui/button';
	import { LuSearch } from 'svelte-icons-pack/lu';
	import { Icon } from 'svelte-icons-pack';
	let { data }: { data: PageData } = $props();
</script>

<div class="page">
	<div class="w-screen max-w-5xl">
		<h1 class="text-center font-black text-4xl md:text-7xl mb-8 pt-16">providers.app</h1>
		<p class="md:text-xl font- text-center pb-4 dark:text-slate-400 text-slate-500">
			Search for service providers in your area!
		</p>
		<p class="text-sm text-slate-600 text-center pb-8">
			Find service providers to see their availability. Consultants, therapists, instructors,
			coaches, you name it!
		</p>
		<div class="flex space-x-4 max-w-xl m-auto mb-10">
			<Input
				placeholder="Name, city, region, line of business..."
				class="max-w h-12 md:h-14 text-md pl-6 rounded-full"
			/>
			<Button class="h-12 md:h-14 pl-8 pr-8 text-md rounded-full">
				<Icon src={LuSearch} className="w-5 h-5 mr-3" />
				Search
			</Button>
		</div>
		<div class="relative">
			<div class="absolute inset-0 flex items-center">
				<span class="w-full border-t"></span>
			</div>
		</div>
		<div class="grid grid-cols-1 md:grid-cols-3 gap-2 mt-8 pt-10">
			{#each data.providers as { id, name, city, region, lineOfBusiness }}
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
