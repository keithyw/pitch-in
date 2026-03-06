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
}

export interface PaginationParams {
	page?: number
	page_size?: number
	search?: string
	ordering?: string
}
