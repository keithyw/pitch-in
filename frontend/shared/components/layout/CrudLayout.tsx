import { Container } from '@pitch-in/shared/components'
import { Subnavbar, SubnavBarLink } from '@pitch-in/shared/components'

interface CrudLayoutProps {
	children: React.ReactNode
	links: SubnavBarLink[]
	title: string
}

export const CrudLayout = ({ children, links, title }: CrudLayoutProps) => {
	return (
		<div className='min-h-screen bg-gray-50'>
			<header className='items-centered flex justify-between bg-white px-6 py-4 shadow-sm'>
				<h1 className='text-2xl font-semibold text-gray-800'>{title}</h1>
				<Subnavbar links={links} />
			</header>
			<Container as='main' className='px-4 py-8 sm:px-6 lg:px-8'>
				{children}
			</Container>
		</div>
	)
}
