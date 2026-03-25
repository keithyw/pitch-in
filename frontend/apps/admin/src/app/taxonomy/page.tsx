import { TagIcon } from '@heroicons/react/24/outline'
import { DashboardGrid } from '@pitch-in/shared/components'
import { DashboardCardProps } from '@pitch-in/shared/types'
import { DASHBOARD_WRAPPER, ICON_CLASS, TAGS_URL } from '@/lib/constants'

const TaxonomyDashboardPage = () => {
	const cards: DashboardCardProps[] = [
		{
			title: 'Tags',
			description: 'Manage Tags',
			icon: <TagIcon className={ICON_CLASS} />,
			link: TAGS_URL,
		},
	]

	return (
		<div className={DASHBOARD_WRAPPER}>
			<DashboardGrid cards={cards} />
		</div>
	)
}

export default TaxonomyDashboardPage
