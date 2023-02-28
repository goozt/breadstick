<script lang="ts">
	import { ImageSlider } from '$ui';
	import { pageTracker } from '$store';
	import { itemCollection } from '$stores/Menu';
	import { Grid, Item } from '$ui/menu';
	import { CartIcon } from '$icons';
	import { images } from '$data/SliderImages';

	$pageTracker = { name: 'Menu', url: '/menu' };

	export const callback = (event: Event) => {
		const e = event.target as Element;
		console.log(e.getAttribute('data-id'));
		return false;
	};
</script>

<div class="md:container md:mx-auto py-16 px-4">
	<ImageSlider {images} />

	<Grid paddingYClass="py-16 sm:py-20">
		{#each $itemCollection as item}
			<Item
				id={String(item.id)}
				name={item.name}
				price={item.price}
				image={item.imageurl}
				on:click={callback}
			>
				<CartIcon /> Add to Cart
			</Item>
		{/each}
	</Grid>
</div>
