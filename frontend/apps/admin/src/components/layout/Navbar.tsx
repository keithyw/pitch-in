'use client'

import { UserCircle } from 'lucide-react'
import { NavbarLink } from '@pitch-in/shared'
import {
	Navbar as SharedNavbar,
	NavLinkItem,
} from '@pitch-in/shared/components'
import { DASHBOARD_URL } from '@/lib'

const ADMIN_LINKS: NavLinkItem[] = [
	{ label: 'Dashboard', href: DASHBOARD_URL, permission: '' },
]

export default function Navbar() {
	// placeholder
	const isAuthenticated = true

	return (
		<>
			<SharedNavbar
				links={isAuthenticated ? ADMIN_LINKS : []}
				renderLink={(l) => (
					<NavbarLink key={l.href} href={l.href} permission={l.permission}>
						{l.label}
					</NavbarLink>
				)}
				rightContent={
					isAuthenticated ? <div>Authenticated</div> : <div>login</div>
				}
			/>
		</>
	)
}
