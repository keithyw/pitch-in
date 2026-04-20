'use client'

import { useCallback, useEffect, useMemo, useState } from 'react'
import { ChipContainer } from '@pitch-in/shared/components'
import { StepComponentProps } from '@pitch-in/shared/types'
import { IngestionAPI } from '@/lib/clients/api'
import useAIStore from '@/stores/useAIStore'

interface TagObject {
	id: string
	label: string
}

const ContentReviewStep = ({ setSubmitHandler }: StepComponentProps) => {
	const { content, image, setIsSubmitting } = useAIStore()
	const [localTags, setLocalTags] = useState<string[]>(content?.tags || [])

	const chipData = useMemo(() => {
		return localTags.map((t, idx) => ({
			id: `${t}-${idx}`,
			label: t,
		}))
	}, [localTags])

	const handleFinish = useCallback(async (): Promise<boolean> => {
		setIsSubmitting(true)
		if (!content || !image) {
			setIsSubmitting(false)
			return false
		}
		try {
			const item = {
				name: content.name,
				description: content.description,
				tags: localTags,
			}
			await IngestionAPI.ingest(item, image)
		} catch (e: unknown) {
			console.error(e)
			setIsSubmitting(false)
			return false
		} finally {
			setIsSubmitting(false)
		}
		return true
	}, [content, image, setIsSubmitting])

	useEffect(() => {
		setSubmitHandler(handleFinish)
		return () => {
			setSubmitHandler(null)
		}
	}, [handleFinish, setSubmitHandler])

	const handleRemove = useCallback((i: TagObject) => {
		setLocalTags((prev) => prev.filter((t) => t !== i.label))
	}, [])

	if (!content)
		return (
			<div className='mx-auto rounded-lg p-6 shadow-md'>
				<p className='text-gray-900'>No information returned on image</p>
			</div>
		)

	return (
		<div className='mx-auto rounded-lg p-6 shadow-md'>
			<p className='font-semibold text-gray-900 italic'>Review Content</p>
			<div className='mt-4'>
				<div className='relative mt-1 text-gray-900'>
					<div className='relative w-full cursor-default'>
						<div className='flex flex-wrap gap-2 p-2'>
							<h4 className='font-semibold text-gray-800'>{content.name}</h4>
						</div>
						<div className='flex flex-wrap gap-2 p-2'>
							{content.description}
						</div>
						<div className='flex flex-wrap gap-2 p-2'>
							<p>Tags</p>
							{content.tags.length > 0 ? (
								<ChipContainer
									itemName='tags'
									fieldName='label'
									isLoadingData={false}
									data={chipData}
									onRemove={handleRemove}
									errors=''
								/>
							) : (
								<p className='text-sm text-gray-500 italic'>No tags</p>
							)}
						</div>
					</div>
				</div>
			</div>
		</div>
	)
}

export default ContentReviewStep
