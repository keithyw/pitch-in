'use client'

import {
	AIGenerationLayout,
	PermissionGuard,
} from '@pitch-in/shared/components'
import { CONTENT_WRITE } from '@pitch-in/shared/constants'
import { AIPromptContext } from '@pitch-in/shared/contexts'
import { WizardStepType } from '@pitch-in/shared/types'
import { ITEMS_URL } from '@/lib/constants'
import useAIStore from '@/stores/useAIStore'
import ImagePromptStep from './steps/ImagePromptStep'
import ContentReviewStep from './steps/ContentReviewStep'

const steps: WizardStepType[] = [
	{ id: 'prompt', title: 'Generate Prompt', component: ImagePromptStep },
	{ id: 'review', title: 'Review Content', component: ContentReviewStep },
]

const IdentifierPage = () => {
	const store = useAIStore()
	return (
		<PermissionGuard requiredPermission={CONTENT_WRITE}>
			<AIPromptContext.Provider value={store}>
				<AIGenerationLayout
					successUrl={ITEMS_URL}
					title='Identify Item in Image'
					wizardSteps={steps}
				/>
			</AIPromptContext.Provider>
		</PermissionGuard>
	)
}

export default IdentifierPage
