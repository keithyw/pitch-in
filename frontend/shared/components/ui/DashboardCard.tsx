import Link from 'next/link'
import { DashboardCardProps } from '@pitch-in/shared/types'

export const DashboardCard = ({
	title,
	description,
	icon,
	link = '#',
}: DashboardCardProps) => {
	return (
		<Link
			className='block transform rounded-lg bg-white p-4 text-gray-600 shadow-xl transition-transform duration-200 hover:scale-[1.02]'
			href={link}
		>
			<h3 className='mb-1 flex w-full items-center text-lg font-bold'>
				{icon}
				{title}
			</h3>
			<p className='p-6 text-sm text-gray-400'>{description}</p>
		</Link>
	)
}
