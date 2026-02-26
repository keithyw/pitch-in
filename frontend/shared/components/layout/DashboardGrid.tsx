import { DashboardCard } from '@pitch-in/shared/components'
import { DashboardCardProps } from '@pitch-in/shared/types'

export const DashboardGrid = ({ cards }: { cards: DashboardCardProps[] }) => {
	return (
		<div className='grid grid-cols-1 gap-6 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4'>
			{cards.map((card, idx) => (
				<DashboardCard key={idx} {...card} />
			))}
		</div>
	)
}
