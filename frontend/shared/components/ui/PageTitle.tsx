import { cn } from '@pitch-in/shared/utils'

interface PageTitleProps {
	children: React.ReactNode
	className?: string
}

export const PageTitle = ({ children, className }: PageTitleProps) => {
	return (
		<h1
			className={cn(
				'mb-6 text-center text-4xl font-extrabold text-gray-900',
				className,
			)}
		>
			{children}
		</h1>
	)
}
