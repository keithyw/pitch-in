'use client'

import { ReactNode } from 'react'
import { Container, NavbarLink } from '@pitch-in/shared/components'
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
			<div className='container mx-auto flex items-center justify-between'>
				<NavbarLink href='/'>Home</NavbarLink>
				<div className='flex items-center space-x-6'>
					{links.map((l) => renderLink(l))}
					{rightContent}
				</div>
			</div>
		</nav>
	)
}
