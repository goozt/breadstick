<script lang="ts">
	import { Modal, Helper } from 'flowbite-svelte';
	import { Button } from '$ui';
	import { newToast } from '$services/toast';
	import { queryClient } from '$services/api';
	import menuAPI from '$api/menu';
	import type { Item, MenuResult } from '$api/menu.d';

	export let open = false;

	const mutation = menuAPI.create();

	const appendItemtoMenu = async (item: Item) => {
		await queryClient.cancelQueries(['menu']);
		const previousMenu = queryClient.getQueryData<MenuResult>('menu');
		if (previousMenu && previousMenu.data) {
			queryClient.setQueryData('menu', {
				...previousMenu,
				data: {
					...previousMenu.data,
					items: [...(previousMenu.data.items ?? []), item]
				}
			});
		}
	};

	const onSubmitAction = (event: Event) => {
		const formData = new FormData(event.target as HTMLFormElement);
		try {
			const data = menuAPI.getFormData(formData);
			$mutation.mutate(data, {
				onSuccess: (data) => {
					if (data) {
						if (data.data) {
							newToast('green', 'Added ' + data.data.name + ' to menu.');
							appendItemtoMenu(data.data);
							open = false;
						} else if (data.error) {
							newToast('red', 'Failed: ' + data.error.details?.[0] ?? 'unknown reason');
						} else {
							newToast('red', 'Failed: unknown reason');
						}
					}
				},
				onError: (error) => {
					newToast('red', 'Failed: ' + error);
				}
			});
		} catch (e) {
			newToast('red', 'Failed to add new item to menu.');
		}
	};
</script>

<Modal bind:open size="xs" autoclose={false} class="w-full">
	<form class="flex flex-col space-y-6" on:submit|preventDefault={onSubmitAction}>
		<h3 class="text-xl font-medium text-gray-900 dark:text-white p-0">
			Add a new food item to the menu
		</h3>
		<slot />
		{#if $mutation.error}
			<Helper on:click={() => $mutation.reset()} class="mt-2" color="red">
				<span class="font-medium">Error: </span>
				{$mutation.error}
			</Helper>
		{/if}
		<Button type="submit" class="w-full1">Create</Button>
		<Button
			type="cancel"
			on:click={() => {
				open = false;
			}}
			class="w-full1">Cancel</Button
		>
	</form>
</Modal>
