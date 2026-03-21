import {
	IdentificationIcon,
	ShieldExclamationIcon,
	UsersIcon,
} from '@heroicons/react/24/outline'
import { DashboardGrid } from '@pitch-in/shared/components'
import {
	DASHBOARD_WRAPPER,
	ICON_CLASS,
	PERMISSIONS_URL,
	ROLES_URL,
	USERS_URL,
} from '@/lib/constants'
import { DashboardCardProps } from '@pitch-in/shared/types'

const IdentityDashboardPage = () => {
	const cards: DashboardCardProps[] = [
		{
			title: 'Permissions',
			description: 'Manage permissions',
			icon: <ShieldExclamationIcon className={ICON_CLASS} />,
			link: PERMISSIONS_URL,
		},
		{
			title: 'Roles',
			description: 'Manage roles',
			icon: <IdentificationIcon className={ICON_CLASS} />,
			link: ROLES_URL,
		},
		{
			title: 'Users',
			description: 'Manage system users',
			icon: <UsersIcon className={ICON_CLASS} />,
			link: USERS_URL,
		},
	]

	return (
		<div className={DASHBOARD_WRAPPER}>
			<DashboardGrid cards={cards} />
		</div>
	)
}

export default IdentityDashboardPage
