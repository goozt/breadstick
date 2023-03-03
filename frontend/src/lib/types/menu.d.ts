export interface Item {
	id?: string;
	name: string;
	category: string;
	price: number;
}

interface ErrorResult {
	error?: {
		code: number;
		message: string;
		details?: string[];
	};
}

export interface MenuResult extends ErrorResult {
	data?: {
		items: Item[];
	};
}

export interface CreateResult extends ErrorResult {
	data?: Item;
}
