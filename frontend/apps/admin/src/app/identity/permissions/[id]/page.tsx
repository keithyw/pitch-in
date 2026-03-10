'use client'

import { useCallback, useState } from 'react'
import { useParams } from 'next/navigation'
import { DetailsLayout, DetailSectionRow } from '@pitch-in/shared/components'
import { FAILED_LOADING_PERMISSION_ERROR } from '@pitch-in/shared/constants'
import { Permission } from '@pitch-in/shared/types'
import { useDetailsController } from '@pitch-in/shared'
import { PERMISSIONS_URL } from '@/lib'
import { PermissionAPI } from '@/lib/clients/api'

const PermissionDetailsPage = () => {
	const params = useParams()
	const [details, setDetails] = useState<DetailSectionRow[]>([])

	const detailsCallback = useCallback((p: Permission) => {
		setDetails([
			{
				label: 'Code',
				value: p.code,
			},
			{
				label: 'Display Name',
				value: p.display_name,
			},
			{
				label: 'Path',
				value: p.path,
			},
			{
				label: 'Method',
				value: p.method,
			},
		])
	}, [])

	const {
		data: permission,
		isLoading,
		error,
		handleDeleteConfirm,
		handleEditClick,
		isConfirmationModalOpen,
		setIsConfirmationModalOpen,
	} = useDetailsController({
		id: parseInt(params.id as string),
		deleteData: PermissionAPI.delete,
		getData: PermissionAPI.get,
		redirectUrl: PERMISSIONS_URL,
		errorLoadingMessage: FAILED_LOADING_PERMISSION_ERROR,
		handleDetailsCallback: detailsCallback,
	})

	return (
		<DetailsLayout
			title='Permission Details'
			item={permission as Permission}
			details={details}
			handleDeleteConfirm={handleDeleteConfirm}
			handleEditClick={handleEditClick}
			isLoading={isLoading}
			isConfirmationModalOpen={isConfirmationModalOpen}
			setIsConfirmationModalOpen={setIsConfirmationModalOpen}
			error={error}
		/>
	)
}

export default PermissionDetailsPage
