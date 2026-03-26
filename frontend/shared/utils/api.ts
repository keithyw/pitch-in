import { FetchParams, FilterField } from '@pitch-in/shared/types'

export const prepareQueryParams = (
	params: FetchParams,
): Record<string, any> => {
	const { page, pageSize, searchTerm, ordering, filters, fields } = params

	// 1. Map base pagination and search
	const query: Record<string, any> = {
		limit: pageSize,
		offset: (page as number) - 1,
		sort: ordering
			? ordering.startsWith('-')
				? `${ordering.slice(1)}.desc`
				: `${ordering}.asc`
			: undefined,
		...(filters || {}), // 2. Spread additional filters
	}

	if (fields && fields.length > 0) {
		for (const f of fields as FilterField[]) {
			if (!f.operator) {
				query[f.field] = f.value
				continue
			}

			const filter = `${f.field}${f.operator[0]}`
			query[filter] = f.value
		}
	}

	// 3. Clean up undefined keys to keep the URL clean
	Object.keys(query).forEach((key) => {
		if (query[key] === undefined || query[key] === null) {
			delete query[key]
		}
	})

	return query
}
