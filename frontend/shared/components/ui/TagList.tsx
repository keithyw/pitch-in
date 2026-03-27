import Link from 'next/link'
import { Tag } from '@pitch-in/shared/types'

interface TagListProps {
	tags: Tag[]
	label?: string
	tagUrl?: string
}

export const TagList = ({ tags, label, tagUrl }: TagListProps) => {
	if (!tags || tags.length === 0) return null

	return (
		<div className='rounded-lg border border-gray-200 bg-gray-50 p-4 shadow-sm'>
			<h3 className='mb-3 text-sm font-semibold tracking-wider text-gray-500 uppercase'>
				{label}
			</h3>
			<div className='flex flex-wrap gap-2'>
				{tags.map((tag) =>
					tagUrl ? (
						<Link
							key={tag.id}
							href={`${tagUrl}/${tag.id}`}
							className='inline-flex items-center rounded-md border border-blue-200 bg-white px-2.5 py-1.5 text-sm font-medium text-blue-700 transition-colors hover:bg-blue-50'
						>
							{tag.tag}
						</Link>
					) : (
						<span
							key={tag.id}
							className='inline-flex items-center rounded-md border border-gray-200 bg-white px-2.5 py-1.5 text-sm font-medium text-gray-700'
						>
							{tag.tag}
						</span>
					),
				)}
			</div>
		</div>
	)
}
