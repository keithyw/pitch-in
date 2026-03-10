'use client'

import { useMemo } from 'react'
import { useRouter } from 'next/navigation'
import { ListLayout } from '@pitch-in/shared/components'
import { useDataTableController, useDeleteRecord } from '@pitch-in/shared/hooks'
import { CREATE_ROLES_URL, ROLES_URL } from '@/lib'
import { Role, TableColumn, TableRowAction } from '@pitch-in/shared/types'
import { RoleAPI } from '@/lib/clients/api'

const COLUMNS: TableColumn<Role>[] = [
	{
		header: 'ID',
		accessor: 'id',
		sortable: true,
	},
	{
		header: 'Name',
		accessor: 'name',
		sortable: true,
	},
]

const RolePage = () => {
	const router = useRouter()

	const tableController = useDataTableController({
		initialSortField: 'name',
		fetchData: RoleAPI.fetch,
	})

	const deleteController = useDeleteRecord<Role>({
		deleteFn: RoleAPI.delete,
		onSuccess: tableController.loadData,
		itemNameProp: 'name',
	})

	const cols = useMemo(() => COLUMNS, [])

	const actions: TableRowAction<Role>[] = [
		{
			label: 'Details',
			onClick: (r) => {
				router.push(`${ROLES_URL}/${r.id}`)
			},
			actionType: 'view',
			requiredPermission: '',
		},
		{
			label: 'Edit',
			onClick: (r) => {
				router.push(`${ROLES_URL}/${r.id}/edit`)
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
			title='Roles'
			listPermission=''
			createPermission=''
			createUrl={CREATE_ROLES_URL}
			createText='Create New Role'
			data={tableController.data}
			columns={cols}
			actions={actions}
			isLoading={tableController.isLoading}
			tableController={tableController}
			deleteController={deleteController}
			deleteTitle='Confirm Delete Role'
			deleteMessage={(r) => `Are you sure you want to delete ${r?.name}?`}
		/>
	)
}

export default RolePage
