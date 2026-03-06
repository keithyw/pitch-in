import React from 'react'
import {
	RowActionsMenu,
	SearchInput,
	SpinnerSection,
} from '@pitch-in/shared/components'
import { TableColumn, TableRowAction } from '@pitch-in/shared/types'
import { ChevronDownIcon, ChevronUpIcon } from '@heroicons/react/20/solid'

type FilterComponent = React.ReactNode
type FilterSection = React.ReactNode

interface DataTableProps<T> {
	data: T[]
	columns: TableColumn<T>[]
	rowKey: keyof T
	actions?: TableRowAction<T>[]
	emptyMessage?: string
	searchTerm?: string
	onSearch?: (term: string) => void
	searchPlaceholder?: string
	filter?: FilterComponent
	filterSection?: FilterSection

	// pagination
	currentPage?: number
	pageSize?: number
	totalCount?: number
	onPageChange?: (page: number) => void
	onPageSizeChange?: (size: number) => void
	pageSizes?: number[]
	isLoadingRows?: boolean

	// sorting
	onSort?: (field: string) => void
	// onSort?: (field: string) => void
	currentSortField?: string
	currentSortDirection?: 'asc' | 'desc'
}

// eslint-disable-next-line @typescript-eslint/no-explicit-any
function DataTableComponent<T extends Record<string, any>>({
	data,
	columns,
	rowKey,
	actions,
	emptyMessage = 'No data available',
	searchTerm,
	onSearch,
	searchPlaceholder,
	filter,
	filterSection,

	// pagination props
	currentPage = 1,
	pageSize = 10,
	totalCount = 0,
	onPageChange,
	onPageSizeChange,
	pageSizes = [10, 25, 50, 100],
	isLoadingRows = false,

	// sorting props
	onSort,
	currentSortField,
	currentSortDirection,
}: DataTableProps<T>) {
	const totalColumns = columns.length + (actions && actions.length > 0 ? 1 : 0)
	const totalPages = Math.ceil(totalCount / pageSize)

	const handlePreviousPage = () => {
		if (onPageChange && currentPage > 1) {
			onPageChange(currentPage - 1)
		}
	}

	const handleNextPage = () => {
		if (onPageChange && currentPage < totalPages) {
			onPageChange(currentPage + 1)
		}
	}

	const handlePageSizeChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
		if (onPageSizeChange) {
			onPageSizeChange(Number(e.target.value))
			void onPageChange?.(1)
		}
	}

	return (
		<div className='shadow-md sm:rounded-lg'>
			{(onSearch || onPageChange) && (
				<div className='flex items-center justify-between rounded-t-lg border-b bg-white p-4 dark:border-gray-700 dark:bg-gray-800'>
					<div className='flex items-center space-x-4'>
						{onSearch && (
							<SearchInput
								value={searchTerm || ''}
								onChange={onSearch}
								placeholder={searchPlaceholder}
								className='max-w-xs'
							/>
						)}
						{filter && <div className='relative'>{filter}</div>}
						{filterSection && <div className='relative'>{filterSection}</div>}
					</div>
					{onPageChange && (
						<div className='flex items-center space-x-4'>
							<span className='text-sm text-gray-700 dark:text-gray-300'>
								Items per page:
							</span>
							<select
								value={pageSize}
								onChange={handlePageSizeChange}
								className='block w-20 rounded-md border border-gray-300 bg-white px-3 py-1 text-sm text-gray-700 shadow-sm focus:border-blue-500 focus:ring-blue-500 focus:outline-none dark:border-gray-600 dark:bg-gray-700 dark:text-white'
							>
								{pageSizes.map((s) => (
									<option key={s} value={s}>
										{s}
									</option>
								))}
							</select>
							<span className='text-sm text-gray-700 dark:text-gray-300'>
								Page {currentPage} of {totalPages}
							</span>
							<button
								onClick={handlePreviousPage}
								disabled={currentPage === 1}
								className='cursor-pointer rounded-md bg-gray-100 px-3 py-1 text-sm font-medium text-blue-600 hover:bg-gray-200 disabled:cursor-not-allowed disabled:opacity-50 dark:bg-gray-700 dark:text-blue-300 dark:hover:bg-gray-600'
							>
								Previous
							</button>
							<button
								onClick={handleNextPage}
								disabled={currentPage === totalPages}
								className='cursor-pointer rounded-md bg-gray-100 px-3 py-1 text-sm font-medium text-blue-600 hover:bg-gray-200 disabled:cursor-not-allowed disabled:opacity-50 dark:bg-gray-700 dark:text-blue-300 dark:hover:bg-gray-600'
							>
								Next
							</button>
						</div>
					)}
				</div>
			)}
			<table className='data:text-gray-400 w-full text-left text-sm text-gray-500'>
				<thead className='bg-gray-50 text-xs text-gray-700 uppercase dark:bg-gray-700 dark:text-gray-400'>
					<tr>
						{columns.map((col, idx) => (
							<th
								scope='col'
								className={`px-6 py-3 ${col.sortable ? 'cursor-pointer' : ''}`}
								key={col.header || idx}
								onClick={() => {
									if (col.sortable && (col.sortField || col.accessor)) {
										if (onSort) {
											onSort(col.sortField || (col.accessor as string))
										}
									}
								}}
							>
								{col.sortable && onSort ? (
									<span className='flex items-center gap-1 focus:outline-none'>
										{col.header}
										{currentSortField ===
											(col.sortField ||
												(col.accessor ? String(col.accessor) : '')) &&
											(currentSortDirection === 'asc' ? (
												<ChevronUpIcon className='h-4 w-4 text-gray-500' />
											) : (
												<ChevronDownIcon className='h-4 w-4 text-gray-500' />
											))}
									</span>
								) : (
									col.header
								)}
							</th>
						))}
						{actions && actions.length > 0 && (
							<th scope='col' className='px-6 py-3 text-right'>
								<span>Actions</span>
							</th>
						)}
					</tr>
				</thead>
				<tbody>
					{isLoadingRows ? (
						<tr>
							<td colSpan={totalColumns} className='py-8 text-center'>
								<SpinnerSection spinnerMessage='Loading data...' />
							</td>
						</tr>
					) : data && data.length > 0 ? (
						data.map((row) => (
							<tr
								key={String(row[rowKey])}
								className='border-b bg-white hover:bg-gray-50 dark:border-gray-700 dark:bg-gray-800 dark:hover:bg-gray-600'
							>
								{columns.map((col, colIndex) => (
									<td
										className='px-6 py-4'
										key={String(row[rowKey]) + '-' + colIndex}
									>
										{col.render
											? col.render(row)
											: col.accessor
												? row[col.accessor as keyof T]
												: null}
									</td>
								))}
								{actions && actions.length > 0 && (
									<td className='px-6 py-4 text-right whitespace-nowrap'>
										<RowActionsMenu actions={actions} row={row} />
									</td>
								)}
							</tr>
						))
					) : (
						<tr>
							<td
								colSpan={totalColumns}
								className='py-8 text-center text-gray-500'
							>
								{searchTerm && onSearch
									? `No results for "${searchTerm}"`
									: emptyMessage}
							</td>
						</tr>
					)}
				</tbody>
			</table>
		</div>
	)
}
export const DataTable = React.memo(
	DataTableComponent,
) as typeof DataTableComponent
