import { cn } from '@pitch-in/shared/utils'

interface CardContainerProps {
	children: React.ReactNode
	className?: string
}

export const CardContainer = ({ children, className }: CardContainerProps) => {
	return (
		<div
			className={cn(
				'container mx-auto rounded-lg bg-white p-8 shadow-md',
				className,
			)}
		>
			{children}
		</div>
	)
}
