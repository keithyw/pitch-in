'use client'

import { useMemo } from 'react'
import { useRouter } from 'next/navigation'
import { ListLayout } from '@pitch-in/shared/components'
import { useDataTableController, useDeleteRecord } from '@pitch-in/shared/hooks'
import { CREATE_PERMISSIONS_URL, PERMISSIONS_URL } from '@/lib'
import { Permission, TableColumn, TableRowAction } from '@pitch-in/shared/types'
import { PermissionAPI } from '@/lib/clients/api'

const PERMISSION_COLUMNS: TableColumn<Permission>[] = [
	{
		header: 'ID',
		accessor: 'id',
		sortable: true,
	},
	{
		header: 'Code',
		accessor: 'code',
		sortable: true,
	},
	{
		header: 'Display Name',
		accessor: 'display_name',
		sortable: true,
	},
]

const PermissionsPage = () => {
	const router = useRouter()

	const tableController = useDataTableController({
		initialSortField: 'code',
		fetchData: PermissionAPI.fetch,
	})

	const deleteController = useDeleteRecord<Permission>({
		deleteFn: PermissionAPI.delete,
		onSuccess: tableController.loadData,
		itemNameProp: 'display_name',
	})

	const columns = useMemo(() => PERMISSION_COLUMNS, [])

	const actions: TableRowAction<Permission>[] = [
		{
			label: 'Details',
			onClick: (p) => {
				router.push(`${PERMISSIONS_URL}/${p.id}`)
			},
			actionType: 'view',
			requiredPermission: '',
		},
		{
			label: 'Edit',
			onClick: (p) => {
				router.push(`${PERMISSIONS_URL}/${p.id}/edit`)
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
			title='Permissions'
			listPermission=''
			createPermission=''
			createUrl={CREATE_PERMISSIONS_URL}
			createText='Create New Permission'
			data={tableController.data}
			columns={columns}
			actions={actions}
			isLoading={tableController.isLoading}
			tableController={tableController}
			deleteController={deleteController}
			deleteTitle='Confirm Delete Permission'
			deleteMessage={(p) =>
				`Are you sure you want to delete ${p?.display_name}?`
			}
		/>
	)
}

export default PermissionsPage
