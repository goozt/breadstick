<script lang="ts">
	import { SpeedDialButton } from 'flowbite-svelte';
	export let name: string,
		url: string | undefined = undefined;
</script>

<a class="speed-dial-link rounded-full" href={url} on:click>
	<SpeedDialButton {name} tooltip="left">
		<slot />
	</SpeedDialButton>
</a>
