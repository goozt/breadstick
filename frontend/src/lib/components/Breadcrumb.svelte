<script lang="ts">
	import { BreadcrumbItem } from '$ui';
	import { Breadcrumb } from 'flowbite-svelte';
	export let path: string;
	let crumbs: BreadcrumbType[];
	$: {
		const tokens = path.split('/').filter((t) => t !== '');
		let tokenPath = '';
		crumbs = tokens.map((t: string) => {
			tokenPath += '/' + t;
			return {
				label: t.charAt(0).toUpperCase() + t.slice(1),
				href: tokenPath
			};
		});
	}
</script>

<Breadcrumb class="mb-4 dark:text-black">
	<BreadcrumbItem href="/" home>Home</BreadcrumbItem>
	{#each crumbs as item, i}
		{#if i == crumbs.length - 1}
			<BreadcrumbItem>{item.label}</BreadcrumbItem>
		{:else}
			<BreadcrumbItem href={item.href}>{item.label}</BreadcrumbItem>
		{/if}
	{/each}
</Breadcrumb>
