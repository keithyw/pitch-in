'use client'

import { useMemo } from 'react'
import { useRouter } from 'next/navigation'
import { ListLayout } from '@pitch-in/shared/components'
import { useDataTableController, useDeleteRecord } from '@pitch-in/shared/hooks'
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
	{
		header: 'Full Name',
		render: (u: User) => `${u.first_name} ${u.last_name}`,
		sortable: true,
		sortField: 'last_name',
	},
	{
		header: 'Is Active',
		render: (u: User) => `${u.is_active ? 'Yes' : 'No'}`,
	},
]

const UsersPage = () => {
	const router = useRouter()

	const tableController = useDataTableController({
		initialSortField: 'username',
		fetchData: UserAPI.fetch,
	})

	const deleteController = useDeleteRecord<User>({
		deleteFn: UserAPI.delete,
		onSuccess: tableController.loadData,
		itemNameProp: 'username',
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
		{
			label: 'Edit',
			onClick: (u) => {
				router.push(`${USERS_URL}/${u.id}/edit`)
			},
			actionType: 'edit',
			requiredPermission: '',
		},
		{
			label: 'Delete',
			onClick: deleteController.openDeleteModal,
			actionType: 'delete',
			requiredPermission: '',
		},
	]

	return (
		<ListLayout
			title='Users'
			listPermission=''
			createPermission=''
			createUrl={CREATE_USERS_URL}
			createText='Create New User'
			data={tableController.data}
			columns={userColumns}
			actions={actions}
			isLoading={tableController.isLoading}
			tableController={tableController}
			deleteController={deleteController}
			deleteTitle='Confirm Delete User'
			deleteMessage={(u) => `Are you sure you want to delete ${u?.username}?`}
		/>
	)
}

export default UsersPage
