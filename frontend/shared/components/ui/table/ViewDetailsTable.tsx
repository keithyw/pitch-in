'use client'

import { useRouter } from 'next/navigation'
import { TableColumn } from '@pitch-in/shared/types'

export interface ViewDetailsTableProps<T> {
	data: T[]
	columns: TableColumn<T>[]
	rowKey: keyof T
	getRowHref?: (row: T) => string
}

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export const ViewDetailsTable = <T extends Record<string, any>>({
	data,
	columns,
	rowKey,
	getRowHref,
}: ViewDetailsTableProps<T>) => {
	const router = useRouter()

	if (data.length === 0) {
		return (
			<p className='rounded-md border bg-gray-50 p-4 text-gray-500'>
				No data available
			</p>
		)
	}
	return (
		<div className='flow-root'>
			<div className='-mx-4 -my-2 overflow-x-auto sm:-mx-6 lg:-mx-8'>
				<div className='inline-block min-w-full py-2 align-middle md:px-6 lg:px-8'>
					<table className='min-w-full divide-y divide-gray-300'>
						<thead>
							<tr>
								{columns.map((col, idx) => (
									<th
										key={col.header || idx}
										scope='col'
										className='py-3.5 pr-3 pl-4 text-left text-sm font-semibold text-gray-900 sm:pl-0'
									>
										{col.header}
									</th>
								))}
							</tr>
						</thead>
						<tbody className='divide-y divide-gray-200 bg-white'>
							{data.map((row) => (
								<tr
									key={String(row[rowKey])}
									onClick={() => {
										if (getRowHref) {
											router.push(getRowHref(row))
										}
									}}
									className={
										getRowHref
											? 'cursor-pointer transition-colors hover:bg-gray-50'
											: ''
									}
								>
									{columns.map((col, colIndex) => (
										<td
											key={String(row[rowKey]) + '-' + colIndex}
											className='py-4 pr-3 pl-4 text-sm font-medium whitespace-nowrap text-gray-900 sm:pl-0'
										>
											{col.render
												? col.render(row)
												: col.accessor
													? row[col.accessor as keyof T]
													: null}
										</td>
									))}
								</tr>
							))}
						</tbody>
					</table>
				</div>
			</div>
		</div>
	)
}
