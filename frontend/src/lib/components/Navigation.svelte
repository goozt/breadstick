<script lang="ts">
	import { DarkMode, Navbar, NavHamburger, NavLi, NavUl } from 'flowbite-svelte';
	import { ServerStatus } from '$ui';
	import { CartIcon } from '$icons';
	import { getHeading } from '$services/navigation';

	export let page: string;

	export let navMenu: string[];
	let navClass = 'bg-white border-gray-200 px-4 lg:px-6 py-2.5 dark:bg-gray-800';
	let navDivClass = 'flex flex-wrap justify-between items-center mx-auto max-w-screen-xl';
	let btnClass =
		'text-gray-500 dark:text-gray-400 hover:bg-gray-100 dark:hover:bg-gray-700 focus:outline-none focus:ring-4 focus:ring-gray-200 dark:focus:ring-gray-700 rounded-lg text-sm p-2 ml-2';

	let activeClass =
		'text-gray-900 hover:bg-yellow-300 lg:hover:bg-transparent dark:text-yellow-400 dark:hover:bg-gray-700 lg:dark:hover:bg-transparent transition-transform lg:scale-105';
	let nonActiveClass =
		'text-gray-500 hover:text-gray-700 hover:bg-yellow-300 lg:hover:bg-transparent dark:text-gray-400 dark:hover:text-white dark:hover:bg-gray-700 lg:dark:hover:bg-transparent';
</script>

<Navbar let:hidden let:toggle fluid={false} {navClass} {navDivClass}>
	<slot />
	<!-- Right Bar -->
	<div class="flex items-center lg:order-2 text-r">
		<ServerStatus />
		<DarkMode {btnClass} />
		<button class={`${btnClass} bg-yellow-300 hover:bg-yellow-400 dark:hover:bg-yellow-500`}>
			<CartIcon classes="w-6 h-6 text-gray-700 dark:text-black" />
		</button>
		<NavHamburger
			on:click={toggle}
			btnClass="inline-flex items-center p-2 ml-1 text-sm text-gray-500 rounded-lg lg:hidden hover:bg-gray-100 focus:outline-none focus:ring-4 focus:ring-gray-200 dark:text-gray-400 dark:hover:bg-gray-700 dark:focus:ring-gray-600 ml-2"
		/>
	</div>

	<!-- Menu -->
	<NavUl
		{hidden}
		divClass="nav-bar justify-between items-center w-full lg:flex lg:w-auto lg:order-1"
		ulClass="nav-menu flex flex-col mt-4 font-medium lg:flex-row lg:space-x-8 lg:mt-0"
	>
		{#each navMenu as navItem}
			<NavLi href={navItem} active={page == navItem} {activeClass} {nonActiveClass}
				>{getHeading(navItem)}</NavLi
			>
		{/each}
	</NavUl>
</Navbar>
