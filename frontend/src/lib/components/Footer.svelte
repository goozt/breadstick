<script lang="ts">
	import { Footer, FooterBrand, FooterCopyright } from 'flowbite-svelte';
	export let name: string = 'MyRestaurant',
		description: string = 'This is a restaurant';
</script>

<Footer
	footerType="custom"
	customClass="justify-self-end pt-16 pb-8 px-4 lg:pt-24 lg:pb-10 bg-gray-200 dark:bg-gray-800 z-40"
>
	<div class="grid gap-12 lg:grid-cols-6 lg:gap-24">
		<div class="col-span-2">
			<FooterBrand href="/" src="/logo.png" {name} aClass="flex mb-6" />
			<p class="text-gray-600 dark:text-gray-400">
				{description}
			</p>
		</div>
		<slot />
	</div>
	<hr class="my-6 border-gray-200 sm:mx-auto dark:border-gray-700 lg:my-8" />
	<div class="sm:flex sm:items-center sm:justify-between">
		<FooterCopyright href="/" by={name} />
	</div>
</Footer>
