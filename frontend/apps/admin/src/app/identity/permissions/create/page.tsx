'use client'

import { CreateFormLayout, FormInput } from '@pitch-in/shared/components'
import { useCreateRecord } from '@pitch-in/shared/hooks'
import {
	permissionCreateSchema,
	PermissionCreateFormData,
} from '@pitch-in/shared/schemas'
import { IDENTITY_WRITE } from '@pitch-in/shared/constants'
import { FormField } from '@pitch-in/shared/types'
import { PermissionAPI } from '@/lib/clients/api'
import { PERMISSIONS_URL } from '@/lib'

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

const CreatePermissionPage = () => {
	const {
		onSubmit,
		register,
		formState: { errors, isSubmitting },
	} = useCreateRecord({
		schema: permissionCreateSchema,
		defaultValues: {
			code: '',
			display_name: '',
			path: '',
			method: '',
		},
		createFn: PermissionAPI.create,
		redirectUrl: PERMISSIONS_URL,
	})

	return (
		<CreateFormLayout
			title='Create Permission'
			isSubmitting={isSubmitting}
			submitText='Create'
			submittingText='Creating...'
			handleSubmit={onSubmit}
			requiredPermission={IDENTITY_WRITE}
		>
			{fields.map((f, idx) => (
				<FormInput
					key={idx}
					field={f}
					register={register}
					errorMessage={errors[f.name]?.message as string}
				/>
			))}
		</CreateFormLayout>
	)
}

export default CreatePermissionPage
