'use client'

import { useCallback } from 'react'
import { useParams } from 'next/navigation'
import { EditFormLayout } from '@pitch-in/shared/components'
import { FAILED_LOADING_PERMISSION_ERROR } from '@pitch-in/shared/constants'
import { useEditRecord } from '@pitch-in/shared/hooks'
import {
	permissionCreateSchema,
	PermissionCreateFormData,
} from '@pitch-in/shared/schemas'
import { FormField, Permission } from '@pitch-in/shared/types'
import { PERMISSIONS_URL } from '@/lib'
import { PermissionAPI } from '@/lib/clients/api'

const fields: FormField<PermissionCreateFormData>[] = [
	{
		name: 'code',
		label: 'Code',
		placeholder: 'Enter code',
		required: true,
	},
	{
		name: 'display_name',
		label: 'Display Name',
		placeholder: 'Enter display name',
		required: false,
	},
	{
		name: 'path',
		label: 'Path',
		placeholder: 'Enter path',
		required: false,
	},
	{
		name: 'method',
		label: 'Method',
		placeholder: 'Enter method',
		required: false,
	},
]

const EditPermissionPage = () => {
	const params = useParams()

	const {
		data: permission,
		isLoading,
		fieldErrors,
		isSubmitting,
		loadingError,
		register,
		control,
		onSubmit,
	} = useEditRecord<typeof permissionCreateSchema, Permission>({
		id: parseInt(params.id as string),
		defaultValues: {
			code: '',
			display_name: '',
			path: '',
			method: '',
		},
		getData: PermissionAPI.get,
		updateData: PermissionAPI.patch,
		errorLoadingMessage: FAILED_LOADING_PERMISSION_ERROR,
		redirectUrl: PERMISSIONS_URL,
		schema: permissionCreateSchema,
		handleFetchCallback: useCallback((res: Permission) => {
			return {
				code: res.code,
				display_name: res.display_name ?? '',
				path: res.path ?? '',
				method: res.method ?? '',
			}
		}, []),
		transformData: async (data: PermissionCreateFormData) => {
			return data
		},
	})
	return (
		<EditFormLayout
			permission=''
			item={permission as Permission}
			title='Edit Permission'
			fields={fields}
			isLoading={isLoading}
			isSubmitting={isSubmitting}
			loadingError={loadingError}
			cancelUrl={`${PERMISSIONS_URL}/${params.id}`}
			handleSubmit={onSubmit}
			register={register}
			control={control}
			errors={fieldErrors}
		/>
	)
}

export default EditPermissionPage
