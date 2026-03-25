'use client'

import { useMemo } from 'react'
import { useRouter } from 'next/navigation'
import { ListLayout } from '@pitch-in/shared/components'
import { CONTENT_READ, CONTENT_WRITE } from '@pitch-in/shared/constants'
import { useDataTableController, useDeleteRecord } from '@pitch-in/shared/hooks'
import { TableColumn, TableRowAction, Item } from '@pitch-in/shared/types'
import { CREATE_ITEM_URL, ITEMS_URL } from '@/lib/constants'
import { ItemAPI } from '@/lib/clients/api'

const COLS: TableColumn<Item>[] = [
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

const ItemsPage = () => {
	const router = useRouter()

	const tableController = useDataTableController({
		initialSortField: 'name',
		fetchData: ItemAPI.fetch,
	})

	const deleteController = useDeleteRecord<Item>({
		deleteFn: ItemAPI.delete,
		onSuccess: tableController.loadData,
		itemNameProp: 'name',
	})

	const cols = useMemo(() => COLS, [])

	const actions: TableRowAction<Item>[] = [
		{
			label: 'Details',
			onClick: (t) => {
				router.push(`${ITEMS_URL}/${t.id}`)
			},
			actionType: 'view',
			requiredPermission: CONTENT_READ,
		},
		{
			label: 'Edit',
			onClick: (t) => {
				router.push(`${ITEMS_URL}/${t.id}/edit`)
			},
			actionType: 'edit',
			requiredPermission: CONTENT_WRITE,
		},
		{
			label: 'Delete',
			onClick: deleteController.openDeleteModal,
			actionType: 'delete',
			requiredPermission: CONTENT_WRITE,
		},
	]

	return (
		<ListLayout
			title='Items'
			listPermission={CONTENT_READ}
			createPermission={CONTENT_WRITE}
			createUrl={CREATE_ITEM_URL}
			createText='Create New Item'
			data={tableController.data}
			columns={cols}
			actions={actions}
			isLoading={tableController.isLoading}
			tableController={tableController}
			deleteController={deleteController}
			deleteTitle='Confirm Delete Item'
			deleteMessage={(i) => `Are you sure you want to delete ${i?.name}?`}
		/>
	)
}

export default ItemsPage
