'use client'

import { NavbarLink } from '@pitch-in/shared'
import {
	Navbar as SharedNavbar,
	NavLinkItem,
} from '@pitch-in/shared/components'
import { DASHBOARD_URL, IDENTITY_URL } from '@/lib'
import ProfileDropdown from '@/components/layout/ProfileDropdown'

const ADMIN_LINKS: NavLinkItem[] = [
	{ label: 'Dashboard', href: DASHBOARD_URL, permission: '' },
	{ label: 'Identity', href: IDENTITY_URL, permission: '' },
	{ label: 'Images', href: 'blah2', permission: '' },
]

export default function Navbar() {
	// placeholder
	const isAuthenticated = true

	return (
		<SharedNavbar
			links={isAuthenticated ? ADMIN_LINKS : []}
			renderLink={(l) => (
				<NavbarLink key={l.href} href={l.href} permission={l.permission}>
					{l.label}
				</NavbarLink>
			)}
			rightContent={isAuthenticated ? <ProfileDropdown /> : <div>login</div>}
		/>
	)
}
