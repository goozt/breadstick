import { useQuery } from '@sveltestack/svelte-query';
import { queryConfig } from '$services/api';
import { fetchServerData } from '$services/fetch';

type QueryType = {
	data: {
		status: string;
	};
};

const fetchData = async () => {
	return fetchServerData('/health').catch(() => {
		return { status: 'error' };
	});
};

export default () => {
	return useQuery<QueryType>('status', fetchData, queryConfig);
};
