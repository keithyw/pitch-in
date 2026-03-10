import {
	IdentificationIcon,
	ShieldExclamationIcon,
	UsersIcon,
} from '@heroicons/react/24/outline'
import { DashboardGrid } from '@pitch-in/shared/components'
import { PERMISSIONS_URL, ROLES_URL, USERS_URL } from '@/lib'
import { DashboardCardProps } from '@pitch-in/shared/types'

const IdentityDashboardPage = () => {
	const iconClass = 'mr-2 h-5 w-5 text-blue-600'

	const cards: DashboardCardProps[] = [
		{
			title: 'Permissions',
			description: 'Manage permissions',
			icon: <ShieldExclamationIcon className={iconClass} />,
			link: PERMISSIONS_URL,
		},
		{
			title: 'Roles',
			description: 'Manage roles',
			icon: <IdentificationIcon className={iconClass} />,
			link: ROLES_URL,
		},
		{
			title: 'Users',
			description: 'Manage system users',
			icon: <UsersIcon className={iconClass} />,
			link: USERS_URL,
		},
	]

	return (
		<div className='rounded-xl bg-white p-6 shadow-lg'>
			<DashboardGrid cards={cards} />
		</div>
	)
}

export default IdentityDashboardPage
