<script lang="ts">
	import '../app.postcss';
	import { NavBrand, FooterLinkGroup, FooterLink } from 'flowbite-svelte';
	import { Logo } from '$icons';
	import { Navigation, Header, Footer, Toast } from '$ui';
	import { toastChannel } from '$services/toast';
	import { QueryClientProvider } from '@sveltestack/svelte-query';
	import { queryClient } from '$services/api';
	import { page } from '$app/stores';

	const navList = ['/menu', '/dashboard'];
	$: currentPage = $page.url.pathname;
</script>

<QueryClientProvider client={queryClient}>
	<main>
		<div class="flex flex-col min-h-screen bg-white dark:bg-gray-900">
			<header class="sticky top-0 z-40 flex-none mx-auto w-full bg-white dark:bg-gray-900 ">
				<div class="absolute top-20 right-4">
					{#each $toastChannel as toastItem}
						<Toast item={toastItem} />
					{/each}
				</div>
				<!-- Navigation Menu -->
				<Navigation navMenu={navList} bind:page={currentPage}>
					<!-- Left logo -->
					<NavBrand href="/">
						<Logo size="32" />
						<span
							class="transition-transform self-center whitespace-nowrap text-xl font-semibold dark:text-white ml-2"
							class:scale-105={currentPage == '/'}
						>
							BreadStick
						</span>
					</NavBrand>
				</Navigation>
				<hr class="bg-gray-200 rounded border-0 dark:bg-gray-700 w-full h-px" />
			</header>
			<!-- Content -->
			<div class="max-w-8xl dark:bg-gray-900">
				{#if currentPage != '/'}
					<Header page={currentPage} />
				{/if}
				<slot />
			</div>
		</div>
	</main>
	<Footer
		name="BreadStick"
		description="We are a group of food artists and enthusiast who create baked food items for helping people who starve in the world."
	>
		<div>
			<h2 class="mb-6 text-sm font-semibold text-gray-400 uppercase dark:text-white">
				HELP & SUPPORT
			</h2>
			<FooterLinkGroup>
				<FooterLink liClass="mb-4" href="https://discord.gg/beadstick">Discord</FooterLink>
				<FooterLink liClass="mb-4" href="https://github.com/goozt/beadstick/issues"
					>GitHub</FooterLink
				>
			</FooterLinkGroup>
		</div>
		<div>
			<h2 class="mb-6 text-sm font-semibold text-gray-400 uppercase dark:text-white">Follow us</h2>
			<FooterLinkGroup>
				<FooterLink liClass="mb-4" href="https://github.com/goozt/beadstick">Gihub</FooterLink>
				<FooterLink liClass="mb-4" href="https://discord.gg/beadstick">Discord</FooterLink>
				<FooterLink liClass="mb-4" href="https://twitter.com/beadstick">Twitter</FooterLink>
			</FooterLinkGroup>
		</div>
	</Footer>
</QueryClientProvider>

<style lang="scss">
	:global.nav-menu a {
		@media (min-width: 768px) {
			padding: 0.5rem;
		}
	}
</style>
