import { CrudLayout } from '@pitch-in/shared/components'
import { SubnavBarLink } from '@pitch-in/shared/components'
import { USERS_URL } from '@/lib'

const IdentityLayout = ({
	children,
}: Readonly<{ children: React.ReactNode }>) => {
	const links: SubnavBarLink[] = [
		{
			href: `${USERS_URL}`,
			label: 'Users',
		},
	]

	return (
		<CrudLayout title='Identity Dashboard' links={links}>
			{children}
		</CrudLayout>
	)
}

export default IdentityLayout
