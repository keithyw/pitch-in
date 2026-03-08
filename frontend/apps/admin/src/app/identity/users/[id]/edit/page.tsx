'use client'

import { useCallback } from 'react'
import { useParams } from 'next/navigation'
import { EditFormLayout } from '@pitch-in/shared/components'
import { useEditRecord } from '@pitch-in/shared/hooks'
import { userCreateSchema, UserCreateFormData } from '@pitch-in/shared/schemas'
import { FormField, User } from '@pitch-in/shared/types'
import { USERS_URL } from '@/lib'
import { UserAPI } from '@/lib/clients/api'

const fields: FormField<UserCreateFormData>[] = [
	{
		name: 'username',
		label: 'Username',
		placeholder: 'Enter username',
		required: true,
	},
	{
		name: 'email',
		label: 'Email',
		placeholder: 'Enter email',
		required: true,
	},
	{
		name: 'first_name',
		label: 'First Name',
		placeholder: 'Enter a first name',
		required: true,
	},
	{
		name: 'last_name',
		label: 'Last Name',
		placeholder: 'Enter a last name',
		required: true,
	},
]

const EditUserPage = () => {
	const params = useParams()

	const {
		data: user,
		isLoading,
		fieldErrors,
		isSubmitting,
		loadingError,
		register,
		control,
		onSubmit,
	} = useEditRecord({
		id: parseInt(params.id as string),
		defaultValues: {
			username: '',
			email: '',
			first_name: '',
			last_name: '',
			is_active: true,
		},
		getData: UserAPI.get,
		updateData: UserAPI.patch,
		errorLoadingMessage: '',
		redirectUrl: USERS_URL,
		schema: userCreateSchema,
		handleFetchCallback: useCallback((data: User) => {
			return {
				username: data.username,
				email: data.email,
				first_name: data.first_name,
				last_name: data.last_name,
				is_active: data.is_active,
			}
		}, []),
		transformData: async (data: UserCreateFormData) => {
			return {
				username: data.username,
				email: data.email,
				first_name: data.first_name,
				last_name: data.last_name,
				is_active: data.is_active,
			}
		},
	})

	return (
		<EditFormLayout
			permission=''
			item={user as User}
			title='Edit User'
			fields={fields}
			isLoading={isLoading}
			isSubmitting={isSubmitting}
			loadingError={loadingError}
			cancelUrl={`${USERS_URL}/${params.id}`}
			handleSubmit={onSubmit}
			register={register}
			control={control}
			errors={fieldErrors}
		/>
	)
}

export default EditUserPage
