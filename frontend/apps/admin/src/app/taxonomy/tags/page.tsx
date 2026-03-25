'use client'

import { useMemo } from 'react'
import { useRouter } from 'next/navigation'
import { ListLayout } from '@pitch-in/shared/components'
import { CONTENT_READ, CONTENT_WRITE } from '@pitch-in/shared/constants'
import { useDataTableController, useDeleteRecord } from '@pitch-in/shared/hooks'
import { TableColumn, TableRowAction, Tag } from '@pitch-in/shared/types'
import { CREATE_TAG_URL, TAGS_URL } from '@/lib/constants'
import { TagAPI } from '@/lib/clients/api'

const COLS: TableColumn<Tag>[] = [
	{
		header: 'ID',
		accessor: 'id',
		sortable: true,
	},
	{
		header: 'Tag',
		accessor: 'tag',
		sortable: true,
	},
]

const TagsPage = () => {
	const router = useRouter()

	const tableController = useDataTableController({
		initialSortField: 'tag',
		fetchData: TagAPI.fetch,
	})

	const deleteController = useDeleteRecord<Tag>({
		deleteFn: TagAPI.delete,
		onSuccess: tableController.loadData,
		itemNameProp: 'tag',
	})

	const cols = useMemo(() => COLS, [])

	const actions: TableRowAction<Tag>[] = [
		{
			label: 'Details',
			onClick: (t) => {
				router.push(`${TAGS_URL}/${t.id}`)
			},
			actionType: 'view',
			requiredPermission: CONTENT_READ,
		},
		{
			label: 'Edit',
			onClick: (t) => {
				router.push(`${TAGS_URL}/${t.id}/edit`)
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
			title='Tags'
			listPermission={CONTENT_READ}
			createPermission={CONTENT_WRITE}
			createUrl={CREATE_TAG_URL}
			createText='Create New Tag'
			data={tableController.data}
			columns={cols}
			actions={actions}
			isLoading={tableController.isLoading}
			tableController={tableController}
			deleteController={deleteController}
			deleteTitle='Confirm Delete Tag'
			deleteMessage={(t) => `Are you sure you want to delete ${t?.tag}?`}
		/>
	)
}

export default TagsPage
