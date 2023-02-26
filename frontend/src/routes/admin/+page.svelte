<script lang="ts">
	import { Header } from '$ui';
	import { pageTracker, itemCollection } from '$store';
	import { Grid, Item } from '$ui/menu';
	import type { MenuItem } from '$ui-types';
	$pageTracker = 'admin';

	const findItem = (id: string): MenuItem => {
		return;
	};

	const removeItem = (event: Event) => {
		const e = event.target as Element;
		const id = e.getAttribute('data-id');
		if (id != null) {
			const items: MenuItem[] = $itemCollection;
			const item = items.find((item) => item.id === parseInt(id));
			const index = items.indexOf(item);
			if (index !== -1) {
				items.splice(index, 1);
			}
			$itemCollection = items;
			console.log('Removed item', id, 'successfully');
		}
	};
</script>

<Header heading="Admin Dashboard" />
<div class="md:container md:mx-auto py-16 px-4">
	<Grid>
		{#each $itemCollection as item}
			<Item
				id={String(item.id)}
				name={item.name}
				price={item.price}
				image={item.imageurl}
				on:click={removeItem}>Remove</Item
			>
		{/each}
	</Grid>
</div>
