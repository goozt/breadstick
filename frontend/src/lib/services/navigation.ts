export const getHeading = (path: string): string => {
	const tokens = path.split('/').filter((t) => t !== '');
	const heading = tokens[tokens.length - 1];
	return heading.charAt(0).toUpperCase() + heading.slice(1);
};
