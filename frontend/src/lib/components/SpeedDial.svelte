<script lang="ts">
	import { SpeedDial } from 'flowbite-svelte';
</script>

<SpeedDial trigger="click" defaultClass="speed-dial fixed right-6 bottom-6 z-40">
	<slot />
</SpeedDial>

<style lang="scss">
	:global.speed-dial > button {
		--tw-bg-opacity: 1;
		background-color: rgb(227 160 8 / var(--tw-bg-opacity));
	}
	:global.speed-dial > button:hover {
		--tw-bg-opacity: 1;
		background-color: rgb(227 160 8 / var(--tw-bg-opacity)) !important;
	}
	:global.speed-dial > button:focus {
		--tw-ring-opacity: 1;
		--tw-ring-color: rgb(227 160 8 / var(--tw-ring-opacity));
	}
	:global.dark .speed-dial > button {
		--tw-text-opacity: 1;
		color: rgb(0 0 0 / var(--tw-text-opacity));
		--tw-bg-opacity: 1;
		background-color: rgb(247 183 37 / var(--tw-bg-opacity)) !important;
	}
	:global.dark .speed-dial > button:hover {
		--tw-text-opacity: 1;
		color: rgb(0 0 0 / var(--tw-text-opacity));
		--tw-bg-opacity: 1;
		background-color: rgb(247 183 37 / var(--tw-bg-opacity)) !important;
	}
	:global.dark .speed-dial > button:focus {
		--tw-ring-opacity: 1;
		--tw-ring-color: rgb(196 137 3 / var(--tw-ring-opacity));
	}
</style>
