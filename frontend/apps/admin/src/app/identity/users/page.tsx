'use client'

import { useMemo } from 'react'
import { useRouter } from 'next/navigation'
import { CreateItemSection, PageTitle } from '@pitch-in/shared/components'
import { DataTable } from '@pitch-in/shared/components/ui/table'
import { DEFAULT_PAGE_SIZE } from '@pitch-in/shared/constants'
import { useDataTableController } from '@pitch-in/shared/hooks'
import { TableColumn, TableRowAction, User } from '@pitch-in/shared/types'
import { UserAPI } from '@/lib/clients/api'
import { CREATE_USERS_URL, USERS_URL } from '@/lib'

const USER_COLUMNS: TableColumn<User>[] = [
	{
		header: 'ID',
		accessor: 'id',
		sortable: true,
	},
	{
		header: 'Username',
		accessor: 'username',
		sortable: true,
	},
]

const UsersPage = () => {
	const router = useRouter()

	const {
		data: users,
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
		loadData,
	} = useDataTableController({
		initialSortField: 'username',
		defaultPageSize: DEFAULT_PAGE_SIZE,
		fetchData: UserAPI.fetch,
	})

	const userColumns = useMemo(() => USER_COLUMNS, [])

	const actions: TableRowAction<User>[] = [
		{
			label: 'Details',
			onClick: (u) => {
				router.push(`${USERS_URL}/${u.id}`)
			},
			actionType: 'view',
			requiredPermission: '',
		},
	]

	return (
		<>
			<PageTitle>Users</PageTitle>
			<CreateItemSection permission='' href={CREATE_USERS_URL}>
				Create New User
			</CreateItemSection>

			<DataTable
				data={users}
				columns={userColumns}
				rowKey='id'
				actions={actions}
				searchTerm={searchTerm}
				onSearch={handleSearch}
				currentPage={currentPage}
				pageSize={pageSize}
				totalCount={totalCount}
				onPageChange={handlePageChange}
				onPageSizeChange={handlePageSizeChange}
				onSort={handleSort}
				currentSortField={sortField}
				currentSortDirection={sortDirection}
				isLoadingRows={isLoading}
			/>
		</>
	)
}

export default UsersPage
