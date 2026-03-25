'use client'

import { useCallback, useState } from 'react'
import { useParams } from 'next/navigation'
import { DetailsLayout, DetailSectionRow } from '@pitch-in/shared/components'
import { CONTENT_READ, CONTENT_WRITE } from '@pitch-in/shared/constants'
import { useDetailsController } from '@pitch-in/shared'
import { Item } from '@pitch-in/shared/types'
import { failedLoadingError } from '@pitch-in/shared/utils'
import { ITEMS_URL } from '@/lib/constants'
import { ItemAPI } from '@/lib/clients/api'

const ItemDetailsPage = () => {
	const params = useParams()
	const [details, setDetails] = useState<DetailSectionRow[]>([])

	const detailsCallback = useCallback((i: Item) => {
		setDetails([
			{
				label: 'Id',
				value: i.id.toString(),
			},
			{
				label: 'Name',
				value: i.name,
			},
			{
				label: 'Slug',
				value: i.slug,
			},
			{
				label: 'Description',
				value: i.description as string,
			},
		])
	}, [])

	const detailsController = useDetailsController({
		id: parseInt(params.id as string),
		deleteData: ItemAPI.delete,
		getData: ItemAPI.get,
		redirectUrl: ITEMS_URL,
		errorLoadingMessage: failedLoadingError('item'),
		handleDetailsCallback: detailsCallback,
	})

	return (
		<DetailsLayout
			title='Item Details'
			item={detailsController.data as Item}
			details={details}
			handleDeleteConfirm={detailsController.handleDeleteConfirm}
			handleEditClick={detailsController.handleEditClick}
			isLoading={detailsController.isLoading}
			isConfirmationModalOpen={detailsController.isConfirmationModalOpen}
			setIsConfirmationModalOpen={detailsController.setIsConfirmationModalOpen}
			error={detailsController.error}
			viewPermission={CONTENT_READ}
			writePermission={CONTENT_WRITE}
		/>
	)
}

export default ItemDetailsPage
