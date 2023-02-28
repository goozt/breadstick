<script lang="ts">
	import { page } from '$app/stores';
	import { ErrorPage } from '$ui/error';
	import { Button } from 'flowbite-svelte';
	let error = $page.error;
	let msg: string;
	$: msg = error ? error.message : "Something's missing";
	$: description = error
		? error.message != 'Not Found'
			? 'Something went wrong at our side. Please try again later.'
			: "Sorry, we can't find that page. You'll find lots to explore on the home page."
		: "Sorry, we can't find that page. You'll find lots to explore on the home page.";
</script>

<div class="flex items-center justify-center h-screen">
	<ErrorPage>
		<div slot="paragraph">
			<p class="mb-4 text-3xl tracking-tight font-bold text-gray-900 md:text-4xl dark:text-white">
				{msg}
			</p>
			<p class="mb-4 text-lg font-light text-gray-500 dark:text-gray-400">
				{description}
			</p>
			<Button shadow="blue" gradient color="blue" href="/" size="lg">Back to Home</Button>
		</div>
	</ErrorPage>
</div>
