'use client'

import { useCallback, useEffect, useState, ComponentType } from 'react'
import { BaseModal, Button, TextareaInput } from '@pitch-in/shared/components'
import { useAIPromptContext } from '@pitch-in/shared/contexts'
import {
	AIServiceException,
	PromptHintComponentProps,
	StepComponentProps,
} from '@pitch-in/shared/types'

interface AIPromptStepProps extends StepComponentProps {
	children?: React.ReactNode
	promptHintComponent?: ComponentType<PromptHintComponentProps>
	promptHintMessage?: string
	onGenerate: () => Promise<boolean>
}

export const AIPromptStep = ({
	setSubmitHandler,
	children,
	promptHintComponent: PromptHintComponent,
	promptHintMessage,
	onGenerate,
}: AIPromptStepProps) => {
	const {
		prompt,
		setPrompt,
		hasPromptHint,
		isPromptDisabled,
		setError,
		setIsCurrentStepValid,
		setIsSubmitting,
	} = useAIPromptContext()
	const [isOpen, setIsOpen] = useState(false)

	const handleStepSubmit = useCallback(async (): Promise<boolean> => {
		let isValid = false
		setIsSubmitting(true)
		setIsCurrentStepValid(isValid)
		setError('')
		try {
			isValid = await onGenerate()
			setIsCurrentStepValid(isValid)
		} catch (e: unknown) {
			// isValid should still be false here so no need to set
			if (e instanceof AIServiceException) {
				console.error('AIServiceException ' + e.message)
				setError(e.message)
			} else if (e instanceof Error) {
				console.error('generic error ', e.message)
				setError(e.message)
			}
		} finally {
			setIsSubmitting(false)
		}
		return isValid
	}, [onGenerate, setError, setIsCurrentStepValid, setIsSubmitting])

	useEffect(() => {
		setSubmitHandler(handleStepSubmit)
		return () => {
			setSubmitHandler(null)
		}
	}, [handleStepSubmit, setSubmitHandler])

	return (
		<div className='space-y-6 p-4'>
			{children}
			<TextareaInput
				id='prompt'
				label='Generate content for:'
				placeholder={promptHintMessage}
				value={prompt}
				rows={6}
				className='min-h-[120px] resize-y'
				onChange={(e) => setPrompt(e.target.value)}
				disabled={isPromptDisabled}
			/>
			{hasPromptHint && (
				<Button actionType='neutral' onClick={() => setIsOpen(true)}>
					Get Help Generating Prompt
				</Button>
			)}
			{PromptHintComponent && (
				<BaseModal isOpen={isOpen} onClose={() => setIsOpen(false)}>
					<div className='p-4'>
						<PromptHintComponent onHandleSubmit={() => setIsOpen(false)} />
					</div>
				</BaseModal>
			)}
		</div>
	)
}
