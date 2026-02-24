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
				'hover:text-primary transition-colors',
				className,
				isActive && activeClassName,
			)}
		>
			{children}
		</Link>
	)
}
