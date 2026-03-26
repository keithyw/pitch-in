const OPERATORS = ['>', '<', '<=', '>=', '~=']

export interface FilterField {
	field: string
	operator?: (typeof OPERATORS)[number]
	value: string
}

export type FilterParams = Record<string, string | number | boolean | undefined>
export interface FilterOption {
	key: string
	label: string
}

export interface FetchParams {
	page?: number
	pageSize?: number
	searchTerm?: string
	ordering?: string
	filters?: FilterParams
	fields?: FilterField[]
}

export interface PaginationParams {
	page?: number
	page_size?: number
	search?: string
	ordering?: string
}
