'use client'

import { useCallback, useState } from 'react'
import { useParams } from 'next/navigation'
import { DetailsLayout, DetailSectionRow } from '@pitch-in/shared/components'
import { FAILED_LOADING_ASSETS_ERROR } from '@pitch-in/shared/constants'
import { User } from '@pitch-in/shared/types'
import { useDetailsController } from '@pitch-in/shared/hooks'
import { USERS_URL } from '@/lib'
import { UserAPI } from '@/lib/clients/api'

const UserDetailsPage = () => {
	const params = useParams()
	const [details, setDetails] = useState<DetailSectionRow[]>([])

	const detailsCallback = useCallback((res: User) => {
		setDetails([
			{
				label: 'Username',
				value: res.username,
			},
			{
				label: 'Email',
				value: res.email,
			},
			{
				label: 'First Name',
				value: res.first_name,
			},
			{
				label: 'Last Name',
				value: res.last_name,
			},
			{
				label: 'Is Active',
				value: res.is_active ? 'Yes' : 'No',
			},
		])
	}, [])

	const {
		data: user,
		isLoading,
		error,
		handleDeleteConfirm,
		handleEditClick,
		isConfirmationModalOpen,
		setIsConfirmationModalOpen,
	} = useDetailsController({
		id: parseInt(params.id as string),
		deleteData: UserAPI.delete,
		getData: UserAPI.get,
		redirectUrl: USERS_URL,
		errorLoadingMessage: FAILED_LOADING_ASSETS_ERROR,
		handleDetailsCallback: detailsCallback,
	})

	return (
		<DetailsLayout
			title='User Details'
			item={user as User}
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

export default UserDetailsPage
