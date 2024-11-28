<script lang="ts">
	// Todo: Fetch account info here and display it
	// allow editing info via another page
	import * as Table from '$components/ui/table';
	import { user } from '$lib/stores/user.svelte';

	const providers = user.info?.providerAccounts;
</script>

<div class="w-full flex items-center justify-center">
	<div class="w-full max-w-5xl">
		<h1 class="tracking-tight">My account</h1>
		<h3>Providers</h3>
		{#if providers}
			<Table.Root>
				<Table.Header>
					<Table.Row>
						<Table.Head class="">Name</Table.Head>
						<Table.Head>Area</Table.Head>
						<Table.Head>Line of Business</Table.Head>
					</Table.Row>
				</Table.Header>
				<Table.Body>
					{#each providers as provider, i (i)}
						<Table.Row>
							<Table.Cell class="font-medium">{provider.name}</Table.Cell>
							<Table.Cell>{provider.city}, { provider.region }</Table.Cell>
							<Table.Cell>{provider.lineOfBusiness}</Table.Cell>
						</Table.Row>
					{/each}
				</Table.Body>
			</Table.Root>
		{/if}
	</div>
</div>
