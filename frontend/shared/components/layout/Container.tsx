import { cn } from '@pitch-in/shared/utils'

interface ContainerProps {
	children: React.ReactNode
	className?: string
	as?: 'div' | 'section' | 'main'
}

export const Container = ({
	children,
	className = '',
	as: Component = 'div',
}: ContainerProps) => {
	return (
		<Component
			className={cn(
				'container mx-auto w-full max-w-7xl px-4 sm:px-6 lg:px-8',
				className,
			)}
		>
			{children}
		</Component>
	)
}
