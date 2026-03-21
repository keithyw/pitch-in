'use client'

import { useCallback, useState } from 'react'
import { useParams } from 'next/navigation'
import { DetailsLayout, DetailSectionRow } from '@pitch-in/shared/components'
import {
	ASSETS_READ,
	ASSETS_WRITE,
	FAILED_LOADING_ASSETS_ERROR,
} from '@pitch-in/shared/constants'
import { useDetailsController } from '@pitch-in/shared/hooks'
import { Asset } from '@pitch-in/shared/types'
import { ASSETS_URL } from '@/lib/constants'
import { AssetAPI } from '@/lib/clients/api'

const AssetDetailsPage = () => {
	const params = useParams()
	const [details, setDetails] = useState<DetailSectionRow[]>([])

	const detailsCallback = useCallback((a: Asset) => {
		setDetails([
			{
				label: 'File',
				value: a.url,
				isAsset: true,
				type: a.mime_type.startsWith('image/') ? 'image' : a.mime_type,
			},
			{
				label: 'Object Key',
				value: a.object_key,
			},
			{
				label: 'mime_type',
				value: a.mime_type,
			},
			{
				label: 'Size (Bytes)',
				value: a.size_bytes.toString(),
			},
			{
				label: 'Dimensions',
				value:
					a.width && a.height
						? [a.width, 'px', ' X ', a.height, 'px'].join('')
						: 'N/A',
			},
		])
	}, [])

	const controller = useDetailsController({
		id: parseInt(params.id as string),
		deleteData: AssetAPI.delete,
		getData: AssetAPI.get,
		redirectUrl: ASSETS_URL,
		errorLoadingMessage: FAILED_LOADING_ASSETS_ERROR,
		handleDetailsCallback: detailsCallback,
	})

	return (
		<DetailsLayout
			title='Asset Details'
			item={controller.data as Asset}
			details={details}
			handleDeleteConfirm={controller.handleDeleteConfirm}
			handleEditClick={controller.handleEditClick}
			isLoading={controller.isLoading}
			isConfirmationModalOpen={controller.isConfirmationModalOpen}
			setIsConfirmationModalOpen={controller.setIsConfirmationModalOpen}
			error={controller.error}
			viewPermission={ASSETS_READ}
			writePermission={ASSETS_WRITE}
		/>
	)
}

export default AssetDetailsPage
