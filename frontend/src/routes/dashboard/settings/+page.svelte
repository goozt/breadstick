<script lang="ts">
	import { Tabs, TabItem, Timeline, TimelineItem, Button } from 'flowbite-svelte';
</script>

<div class="md:container md:mx-auto py-16 px-4">
	<Tabs>
		<TabItem open>
			<span slot="title">General</span>
			<Timeline>
				<TimelineItem title="Application UI code in Tailwind CSS" date="February 2022">
					<p class="mb-4 text-base font-normal text-gray-500 dark:text-gray-400">
						Get access to over 20+ pages including a dashboard layout, charts, kanban board,
						calendar, and pre-order E-commerce & Marketing pages.
					</p>
					<Button color="alternative"
						>Learn more
						<svg
							class="ml-2 w-3 h-3"
							fill="currentColor"
							viewBox="0 0 20 20"
							xmlns="http://www.w3.org/2000/svg"
							><path
								fill-rule="evenodd"
								d="M12.293 5.293a1 1 0 011.414 0l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414-1.414L14.586 11H3a1 1 0 110-2h11.586l-2.293-2.293a1 1 0 010-1.414z"
								clip-rule="evenodd"
							/></svg
						>
					</Button>
				</TimelineItem>
				<TimelineItem title="Application UI code in Tailwind CSS" date="March 2022">
					<p class="text-base font-normal text-gray-500 dark:text-gray-400">
						All of the pages and components are first designed in Figma and we keep a parity between
						the two versions even as we update the project.
					</p>
				</TimelineItem>
				<TimelineItem title="Application UI code in Tailwind CSS" date="April 2022">
					<p class="text-base font-normal text-gray-500 dark:text-gray-400">
						Get started with dozens of web components and interactive elements built on top of
						Tailwind CSS.
					</p>
				</TimelineItem>
			</Timeline>
		</TabItem>
		<TabItem>
			<span slot="title">Users</span>
			<p class="text-sm text-gray-500 dark:text-gray-400">
				<b>Users:</b> Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor
				incididunt ut labore et dolore magna aliqua.
			</p>
		</TabItem>
		<TabItem>
			<span slot="title">Dashboard</span>
			<p class="text-sm text-gray-500 dark:text-gray-400">
				<b>Dashboard:</b> Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor
				incididunt ut labore et dolore magna aliqua.
			</p>
		</TabItem>
		<TabItem>
			<span slot="title">Advance</span>
			<p class="text-sm text-gray-500 dark:text-gray-400">
				<b>Settings:</b> Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor
				incididunt ut labore et dolore magna aliqua.
			</p>
		</TabItem>
	</Tabs>
</div>
