<script lang="ts">
	import type { PageData } from './$types';
	let { data }: { data: PageData } = $props();
</script>

<div class="page">
	<div class="main">
		<h1>Hello!</h1>
		<p>Here is a list of providers in the system:</p>
		<ol>
			{#each data.providers as { id, name, city, region }}
				<li>{name}, {city}, {region}</li>
			{/each}
		</ol>
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

	.main {
		max-width: 480px;
		width: 100%;
	}

	.main ol {
		font-family: var(--font-geist-mono);
		padding-left: 0;
		margin: 0;
		font-size: 14px;
		line-height: 24px;
		letter-spacing: -0.01em;
		list-style-position: inside;
	}

	.main li:not(:last-of-type) {
		margin-bottom: 8px;
	}

	.page {
		padding: 32px;
		padding-bottom: 80px;
	}

	.main {
		align-items: center;
	}
</style>
