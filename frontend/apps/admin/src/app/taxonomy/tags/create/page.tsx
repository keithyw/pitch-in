'use client'

import { CreateFormLayout, FormInput } from '@pitch-in/shared/components'
import { CONTENT_WRITE } from '@pitch-in/shared/constants'
import { useCreateRecord } from '@pitch-in/shared/hooks'
import { tagCreateSchema, TagCreateFormData } from '@pitch-in/shared/schemas'
import { FormField } from '@pitch-in/shared/types'
import { TagAPI } from '@/lib/clients/api'
import { TAGS_URL } from '@/lib'

const fields: FormField<TagCreateFormData>[] = [
	{
		name: 'tag',
		label: 'Tag',
		placeholder: 'Enter tag',
		required: true,
	},
]

const CreateTagPage = () => {
	const createConntroller = useCreateRecord({
		schema: tagCreateSchema,
		defaultValues: {
			tag: '',
		},
		createFn: TagAPI.create,
		redirectUrl: TAGS_URL,
	})

	return (
		<CreateFormLayout
			title='Create Tag'
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

export default CreateTagPage
