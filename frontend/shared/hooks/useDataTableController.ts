'use client'

import { useCallback, useEffect, useState } from 'react'
import { DEFAULT_PAGE_SIZE } from '@pitch-in/shared/constants'
import { FetchParams, FilterParams, ListResponse } from '@pitch-in/shared/types'

interface UseDataTableControllerProps<T> {
	defaultPageSize?: number
	initialFilters?: FilterParams
	initialSortField?: string
	initSortDirection?: 'asc' | 'desc'
	fetchData: (params: FetchParams) => Promise<ListResponse<T>>
}

export function useDataTableController<T>({
	defaultPageSize = DEFAULT_PAGE_SIZE,
	initialFilters = {},
	initialSortField,
	initSortDirection = 'asc',
	fetchData,
}: UseDataTableControllerProps<T>) {
	const [searchTerm, setSearchTerm] = useState('')
	const [currentPage, setCurrentPage] = useState(1)
	const [pageSize, setPageSize] = useState(defaultPageSize)
	const [filters, setFilters] = useState<FilterParams>(initialFilters)
	const [totalCount, setTotalCount] = useState(0)
	const [sortField, setSortField] = useState<string | undefined>(
		initialSortField,
	)
	const [sortDirection, setSortDirection] = useState<
		'asc' | 'desc' | undefined
	>(initSortDirection)
	const [data, setData] = useState<T[]>([])
	const [isLoading, setIsLoading] = useState<boolean>(false)

	const loadData = useCallback(async () => {
		setIsLoading(true)
		try {
			let orderBy: string | undefined
			if (sortField && sortDirection) {
				orderBy = sortDirection === 'asc' ? sortField : `-${sortField}`
			}
			const res = await fetchData({
				page: currentPage,
				pageSize,
				searchTerm,
				ordering: orderBy,
				filters,
			})
			if (res) {
				setData(res.results)
				setTotalCount(res.count)
			}
		} catch (e: unknown) {
		} finally {
			setIsLoading(false)
		}
	}, [
		currentPage,
		pageSize,
		searchTerm,
		sortField,
		sortDirection,
		fetchData,
		filters,
	])

	useEffect(() => {
		void loadData()
	}, [loadData])

	const handleSearch = useCallback((term: string) => {
		setSearchTerm(term)
		setCurrentPage(1)
	}, [])

	const handlePageChange = useCallback((page: number) => {
		setCurrentPage(page)
	}, [])

	const handlePageSizeChange = useCallback((size: number) => {
		setPageSize(size)
		setCurrentPage(1)
	}, [])

	const handleSort = useCallback(
		(f: string) => {
			let newDirection: 'asc' | 'desc' = 'asc'
			if (sortField === f) {
				newDirection = sortDirection === 'asc' ? 'desc' : 'desc'
			}
			setSortField(f)
			setSortDirection(newDirection)
			setCurrentPage(1)
		},
		[sortField, sortDirection],
	)

	const handleFilters = useCallback(
		(f: FilterParams) => {
			if (JSON.stringify(f) !== JSON.stringify(filters)) {
				setFilters(f)
				setCurrentPage(1)
			}
		},
		[filters],
	)

	return {
		data,
		isLoading,
		searchTerm,
		totalCount,
		currentPage,
		pageSize,
		sortField,
		sortDirection,
		handleSearch,
		handlePageChange,
		handlePageSizeChange,
		handleSort,
		handleFilters,
		loadData,
	}
}
