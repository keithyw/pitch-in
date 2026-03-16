'use client'

import Link from 'next/link'
import { NavbarLink } from '@pitch-in/shared'
import {
	Navbar as SharedNavbar,
	NavLinkItem,
} from '@pitch-in/shared/components'
import ProfileDropdown from '@/components/layout/ProfileDropdown'
import { DASHBOARD_URL, IDENTITY_URL, LOGIN_URL } from '@/lib/constants'
import useAuthStore from '@/stores/useAuthStore'

const ADMIN_LINKS: NavLinkItem[] = [
	{ label: 'Dashboard', href: DASHBOARD_URL, permission: '' },
	{ label: 'Identity', href: IDENTITY_URL, permission: '' },
	{ label: 'Images', href: 'blah2', permission: '' },
]

export default function Navbar() {
	// placeholder
	const { isAuthenticated } = useAuthStore()

	return (
		<SharedNavbar
			links={isAuthenticated ? ADMIN_LINKS : []}
			renderLink={(l) => (
				<NavbarLink key={l.href} href={l.href} permission={l.permission}>
					{l.label}
				</NavbarLink>
			)}
			rightContent={
				isAuthenticated ? (
					<ProfileDropdown />
				) : (
					<Link
						href={LOGIN_URL}
						className='focus:ring-opacity-50 rounded-full bg-blue-500 px-4 py-2 font-bold text-white transition-colors duration-200 hover:bg-blue-600 focus:ring-2 focus:ring-blue-400 focus:outline-none'
					>
						Login
					</Link>
				)
			}
		/>
	)
}
