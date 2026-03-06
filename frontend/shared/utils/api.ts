import { FetchParams } from '@pitch-in/shared/types'

export const prepareQueryParams = (
	params: FetchParams,
): Record<string, any> => {
	const { page, pageSize, searchTerm, ordering, filters } = params

	// 1. Map base pagination and search
	const query: Record<string, any> = {
		limit: pageSize,
		offset: page,
		sort: ordering
			? ordering.startsWith('-')
				? `${ordering.slice(1)}.desc`
				: `${ordering}.asc`
			: undefined,
		...(filters || {}), // 2. Spread additional filters
	}
	// 3. Clean up undefined keys to keep the URL clean
	Object.keys(query).forEach((key) => {
		if (query[key] === undefined || query[key] === null) {
			delete query[key]
		}
	})

	return query
}
