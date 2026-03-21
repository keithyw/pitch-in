'use client'

import { useState } from 'react'
import { useForm, Controller } from 'react-hook-form'
import toast from 'react-hot-toast'
import { useRouter } from 'next/navigation'
import { zodResolver } from '@hookform/resolvers/zod'
import {
	CreateFormLayout,
	FileUploadInput,
	FormInput,
} from '@pitch-in/shared/components'
import { ASSETS_WRITE } from '@pitch-in/shared/constants'
import {
	assetCreateSchema,
	AssetCreateFormData,
} from '@pitch-in/shared/schemas'
import { FormField, CreateAssetRequest } from '@pitch-in/shared/types'
import {
	getImageDimensionsFromFile,
	getFileMetadata,
	handleFormErrors,
} from '@pitch-in/shared/utils'
import { ASSETS_URL } from '@/lib/constants'
import { AssetAPI } from '@/lib/clients/api'

const fields: FormField<AssetCreateFormData>[] = [
	{
		name: 'object_key',
		label: 'Object Key',
		placeholder: 'Enter a name',
		required: true,
	},
]

const CreateAssetPage = () => {
	const router = useRouter()
	const [dimensions, setDimensions] = useState<{
		width: number
		height: number
	} | null>(null)
	const [preview, setPreview] = useState<string | null>(null)
	const {
		register,
		handleSubmit,
		setError,
		setValue,
		formState: { errors, isSubmitting },
		reset,
		control,
		watch,
	} = useForm<AssetCreateFormData>({
		resolver: zodResolver(assetCreateSchema),
		defaultValues: {
			object_key: '',
		},
	})

	const watchedFile = watch('file')

	const onFileChange = async (
		file: File | null,
		onChange: (f: File | null) => void,
	) => {
		if (!file) {
			setDimensions(null)
			setPreview(null)
			onChange(null)
			return
		}

		const meta = await getFileMetadata(file)
		setValue('object_key', meta.fileName, { shouldValidate: true })
		if (meta.previewUrl) setPreview(meta.previewUrl)
		if (meta.width && meta.height)
			setDimensions({ width: meta.width, height: meta.height })
		onChange(file)
	}

	const onSubmit = async (data: AssetCreateFormData) => {
		try {
			const { width, height } = await getImageDimensionsFromFile(data.file)
			const req: CreateAssetRequest = {
				object_key: data.object_key,
				file: data.file,
				width: width,
				height: height,
			}
			await AssetAPI.createWithFile(req)
			toast.success(`Asset ${req.object_key} created successfully!`)
			reset()
			router.push(ASSETS_URL)
		} catch (e: unknown) {
			handleFormErrors<AssetCreateFormData>(
				e,
				setError,
				'Failed to create asset.',
			)
		}
	}

	return (
		<CreateFormLayout
			title='Create Asset'
			isSubmitting={isSubmitting}
			submitText='Create'
			submittingText='Creating...'
			handleSubmit={handleSubmit(onSubmit)}
			requiredPermission={ASSETS_WRITE}
		>
			<Controller
				name='file'
				control={control}
				render={({ field: { onChange } }) => (
					<FileUploadInput
						label='Asset'
						onChange={(file) => onFileChange(file, onChange)}
						currentFile={watchedFile as File | null}
					/>
				)}
			/>
			{preview && (
				<div className='mt-4 rounded-md border bg-gray-50 p-2'>
					<img
						src={preview}
						alt='Preview'
						className='mx-auto max-h-40 rounded shadow-sm'
					/>
					{dimensions && (
						<p className='mt-2 text-center text-xs text-gray-500'>
							Dimensions: {dimensions.width}px X {dimensions.height}px
						</p>
					)}
				</div>
			)}
			{!preview && watchedFile && (
				<div className='mmt-4 dashed flex items-center justify-center rounded-md border bg-gray-50 p-4'>
					<span className='text-sm text-gray-600'>
						File Type: {(watchedFile as File).type || 'Unknown'}
					</span>
				</div>
			)}
			{fields.map((f, idx) => (
				<FormInput
					key={idx}
					field={f}
					register={register}
					control={control}
					errorMessage={errors[f.name]?.message as string}
				/>
			))}
		</CreateFormLayout>
	)
}

export default CreateAssetPage
