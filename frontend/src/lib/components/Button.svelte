<script lang="ts">
	import classNames from 'classnames';
	export let href: string | undefined = undefined;
	export let id: string | undefined = undefined;

	const theme: ThemeClass = {
		light: {
			normal: {
				text: 'text-gray-100',
				bg: 'bg-yellow-400'
			},
			hover: {
				text: 'hover:text-white',
				bg: 'hover:bg-yellow-500'
			},
			focus: 'focus:ring-yellow-300'
		},
		dark: {
			normal: {
				text: 'dark:text-gray-700',
				bg: 'dark:bg-gray-200'
			},
			hover: {
				text: 'dark:hover:text-gray-900',
				bg: 'dark:hover:bg-yellow-300'
			},
			focus: 'dark:focus:ring-yellow-500'
		}
	};

	function compile(themes: ThemeClass): string {
		let output: string[] = [
			themes.light.normal.text,
			themes.light.normal.bg,
			themes.light.hover.text,
			themes.light.hover.bg,
			themes.light.focus,
			themes.dark.normal.text,
			themes.dark.normal.bg,
			themes.dark.hover.text,
			themes.dark.hover.bg,
			themes.dark.focus
		];
		return output.join(' ');
	}

	const decorations =
		'transition-colors cursor-pointer text-center font-medium focus:ring-4 focus:outline-none inline-flex items-center justify-center px-5 py-2.5 text-sm rounded-full';
	$: classes = classNames(decorations, compile(theme));
</script>

<a type="button" data-id={id} class={classes} {href} on:click>
	<slot />
</a>
