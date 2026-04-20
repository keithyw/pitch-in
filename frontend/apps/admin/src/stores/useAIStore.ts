import { create } from 'zustand'
import { IdentifierResponse } from '@pitch-in/shared/types'
import { WizardSlice, createWizardSlice } from './wizardStoreSlice'

interface AITools {
	content: IdentifierResponse | null
	image: File | null
	prompt: string
	hasPromptHint: boolean
	isPromptDisabled: boolean

	setContent(content: IdentifierResponse): void
	setImage(image: File): void
	setPrompt(prompt: string): void
	setHasPromptHint(hasPromptHint: boolean): void
	setIsPromptDisabled(isPromptDisabled: boolean): void
	clearDraft: () => void
}

type AIToolsStore = WizardSlice & AITools

const useAIStore = create<AIToolsStore>()((set, get) => ({
	...createWizardSlice(set, () => ({}) as AIToolsStore, {
		setState: set,
		getState: () => ({}) as AIToolsStore,
		subscribe: () => () => {},
		getInitialState: () => ({}) as AIToolsStore,
	}),
	content: null,
	image: null,
	prompt: '',
	hasPromptHint: false,
	isPromptDisabled: false,
	setContent: (c: IdentifierResponse) => set({ content: c }),
	setImage: (i: File) => set({ image: i }),
	setPrompt: (prompt) => set({ prompt }),
	setHasPromptHint: (hasPromptHint) => set({ hasPromptHint }),
	setIsPromptDisabled: (isPromptDisabled) => set({ isPromptDisabled }),
	clearDraft: () => {
		get().reset()
		set({
			content: null,
			image: null,
			prompt: '',
			hasPromptHint: false,
			isPromptDisabled: false,
		})
	},
}))

export default useAIStore
