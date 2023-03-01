<script lang="ts">
	import { page } from '$app/stores';
	import { ErrorPage } from '$ui/error';
	import { Button } from 'flowbite-svelte';

	type Error = {
		code: number;
		message: string;
		description: string;
	};

	const errors: Error[] = [
		{
			code: 404,
			message: 'Page Not Found',
			description: "Sorry, we can't find that page."
		},
		{
			code: 500,
			message: 'Oops!',
			description: 'Something went wrong at our side.'
		}
	];

	$: error =
		errors.find((err) => {
			return err.code === $page.status;
		}) ?? errors[0];
</script>

<div class="flex items-center justify-center h-screen">
	<ErrorPage>
		<div slot="paragraph">
			<p class="text-2xl font-semibold text-yellow-400 pt-4">{error.code}</p>
			<h1 class="mb-4 text-3xl tracking-tight font-bold text-gray-900 md:text-4xl dark:text-white">
				{error.message}
			</h1>
			<p class="mb-4 text-lg font-light text-gray-500 dark:text-gray-400">
				{error.description}
			</p>
			<Button shadow="blue" gradient color="blue" href="/" size="lg">Back to Home</Button>
		</div>
	</ErrorPage>
</div>
