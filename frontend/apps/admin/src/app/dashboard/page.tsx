import { IdentificationIcon, PhotoIcon } from '@heroicons/react/24/outline'
import { DashboardGrid } from '@pitch-in/shared/components'
import { DashboardCardProps } from '@pitch-in/shared/types'
import {
	ASSETS_URL,
	DASHBOARD_WRAPPER,
	ICON_CLASS,
	IDENTITY_URL,
} from '@/lib/constants'

const DashboardPage = () => {
	const cards: DashboardCardProps[] = [
		{
			title: 'Assets',
			description: 'Manage Assets',
			icon: <PhotoIcon className={ICON_CLASS} />,
			link: ASSETS_URL,
		},
		{
			title: 'Identity',
			description: 'Manage Users, Roles and Permissions',
			icon: <IdentificationIcon className={ICON_CLASS} />,
			link: IDENTITY_URL,
		},
	]

	return (
		<div className={DASHBOARD_WRAPPER}>
			<DashboardGrid cards={cards} />
		</div>
	)
}

export default DashboardPage
