'use client'

import { useMemo } from 'react'
import { useRouter } from 'next/navigation'
import { ListLayout } from '@pitch-in/shared/components'
import { ASSETS_READ, ASSETS_WRITE } from '@pitch-in/shared/constants'
import { useDataTableController, useDeleteRecord } from '@pitch-in/shared/hooks'
import { Asset, TableColumn, TableRowAction } from '@pitch-in/shared/types'
import { getFileIcon } from '@pitch-in/shared/utils'
import { ASSETS_URL, CREATE_ASSET_URL, ICON_CLASS } from '@/lib'
import { AssetAPI } from '@/lib/clients/api'

const COLS: TableColumn<Asset>[] = [
	{
		header: 'ID',
		accessor: 'id',
		sortable: true,
	},
	{
		header: 'Mime Type',
		render: (a: Asset) => {
			const Icon = getFileIcon(a.mime_type)
			return (
				<div className='flex items-center space-x-2'>
					<Icon className={ICON_CLASS} />
					<span className='max-w-[100px] truncate text-xs text-gray-500'>
						{a.mime_type.split('/')[1].toUpperCase()}
					</span>
				</div>
			)
		},
		sortable: false,
	},
	{
		header: 'Object Key',
		accessor: 'object_key',
		sortable: true,
	},
]

const AssetsPage = () => {
	const router = useRouter()

	const tableController = useDataTableController({
		initialSortField: 'object_key',
		fetchData: AssetAPI.fetch,
	})

	const deleteController = useDeleteRecord<Asset>({
		deleteFn: AssetAPI.delete,
		onSuccess: tableController.loadData,
		itemNameProp: 'object_key',
	})

	const cols = useMemo(() => COLS, [])

	const actions: TableRowAction<Asset>[] = [
		{
			label: 'Details',
			onClick: (a) => {
				router.push(`${ASSETS_URL}/${a.id}`)
			},
			actionType: 'view',
			requiredPermission: ASSETS_READ,
		},
		{
			label: 'Edit',
			onClick: (a) => {
				router.push(`${ASSETS_URL}/${a.id}/edit`)
			},
			actionType: 'edit',
			requiredPermission: ASSETS_WRITE,
		},
		{
			label: 'Delete',
			onClick: deleteController.openDeleteModal,
			actionType: 'delete',
			requiredPermission: ASSETS_WRITE,
		},
	]

	return (
		<ListLayout
			title='Assets'
			listPermission={ASSETS_READ}
			createPermission={ASSETS_WRITE}
			createUrl={CREATE_ASSET_URL}
			createText='Create New Asset'
			data={tableController.data}
			columns={cols}
			actions={actions}
			isLoading={tableController.isLoading}
			tableController={tableController}
			deleteController={deleteController}
			deleteTitle='Confirm Asset Deletion'
			deleteMessage={(a) => `Are you sure you want to delete ${a?.object_key}?`}
		/>
	)
}

export default AssetsPage
