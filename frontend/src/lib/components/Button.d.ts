type Colors = {
	text: string;
	bg: string;
};

type Actions = {
	normal: Colors;
	hover: Colors;
	focus: string;
};

export type ThemeClass = {
	light: Actions;
	dark: Actions;
};
