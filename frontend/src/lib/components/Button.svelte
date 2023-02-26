<script lang="ts">
	export let href: string | undefined = undefined;
	export let id: string | undefined = undefined;

	interface Colors {
		text: string;
		bg: string;
	}

	interface Actions {
		normal: Colors;
		hover: Colors;
		focus: string | undefined;
	}

	interface Themes {
		light: Actions;
		dark: Actions;
	}

	function compile(themes: Themes): string {
		let output: string[] = [
			'text-' + themes.light.normal.text,
			'bg-' + themes.light.normal.bg,
			'hover:text-' + themes.light.hover.text,
			'hover:bg-' + themes.light.hover.bg,
			'focus:ring-' + themes.light.focus,
			'dark:text-' + themes.dark.normal.text,
			'dark:bg-' + themes.dark.normal.bg,
			'dark:hover:text-' + themes.dark.hover.text,
			'dark:hover:bg-' + themes.dark.hover.bg,
			'dark:focus:ring-' + themes.dark.focus
		];
		return output.join(' ');
	}

	const decorations =
		'transition-colors cursor-pointer text-center font-medium focus:ring-4 focus:outline-none inline-flex items-center justify-center px-5 py-2.5 text-sm rounded-full';
	const theme = compile({
		light: {
			normal: {
				text: 'gray-100',
				bg: 'yellow-400'
			},
			hover: {
				text: 'white',
				bg: 'yellow-500'
			},
			focus: 'yellow-300'
		},
		dark: {
			normal: {
				text: 'gray-700',
				bg: 'gray-200'
			},
			hover: {
				text: 'gray-900',
				bg: 'yellow-300'
			},
			focus: 'yellow-500'
		}
	});
	$: classes = [decorations, theme].join(' ');
</script>

<a type="button" data-id={id} class={classes} {href} on:click>
	<slot />
</a>
