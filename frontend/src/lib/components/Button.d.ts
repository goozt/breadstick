interface Colors {
	text: string;
	bg: string;
}

interface Actions {
	normal: Colors;
	hover: Colors;
	focus: string;
}

export interface ThemeClass {
	light: Actions;
	dark: Actions;
}
