<script lang="ts">
	import { Heading } from 'flowbite-svelte';
	import { Breadcrumb } from '$ui';
	import { getHeading } from '$services/navigation';
	export let page: string;
</script>

<header class="bg-yellow-300 shadow dark:bg-yellow-400">
	<div class="mx-auto max-w-7xl pt-5 pb-1 px-4 sm:px-6 lg:px-8">
		<Breadcrumb path={page} />
		<Heading tag="h3" color="text-gray-700 dark:text-gray-900" class="mb-4"
			>{getHeading(page)}</Heading
		>
	</div>
</header>
