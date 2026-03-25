import { Square3Stack3DIcon } from '@heroicons/react/24/outline'
import { DashboardGrid } from '@pitch-in/shared/components'
import { DashboardCardProps } from '@pitch-in/shared/types'
import { DASHBOARD_WRAPPER, ICON_CLASS, ITEMS_URL } from '@/lib/constants'

const ContentDashboardPage = () => {
	const cards: DashboardCardProps[] = [
		{
			title: 'Items',
			description: 'Manage Items',
			icon: <Square3Stack3DIcon className={ICON_CLASS} />,
			link: ITEMS_URL,
		},
	]

	return (
		<div className={DASHBOARD_WRAPPER}>
			<DashboardGrid cards={cards} />
		</div>
	)
}

export default ContentDashboardPage
