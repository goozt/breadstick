<script lang="ts">
	import { pageTracker } from '$stores/navigation';
	import { addMenuItem, removeMenuItem, itemCollection } from '$stores/menu';
	import { Grid, Item } from '$ui/menu';
	import { SpeedDial, SpeedDialButton } from '$ui';
	$pageTracker = { name: 'Dashboard', url: 'dashboard' };
</script>

<div class="md:container md:mx-auto py-16 px-4">
	<Grid>
		{#each $itemCollection as item}
			<Item
				id={String(item.id)}
				name={item.name}
				price={item.price}
				image={item.imageurl}
				on:click={removeMenuItem}>Remove</Item
			>
		{/each}
	</Grid>
</div>

<SpeedDial>
	<SpeedDialButton name="Add&nbsp;Menu&nbsp;Item" on:click={addMenuItem}>
		<svg
			width="75%"
			height="75%"
			aria-hidden="true"
			fill="currentColor"
			viewBox="0 0 20 20"
			xmlns="http://www.w3.org/2000/svg"
		>
			<path
				clip-rule="evenodd"
				d="M4.5 2A1.5 1.5 0 003 3.5v13A1.5 1.5 0 004.5 18h11a1.5 1.5 0 001.5-1.5V7.621a1.5 1.5 0 00-.44-1.06l-4.12-4.122A1.5 1.5 0 0011.378 2H4.5zM10 8a.75.75 0 01.75.75v1.5h1.5a.75.75 0 010 1.5h-1.5v1.5a.75.75 0 01-1.5 0v-1.5h-1.5a.75.75 0 010-1.5h1.5v-1.5A.75.75 0 0110 8z"
				fill-rule="evenodd"
			/>
		</svg>
	</SpeedDialButton>
	<SpeedDialButton name="Settings" url="/dashboard/settings">
		<svg
			width="60%"
			height="60%"
			aria-hidden="true"
			fill="currentColor"
			viewBox="0 0 20 20"
			xmlns="http://www.w3.org/2000/svg"
		>
			<path
				clip-rule="evenodd"
				d="M14.5 10a4.5 4.5 0 004.284-5.882c-.105-.324-.51-.391-.752-.15L15.34 6.66a.454.454 0 01-.493.11 3.01 3.01 0 01-1.618-1.616.455.455 0 01.11-.494l2.694-2.692c.24-.241.174-.647-.15-.752a4.5 4.5 0 00-5.873 4.575c.055.873-.128 1.808-.8 2.368l-7.23 6.024a2.724 2.724 0 103.837 3.837l6.024-7.23c.56-.672 1.495-.855 2.368-.8.096.007.193.01.291.01zM5 16a1 1 0 11-2 0 1 1 0 012 0z"
				fill-rule="evenodd"
			/>
			<path
				d="M14.5 11.5c.173 0 .345-.007.514-.022l3.754 3.754a2.5 2.5 0 01-3.536 3.536l-4.41-4.41 2.172-2.607c.052-.063.147-.138.342-.196.202-.06.469-.087.777-.067.128.008.257.012.387.012zM6 4.586l2.33 2.33a.452.452 0 01-.08.09L6.8 8.214 4.586 6H3.309a.5.5 0 01-.447-.276l-1.7-3.402a.5.5 0 01.093-.577l.49-.49a.5.5 0 01.577-.094l3.402 1.7A.5.5 0 016 3.31v1.277z"
			/>
		</svg>
	</SpeedDialButton>
</SpeedDial>
