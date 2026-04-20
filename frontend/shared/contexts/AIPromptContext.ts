'use client'

import { createContext, useContext } from 'react'

export interface AIPromptContextValue {
	currentStep: number
	isCurrentStepValid: boolean
	isSubmitting: boolean
	error: string | null
	prompt: string
	setPrompt: (prompt: string) => void
	isPromptDisabled: boolean
	hasPromptHint: boolean
	setCurrentStep: (step: number) => void
	setError: (error: string) => void
	setIsCurrentStepValid: (isValid: boolean) => void
	setIsSubmitting: (isSubmitting: boolean) => void
	clearDraft: () => void
}

export const AIPromptContext = createContext<AIPromptContextValue | null>(null)

export const useAIPromptContext = () => {
	const ctx = useContext(AIPromptContext)
	if (!ctx)
		throw new Error('AIPromptStep must be used within AIPromptContext.Provider')
	return ctx
}
