'use client'

import { CreateFormLayout, FormInput } from '@pitch-in/shared/components'
import { CONTENT_WRITE } from '@pitch-in/shared/constants'
import { useCreateRecord } from '@pitch-in/shared/hooks'
import { itemCreateSchema, ItemCreateFormData } from '@pitch-in/shared/schemas'
import { FormField } from '@pitch-in/shared/types'
import { ItemAPI } from '@/lib/clients/api'
import { ITEMS_URL } from '@/lib/constants'

const fields: FormField<ItemCreateFormData>[] = [
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

const CreateItemPage = () => {
	const createConntroller = useCreateRecord({
		schema: itemCreateSchema,
		defaultValues: {
			name: '',
			description: '',
		},
		createFn: ItemAPI.create,
		redirectUrl: ITEMS_URL,
	})

	return (
		<CreateFormLayout
			title='Create Item'
			isSubmitting={createConntroller.formState.isSubmitting}
			submitText='Create'
			submittingText='Creating...'
			handleSubmit={createConntroller.onSubmit}
			requiredPermission={CONTENT_WRITE}
		>
			{fields.map((f, idx) => (
				<FormInput
					key={idx}
					field={f}
					register={createConntroller.register}
					errorMessage={
						createConntroller.formState.errors[f.name]?.message as string
					}
				/>
			))}
		</CreateFormLayout>
	)
}

export default CreateItemPage
