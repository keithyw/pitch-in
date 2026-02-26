import { UsersIcon } from '@heroicons/react/24/outline'
import { DashboardGrid } from '@pitch-in/shared/components'
import { USERS_URL } from '@/lib'
import { DashboardCardProps } from '@pitch-in/shared/types'

const IdentityDashboardPage = () => {
	const cards: DashboardCardProps[] = [
		{
			title: 'Users',
			description: 'Manage system users',
			icon: <UsersIcon className='mr-2 h-5 w-5 text-blue-600' />,
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
