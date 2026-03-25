'use client'

import { useCallback } from 'react'
import { useParams } from 'next/navigation'
import { EditFormLayout } from '@pitch-in/shared/components'
import { CONTENT_WRITE } from '@pitch-in/shared/constants'
import { useEditRecord } from '@pitch-in/shared/hooks'
import { itemCreateSchema, ItemCreateFormData } from '@pitch-in/shared/schemas'
import { FormField, Item } from '@pitch-in/shared/types'
import { failedLoadingError } from '@pitch-in/shared/utils'
import { ITEMS_URL } from '@/lib/constants'
import { ItemAPI } from '@/lib/clients/api'

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
		required: true,
	},
]

const EditItemPage = () => {
	const params = useParams()

	const editController = useEditRecord<typeof itemCreateSchema, Item>({
		id: parseInt(params.id as string),
		defaultValues: {
			name: '',
			description: '',
		},
		getData: ItemAPI.get,
		updateData: ItemAPI.patch,
		errorLoadingMessage: failedLoadingError('item'),
		redirectUrl: ITEMS_URL,
		schema: itemCreateSchema,
		handleFetchCallback: useCallback((i: Item) => {
			return {
				name: i.name,
				description: i.description,
			}
		}, []),
		transformData: async (data: ItemCreateFormData) => {
			return data
		},
	})

	return (
		<EditFormLayout
			permission={CONTENT_WRITE}
			item={editController.data as Item}
			title='Edit Item'
			fields={fields}
			isLoading={editController.isLoading}
			isSubmitting={editController.isSubmitting}
			loadingError={editController.loadingError}
			cancelUrl={`${ITEMS_URL}/${params.id}`}
			handleSubmit={editController.onSubmit}
			register={editController.register}
			control={editController.control}
			errors={editController.fieldErrors}
		/>
	)
}

export default EditItemPage
