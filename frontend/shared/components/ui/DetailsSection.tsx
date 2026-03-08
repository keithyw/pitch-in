import { AssetPreview } from '@pitch-in/shared/components'

export interface DetailSectionRow {
	label: string
	value: string
	isAsset?: boolean
	type?: string
}

interface DetailSectionProps {
	rows: DetailSectionRow[]
}

export const DetailSection = ({ rows }: DetailSectionProps) => {
	return (
		<div className='border-t border-gray-200 px-4 py-5 sm:p-0'>
			<dl className='sm-divide-y sm:divide-gray-200'>
				{rows.map((row, idx) => (
					<div
						key={idx}
						className='py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6 sm:py-5'
					>
						<dt className='text-sm font-medium text-gray-500'>{row.label}</dt>
						<dd className='mt-1 overflow-x-hidden text-sm whitespace-nowrap text-gray-900 sm:col-span-2 sm:mt-0'>
							{row?.isAsset ? (
								<AssetPreview
									url={row.value}
									type={row?.type || 'image'}
									alt={row.label || 'Asset preview'}
								/>
							) : row.value ? (
								row.value
							) : (
								''
							)}
						</dd>
					</div>
				))}
			</dl>
		</div>
	)
}

export default DetailSection
