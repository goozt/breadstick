<script lang="ts">
	import { Heading } from 'flowbite-svelte';
	import type { NavMenu } from '$types';
	import { Breadcrumb } from '$ui';
	import { page } from '$app/stores';
	export let heading: NavMenu;
</script>

<header class="bg-yellow-300 shadow dark:bg-yellow-400">
	<div class="mx-auto max-w-7xl pt-5 pb-1 px-4 sm:px-6 lg:px-8">
		<Breadcrumb path={$page.url.pathname} />
		<Heading tag="h3" color="text-gray-700 dark:text-gray-900" class="mb-4">{heading.name}</Heading>
	</div>
</header>
