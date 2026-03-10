'use client'

import { useCallback, useState } from 'react'
import { useParams } from 'next/navigation'
import { DetailsLayout, DetailSectionRow } from '@pitch-in/shared/components'
import { Role } from '@pitch-in/shared/types'
import { failedLoadingError } from '@pitch-in/shared/utils'
import { useDetailsController } from '@pitch-in/shared'
import { ROLES_URL } from '@/lib'
import { RoleAPI } from '@/lib/clients/api'

const RoleDetailsPage = () => {
	const params = useParams()
	const [details, setDetails] = useState<DetailSectionRow[]>([])

	const detailsCallback = useCallback((r: Role) => {
		setDetails([
			{
				label: 'Name',
				value: r.name,
			},
			{
				label: 'Description',
				value: r.description,
			},
		])
	}, [])

	const detailsController = useDetailsController({
		id: parseInt(params.id as string),
		deleteData: RoleAPI.delete,
		getData: RoleAPI.get,
		redirectUrl: ROLES_URL,
		errorLoadingMessage: failedLoadingError('role'),
		handleDetailsCallback: detailsCallback,
	})

	return (
		<DetailsLayout
			title='Role Details'
			item={detailsController.data as Role}
			details={details}
			handleDeleteConfirm={detailsController.handleDeleteConfirm}
			handleEditClick={detailsController.handleEditClick}
			isLoading={detailsController.isLoading}
			isConfirmationModalOpen={detailsController.isConfirmationModalOpen}
			setIsConfirmationModalOpen={detailsController.setIsConfirmationModalOpen}
			error={detailsController.error}
		/>
	)
}

export default RoleDetailsPage
