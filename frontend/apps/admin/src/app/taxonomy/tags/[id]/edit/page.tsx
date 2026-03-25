'use client'

import { useCallback } from 'react'
import { useParams } from 'next/navigation'
import { EditFormLayout } from '@pitch-in/shared/components'
import { CONTENT_WRITE } from '@pitch-in/shared/constants'
import { useEditRecord } from '@pitch-in/shared/hooks'
import { tagCreateSchema, TagCreateFormData } from '@pitch-in/shared/schemas'
import { FormField, Tag } from '@pitch-in/shared/types'
import { failedLoadingError } from '@pitch-in/shared/utils'
import { TAGS_URL } from '@/lib/constants'
import { TagAPI } from '@/lib/clients/api'

const fields: FormField<TagCreateFormData>[] = [
	{
		name: 'tag',
		label: 'Tag',
		placeholder: 'Enter tag',
		required: true,
	},
]

const EditTagPage = () => {
	const params = useParams()

	const editController = useEditRecord<typeof tagCreateSchema, Tag>({
		id: parseInt(params.id as string),
		defaultValues: {
			tag: '',
		},
		getData: TagAPI.get,
		updateData: TagAPI.patch,
		errorLoadingMessage: failedLoadingError('tag'),
		redirectUrl: TAGS_URL,
		schema: tagCreateSchema,
		handleFetchCallback: useCallback((t: Tag) => {
			return {
				tag: t.tag,
			}
		}, []),
		transformData: async (data: TagCreateFormData) => {
			return data
		},
	})

	return (
		<EditFormLayout
			permission={CONTENT_WRITE}
			item={editController.data as Tag}
			title='Edit Tag'
			fields={fields}
			isLoading={editController.isLoading}
			isSubmitting={editController.isSubmitting}
			loadingError={editController.loadingError}
			cancelUrl={`${TAGS_URL}/${params.id}`}
			handleSubmit={editController.onSubmit}
			register={editController.register}
			control={editController.control}
			errors={editController.fieldErrors}
		/>
	)
}

export default EditTagPage
