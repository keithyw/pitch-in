'use client'

import { useCallback } from 'react'
import { useParams } from 'next/navigation'
import { EditFormLayout } from '@pitch-in/shared/components'
import { useEditRecord } from '@pitch-in/shared/hooks'
import { roleCreateSchema, RoleCreateFormData } from '@pitch-in/shared/schemas'
import { FormField, Role } from '@pitch-in/shared/types'
import { failedLoadingError } from '@pitch-in/shared/utils'
import { ROLES_URL } from '@/lib'
import { RoleAPI } from '@/lib/clients/api'
import { describe } from 'zod/v4/core'

const fields: FormField<RoleCreateFormData>[] = [
	{
		name: 'name',
		label: 'Name',
		placeholder: 'Enter name',
		required: true,
	},
	{
		name: 'description',
		label: 'Description',
		placeholder: 'Enter description',
		required: false,
	},
]

const EditRolePage = () => {
	const params = useParams()

	const editController = useEditRecord<typeof roleCreateSchema, Role>({
		id: parseInt(params.id as string),
		defaultValues: {
			name: '',
			description: '',
		},
		getData: RoleAPI.get,
		updateData: RoleAPI.patch,
		errorLoadingMessage: failedLoadingError('role'),
		redirectUrl: ROLES_URL,
		schema: roleCreateSchema,
		handleFetchCallback: useCallback((r: Role) => {
			return {
				name: r.name,
				description: r.description ?? '',
			}
		}, []),
		transformData: async (data: RoleCreateFormData) => {
			return data
		},
	})

	return (
		<EditFormLayout
			permission=''
			item={editController.data as Role}
			title='Edit Role'
			fields={fields}
			isLoading={editController.isLoading}
			isSubmitting={editController.isSubmitting}
			loadingError={editController.loadingError}
			cancelUrl={`${ROLES_URL}/${params.id}`}
			handleSubmit={editController.onSubmit}
			register={editController.register}
			control={editController.control}
			errors={editController.fieldErrors}
		/>
	)
}

export default EditRolePage
