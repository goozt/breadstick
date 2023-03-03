export interface Item {
	id?: string;
	name: string;
	category: string;
	price: number;
}

export interface EditableItem {
	name?: string;
	category?: string;
	price?: number;
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

export interface ItemResult extends ErrorResult {
	data?: Item;
}

export interface CreateResult extends ErrorResult {
	data?: Item;
}

export interface DeleteResult extends ErrorResult {
	data?: {
		deleted_item: string;
	};
}

export interface DeleteMenuResult extends ErrorResult {
	data?: {
		detelted_menu: boolean;
	};
}
