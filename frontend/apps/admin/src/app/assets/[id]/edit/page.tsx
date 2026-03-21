'use client'

import { useEffect, useState } from 'react'
import { useForm, Controller } from 'react-hook-form'
import toast from 'react-hot-toast'
import { useParams, useRouter } from 'next/navigation'
import { zodResolver } from '@hookform/resolvers/zod'
import {
	CreateFormLayout,
	FileUploadInput,
	FormInput,
} from '@pitch-in/shared/components'
import {
	ASSETS_WRITE,
	FAILED_LOADING_ASSETS_ERROR,
} from '@pitch-in/shared/constants'
import {
	assetCreateSchema,
	AssetCreateFormData,
} from '@pitch-in/shared/schemas'
import { Asset, FormField, CreateAssetRequest } from '@pitch-in/shared/types'
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

const EditAssetPage = () => {
	const params = useParams()
	const router = useRouter()
	const [isLoading, setIsLoading] = useState<boolean>(false)
	const [asset, setAsset] = useState<Asset | null>(null)
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

	useEffect(() => {
		if (!params.id) return
		setIsLoading(true)
		const fetchAsset = async () => {
			try {
				const res = await AssetAPI.get(parseInt(params.id as string))
				setAsset(res)
				reset({ object_key: res.object_key })

				if (res.mime_type.startsWith('image/')) {
					setPreview(res.url)
					if (res.width && res.height) {
						setDimensions({ width: res.width, height: res.height })
					}
				}
			} catch (e: unknown) {
				if (e instanceof Error) {
					console.error(e.message)
					toast.error(FAILED_LOADING_ASSETS_ERROR)
					router.push(ASSETS_URL)
				}
			} finally {
				setIsLoading(false)
			}
		}
		fetchAsset()
	}, [params.id])

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
			const req: Partial<CreateAssetRequest> = {
				object_key: data.object_key,
			}

			if (data.file) {
				const meta = await getImageDimensionsFromFile(data.file)
				req.file = data.file
				req.width = meta.width
				req.height = meta.height
			}
			await AssetAPI.patch(parseInt(params.id as string), req)
			toast.success(`Asset ${req.object_key} udpated successfully!`)
			reset()
			router.push(ASSETS_URL)
		} catch (e: unknown) {
			handleFormErrors<AssetCreateFormData>(
				e,
				setError,
				'Failed to update asset.',
			)
		}
	}

	return (
		<CreateFormLayout
			title='Edit Asset'
			isSubmitting={isSubmitting}
			submitText='Edit'
			submittingText='Saving...'
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
					<p className='mb-2 text-center text-[10px] font-bold text-gray-400 uppercase'>
						{watchedFile ? 'New Selection Preview' : 'Current Asset'}
					</p>
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

export default EditAssetPage
