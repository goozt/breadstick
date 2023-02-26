<script lang="ts">
	export let paddingXClass = 'px-4 sm:px-6 lg:px-8';
	export let paddingYClass = 'py-1 sm:py-2';
	export let widthClass = 'max-w-2xl lg:max-w-7xl';
	const divClass = `mx-auto ${widthClass} ${paddingYClass} ${paddingXClass}`;
</script>

<div class={divClass}>
	<h2 class="sr-only">Menu</h2>
	<div
		class="grid grid-cols-1 gap-y-10 gap-x-6 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 xl:gap-x-8"
	>
		<slot />
	</div>
</div>
