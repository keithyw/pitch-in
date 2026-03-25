'use client'

import Link from 'next/link'
import { NavbarLink } from '@pitch-in/shared'
import {
	Navbar as SharedNavbar,
	NavLinkItem,
} from '@pitch-in/shared/components'
import {
	ASSETS_READ,
	CONTENT_READ,
	IDENTITY_READ,
} from '@pitch-in/shared/constants'
import ProfileDropdown from '@/components/layout/ProfileDropdown'
import {
	ASSETS_URL,
	DASHBOARD_URL,
	IDENTITY_URL,
	TAXONOMY_URL,
	LOGIN_URL,
} from '@/lib/constants'
import useAuthStore from '@/stores/useAuthStore'

const ADMIN_LINKS: NavLinkItem[] = [
	{ label: 'Dashboard', href: DASHBOARD_URL, permission: IDENTITY_READ },
	{ label: 'Assets', href: ASSETS_URL, permission: ASSETS_READ },
	{ label: 'Identity', href: IDENTITY_URL, permission: IDENTITY_READ },
	{ label: 'Taxonomy', href: TAXONOMY_URL, permission: CONTENT_READ },
]

export default function Navbar() {
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
