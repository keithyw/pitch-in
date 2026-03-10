'use client'

import {
	CreateFormLayout,
	FormInput,
	PermissionGuard,
} from '@pitch-in/shared/components'
import { useCreateRecord } from '@pitch-in/shared/hooks'
import { roleCreateSchema, RoleCreateFormData } from '@pitch-in/shared/schemas'
import { FormField, Role } from '@pitch-in/shared/types'
import { RoleAPI } from '@/lib/clients/api'
import { ROLES_URL } from '@/lib'

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

const CreateRolePage = () => {
	const {
		onSubmit,
		register,
		formState: { errors, isSubmitting },
	} = useCreateRecord({
		schema: roleCreateSchema,
		defaultValues: {
			name: '',
			description: '',
		},
		createFn: RoleAPI.create,
		redirectUrl: ROLES_URL,
	})
	return (
		<PermissionGuard>
			<CreateFormLayout
				title='Create Role'
				isSubmitting={isSubmitting}
				submitText='Create'
				submittingText='Creating...'
				handleSubmit={onSubmit}
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
		</PermissionGuard>
	)
}

export default CreateRolePage
