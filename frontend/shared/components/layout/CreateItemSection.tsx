import Link from 'next/link'
import { Button } from '@pitch-in/shared/components'

interface CreateItemSectionProps {
	href: string
	permission?: string
	children: React.ReactNode
}

export const CreateItemSection = ({
	href,
	permission,
	children,
}: CreateItemSectionProps) => {
	return (
		<>
			<div className='mb-4 flex justify-end'>
				<Link href={href} passHref>
					<Button actionType='edit' type='button'>
						{children}
					</Button>
				</Link>
			</div>
		</>
	)
}
