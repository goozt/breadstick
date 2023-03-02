export declare type ImageType = {
	name: string;
	imgurl: string;
};

export declare type MenuItem = {
	id: string;
	name: string;
	price: number;
	imageurl: string;
};

export declare type ToastItem = {
	id: string;
	type: ToastItemType;
	message: string;
};

export declare type ToastItemType =
	| 'form'
	| 'none'
	| 'default'
	| 'gray'
	| 'red'
	| 'yellow'
	| 'green'
	| 'indigo'
	| 'purple'
	| 'pink'
	| 'blue'
	| 'light'
	| 'dark'
	| 'dropdown'
	| 'navbar'
	| 'navbarUl'
	| undefined;
