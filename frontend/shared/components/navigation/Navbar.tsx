'use client'

import { ReactNode } from 'react'
import { Container } from '@pitch-in/shared/components'
import { cn } from '@pitch-in/shared/utils'

export interface NavLinkItem {
	label: string
	href: string
	permission?: string
}

interface NavbarProps {
	links: NavLinkItem[]
	className?: string
	renderLink: (link: NavLinkItem) => ReactNode
	rightContent?: ReactNode
}

export const Navbar = ({
	links,
	className,
	renderLink,
	rightContent,
}: NavbarProps) => {
	return (
		<nav className={cn('bg-gray-800 p-4 text-white shadow-md', className)}>
			<Container className='flex items-center justify-between'>
				Home
				<div className='flex items-center space-x-6'>
					{links.map((l) => renderLink(l))}
				</div>
				<div className='flex items-center space-x-4'>{rightContent}</div>
			</Container>
		</nav>
	)
}
