<script lang="ts">
	import { Button } from '$ui';
	export let id: string | undefined = undefined,
		name: string = 'Apple',
		price: number = 999,
		image: string = '';
	$: price.toFixed(2);
</script>

<div
	class="group bg-white dark:bg-gray-800 text-gray-500 dark:text-gray-400 rounded-xl border border-gray-200 dark:border-gray-700 shadow-md flex sm:max-w-sm flex-col p-0"
>
	<div
		class="aspect-w-1 aspect-h-1 w-full overflow-hidden rounded-t-xl bg-gray-200 xl:aspect-w-7 xl:aspect-h-8"
	>
		<img
			src={image}
			alt="Tall slender porcelain bottle with natural clay textured body and cork stopper."
			class="h-full w-full object-cover object-center group-hover:opacity-75"
		/>
	</div>
	<div class="p-5">
		<h5 class="text-xl font-semibold tracking-tight text-gray-900 dark:text-white">
			{name}
		</h5>
		<div class="flex justify-between items-centermt-1">
			<span class="text-lg font-medium text-gray-900 dark:text-white">
				${price}
			</span>
			<Button {id} on:click><slot /></Button>
		</div>
	</div>
</div>
