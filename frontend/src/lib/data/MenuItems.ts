import type { MenuItem } from '$ui-types';
import { newUUID } from '$services/uuid';

export const defaultItems: MenuItem[] = [
	{
		id: newUUID(),
		name: 'Item 1',
		price: 45,
		imageurl: '/images/image1.png'
	},
	{
		id: newUUID(),
		name: 'Item 2',
		price: 63,
		imageurl: '/images/image2.png'
	},
	{
		id: newUUID(),
		name: 'Item 3',
		price: 52,
		imageurl: '/images/image3.png'
	},
	{
		id: newUUID(),
		name: 'Item 4',
		price: 95,
		imageurl: '/images/image4.png'
	},
	{
		id: newUUID(),
		name: 'Item 5',
		price: 74,
		imageurl: '/images/image5.png'
	},
	{
		id: newUUID(),
		name: 'Item 6',
		price: 39,
		imageurl: '/images/image6.png'
	}
];
