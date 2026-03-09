'use client'

import { useRouter } from 'next/navigation'
import { Button } from '@pitch-in/shared/components'

interface CancelSubmitButtonProps extends React.ButtonHTMLAttributes<HTMLButtonElement> {
	cancelUrl: string
	children?: React.ReactNode
}

export const CancelSubmitButton = ({
	cancelUrl,
	children,
	...props
}: CancelSubmitButtonProps) => {
	const router = useRouter()

	const handleCancel = () => {
		router.push(cancelUrl)
	}

	return (
		<Button actionType='danger' type='button' onClick={handleCancel} {...props}>
			Cancel
		</Button>
	)
}
