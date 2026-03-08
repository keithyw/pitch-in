'use client'

import {
	CreateFormLayout,
	FormInput,
	PermissionGuard,
} from '@pitch-in/shared/components'
import { useCreateRecord } from '@pitch-in/shared/hooks'
import { userCreateSchema, UserCreateFormData } from '@pitch-in/shared/schemas'
import { FormField } from '@pitch-in/shared/types'
import { UserAPI } from '@/lib/clients/api'
import { USERS_URL } from '@/lib'

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

const CreateUserPage = () => {
	const {
		onSubmit,
		register,
		formState: { errors, isSubmitting },
	} = useCreateRecord({
		schema: userCreateSchema,
		defaultValues: {
			username: '',
			email: '',
			first_name: '',
			last_name: '',
			is_active: true,
		},
		createFn: UserAPI.create,
		redirectUrl: USERS_URL,
	})
	return (
		<PermissionGuard>
			<CreateFormLayout
				title='Create User'
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

export default CreateUserPage
