<script lang="ts">
	import { ImageSlider } from '$ui';
	import { Grid, Item } from '$ui/menu';
	import { CartIcon } from '$icons';
	import { sliderImages } from '$services/slider';
	import menuAPI from '$api/menu';

	export const callback = (event: Event) => {
		const e = event.target as Element;
		console.log(e.getAttribute('data-id'));
		return false;
	};
	let menu = menuAPI.list();

	$: menuItems = $menu.isSuccess ? $menu.data.data?.items ?? [] : [];
</script>

<div class="md:container md:mx-auto py-16 px-4">
	<ImageSlider images={sliderImages} />

	<Grid paddingYClass="py-16 sm:py-20">
		{#each menuItems as item}
			<Item
				id={String(item.id)}
				name={item.name}
				price={item.price}
				image="/images/image1.jpg"
				on:click={callback}
			>
				<CartIcon /> Add to Cart
			</Item>
		{/each}
	</Grid>
</div>
