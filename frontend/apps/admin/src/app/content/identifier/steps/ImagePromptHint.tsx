import { useMemo } from 'react'
import { DialogTitle } from '@headlessui/react'
import { Button } from '@pitch-in/shared/components'
import useAIStore from '@/stores/useAIStore'

interface ImagePromptHintProps {
	onHandleSubmit: () => void
}

const ImagePromptHint = ({ onHandleSubmit }: ImagePromptHintProps) => {
	const setPrompt = useAIStore((state) => state.setPrompt)

	const formattedPrompt = useMemo(() => {
		const promptTemplate = `Identify the item in the image`
		return promptTemplate
	}, [])

	const handleClearAndSubmit = () => {
		setPrompt(formattedPrompt)
		onHandleSubmit()
	}

	return (
		<div>
			<DialogTitle
				as='h3'
				className='text-lg leading-6 font-medium text-gray-900'
			>
				Hint:
			</DialogTitle>
			<div className='mt-4'>
				<p>
					<span className='mr-2 font-bold text-gray-900'>
						Formatted Prompt:
					</span>
					<span className='text-gray-900'>{formattedPrompt}</span>
				</p>
			</div>
			<div className='mt-4 flex justify-end'>
				<Button actionType='submit' onClick={handleClearAndSubmit}>
					Done
				</Button>
			</div>
		</div>
	)
}

export default ImagePromptHint
