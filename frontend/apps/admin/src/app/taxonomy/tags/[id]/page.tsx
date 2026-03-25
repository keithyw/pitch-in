'use client'

import { useCallback, useState } from 'react'
import { useParams } from 'next/navigation'
import { DetailsLayout, DetailSectionRow } from '@pitch-in/shared/components'
import { CONTENT_READ, CONTENT_WRITE } from '@pitch-in/shared/constants'
import { Tag } from '@pitch-in/shared/types'
import { failedLoadingError } from '@pitch-in/shared/utils'
import { useDetailsController } from '@pitch-in/shared'
import { TAGS_URL } from '@/lib/constants'
import { TagAPI } from '@/lib/clients/api'

const TagDetailsPage = () => {
	const params = useParams()
	const [details, setDetails] = useState<DetailSectionRow[]>([])

	const detailsCallback = useCallback((t: Tag) => {
		setDetails([
			{
				label: 'Id',
				value: t.id.toString(),
			},
			{
				label: 'Tag',
				value: t.tag,
			},
			{
				label: 'Slug',
				value: t.slug,
			},
		])
	}, [])

	const detailsController = useDetailsController({
		id: parseInt(params.id as string),
		deleteData: TagAPI.delete,
		getData: TagAPI.get,
		redirectUrl: TAGS_URL,
		errorLoadingMessage: failedLoadingError('tag'),
		handleDetailsCallback: detailsCallback,
	})

	return (
		<DetailsLayout
			title='Tag Details'
			item={detailsController.data as Tag}
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

export default TagDetailsPage
