import { CrudLayout, SubnavBarLink } from '@pitch-in/shared/components'
import { IDENTITY_WRITE } from '@pitch-in/shared'
import { PERMISSIONS_URL, ROLES_URL, USERS_URL } from '@/lib'

const IdentityLayout = ({
	children,
}: Readonly<{ children: React.ReactNode }>) => {
	const links: SubnavBarLink[] = [
		{
			href: PERMISSIONS_URL,
			label: 'Permissions',
			permission: IDENTITY_WRITE,
		},
		{
			href: ROLES_URL,
			label: 'Roles',
			permission: IDENTITY_WRITE,
		},
		{
			href: USERS_URL,
			label: 'Users',
			permission: IDENTITY_WRITE,
		},
	]

	return (
		<CrudLayout title='Identity Dashboard' links={links}>
			{children}
		</CrudLayout>
	)
}

export default IdentityLayout
