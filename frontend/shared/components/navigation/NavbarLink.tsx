'use client'

import Link from 'next/link'
import { usePathname } from 'next/navigation'
import { cn } from '@pitch-in/shared/utils'

interface NavbarLinkProps {
	href: string
	children: React.ReactNode
	className?: string
	permission?: string
	activeClassName?: string
}

export const NavbarLink = ({
	href,
	children,
	className,
	permission,
	activeClassName = 'text primary font-bold',
}: NavbarLinkProps) => {
	const pathname = usePathname()
	const isActive = pathname === href || pathname.startsWith(`${href}/`)

	return (
		<Link
			href={href}
			className={cn(
				'transition-colors duration-200 hover:text-gray-300',
				className,
				isActive && activeClassName,
			)}
		>
			{children}
		</Link>
	)
}
