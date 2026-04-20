'use client'

import { useCallback, useEffect, useState } from 'react'
import Dropzone from 'react-dropzone'
import { AIPromptStep } from '@pitch-in/shared/components'
import { StepComponentProps } from '@pitch-in/shared/types'
import { IdentifierAPI } from '@/lib/clients/api'
import useAIStore from '@/stores/useAIStore'
import ImagePromptHint from './ImagePromptHint'

const ImagePromptStep = ({ setSubmitHandler }: StepComponentProps) => {
	const [files, setFiles] = useState<Blob[]>([])
	const [previewImage, setPreviewImage] = useState<string | null>(null)
	const {
		prompt,
		setContent,
		setImage,
		setHasPromptHint,
		setIsCurrentStepValid,
	} = useAIStore()

	useEffect(() => {
		setHasPromptHint(true)
	}, [setHasPromptHint])

	useEffect(() => {
		setIsCurrentStepValid(prompt.trim().length > 6 && files.length > 0)
	}, [files, prompt, setIsCurrentStepValid])

	useEffect(() => {
		URL.revokeObjectURL(previewImage as string)
	}, [previewImage])

	const handleImage = useCallback(
		(acceptedFiles: Blob[]) => {
			setFiles(acceptedFiles)
			setPreviewImage(URL.createObjectURL(acceptedFiles[0]))
			const file = acceptedFiles[0] as File
			const image = new File([acceptedFiles[0]], file.name, { type: file.type })
			setImage(image)
		},
		[setImage],
	)

	const handleGenerate = useCallback(async (): Promise<boolean> => {
		try {
			const res = await IdentifierAPI.identify(prompt, files[0])
			setContent(res)
			return true
		} catch (e) {
			console.error(e)
			return false
		}
	}, [files, prompt, setContent])

	return (
		<AIPromptStep
			setSubmitHandler={setSubmitHandler}
			promptHintComponent={ImagePromptHint}
			promptHintMessage='Analyze the item in the image'
			onGenerate={handleGenerate}
		>
			<Dropzone multiple={false} onDrop={(f) => handleImage(f)}>
				{({ getRootProps, getInputProps }) => (
					<div className='mt-2 rounded-lg p-8 shadow-md'>
						<div {...getRootProps()}>
							<input {...getInputProps()} />
							<p className='font-bol text-gray-900'>
								Drag and Drop an Image for Analysis Here.
							</p>
						</div>
					</div>
				)}
			</Dropzone>
			{previewImage && (
				<aside className='mt-4 flex flex-row flex-wrap'>
					<div className='mr-2 mb-2 box-border inline-flex h-25 w-25 rounded-sm border-gray-200 p-1'>
						<div className='flex min-w-0 overflow-hidden'>
							<img
								alt='preview image'
								src={previewImage as string}
								className='block h-full w-auto'
								onLoad={() => URL.revokeObjectURL(previewImage)}
							/>
						</div>
					</div>
				</aside>
			)}
		</AIPromptStep>
	)
}

export default ImagePromptStep
