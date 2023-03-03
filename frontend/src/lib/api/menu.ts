import { useQuery, useMutation } from '@sveltestack/svelte-query';
import { queryConfig } from '$services/api';
import { fetchAPIv1 } from '$services/fetch';
import type {
	Item,
	MenuResult,
	CreateResult,
} from '$types/menu';

//  Get menu items

export const getMenu = () => {
	return useQuery<MenuResult>('menu', () => fetchAPIv1('GET', '/menu'), queryConfig);
};

// Create new menu items

export const createItem = () => {
	return useMutation(
		async (newItem: Item): Promise<CreateResult> => fetchAPIv1('POST', '/menu', newItem)
	);
};

};

};

};

};
