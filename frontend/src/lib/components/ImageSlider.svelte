<script lang="ts">
	import { Carousel } from 'flowbite-svelte';
	import type { ImageType } from '$ui-types';

	export let images: ImageType[];

	let showThumbs = false;
	let showCaptions = false;
	let duration = 3000;
	let divClass = 'overflow-hidden relative rounded-lg sm:h-64 xl:h-80 2xl:h-96';
</script>

<div class="flex content-center max-w-4xl m-auto">
	<Carousel {images} loop {showCaptions} {showThumbs} {duration} {divClass} />
</div>

<style lang="scss">
	:global#default-carousel {
		width: 100%;
		img {
			width: 100%;
		}
	}
</style>
