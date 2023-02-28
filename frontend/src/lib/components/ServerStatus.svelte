<script lang="ts">
	import { Tooltip } from 'flowbite-svelte';
	import { useQuery } from '@sveltestack/svelte-query';
	import { api } from '$lib/tools';

	type Query = {
		data: {
			status: string;
		};
	};

	const fetchHealth = async () => {
		return fetch(api.server + '/health', {
			mode: 'cors',
			headers: {
				Accept: 'application/json',
				Authorization: 'Bearer ' + api.token,
				'Access-Control-Allow-Origin': '*'
			}
		})
			.then((res) => res.json())
			.catch(() => {
				return { status: 'error' };
			});
	};

	const result = useQuery<Query>('status', fetchHealth, {
		enabled: api.token != undefined && api.server != undefined,
		refetchInterval: 5000,
		refetchIntervalInBackground: true
	});

	$: status = $result.isSuccess && $result.data.data.status == 'ok' && $result.failureCount == 0;
</script>

<span
	id="placement-left"
	class="text-gray-500 dark:text-gray-400 hover:bg-gray-100 dark:hover:bg-gray-700 rounded-lg text-sm p-2 ml-2 cursor-pointer"
>
	<slot>
		{#if status}
			<svg
				fill="currentColor"
				class="w-5 h-5 server-up"
				viewBox="0 0 32 26"
				enable-background="new 0 0 32 26"
				id="Glyph"
				version="1.1"
				xml:space="preserve"
				xmlns="http://www.w3.org/2000/svg"
				xmlns:xlink="http://www.w3.org/1999/xlink"
				><path
					d="M24.972,12.288C24.608,7.657,20.723,4,16,4c-4.04,0-7.508,2.624-8.628,6.451C4.181,11.559,2,14.583,2,18  c0,4.411,3.589,8,8,8h13c3.859,0,7-3.14,7-7C30,15.851,27.93,13.148,24.972,12.288z M20.707,14.707l-5,5  C15.512,19.902,15.256,20,15,20s-0.512-0.098-0.707-0.293l-3-3c-0.391-0.391-0.391-1.023,0-1.414s1.023-0.391,1.414,0L15,17.586  l4.293-4.293c0.391-0.391,1.023-0.391,1.414,0S21.098,14.316,20.707,14.707z"
					id="XMLID_284_"
				/></svg
			>
		{:else}
			<svg
				fill="currentColor"
				class="w-5 h-5 server-down"
				viewBox="0 0 32 26"
				enable-background="new 0 0 32 26"
				id="Glyph"
				version="1.1"
				xml:space="preserve"
				xmlns="http://www.w3.org/2000/svg"
				xmlns:xlink="http://www.w3.org/1999/xlink"
				><path
					d="M24.972,12.288C24.608,7.657,20.723,4,16,4c-4.04,0-7.508,2.624-8.627,6.451C4.181,11.559,2,14.583,2,18  c0,4.411,3.589,8,8,8h13c3.86,0,7-3.14,7-7C30,15.851,27.93,13.148,24.972,12.288z M19.707,19.293c0.391,0.391,0.391,1.023,0,1.414  s-1.023,0.391-1.414,0L16,18.414l-2.293,2.293c-0.391,0.391-1.023,0.391-1.414,0s-0.391-1.023,0-1.414L14.586,17l-2.293-2.293  c-0.391-0.391-0.391-1.023,0-1.414s1.023-0.391,1.414,0L16,15.586l2.293-2.293c0.391-0.391,1.023-0.391,1.414,0s0.391,1.023,0,1.414  L17.414,17L19.707,19.293z"
					id="XMLID_278_"
				/></svg
			>
		{/if}
	</slot>
</span>
<Tooltip color={status ? 'green' : 'red'} placement="left" trigger="click">
	Server {status ? 'connected' : 'down'}
</Tooltip>
