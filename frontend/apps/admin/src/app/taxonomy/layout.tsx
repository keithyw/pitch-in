import { CrudLayout, SubnavBarLink } from '@pitch-in/shared/components'
import { CONTENT_WRITE } from '@pitch-in/shared/constants'
import { TAXONOMY_URL, TAGS_URL } from '@/lib'
import React from 'react'

const TaxonomyLayout = ({
	children,
}: Readonly<{ children: React.ReactNode }>) => {
	const links: SubnavBarLink[] = [
		{
			href: TAGS_URL,
			label: 'Tags',
			permission: CONTENT_WRITE,
		},
	]

	return (
		<CrudLayout title='Taxonomy Dashboard' links={links}>
			{children}
		</CrudLayout>
	)
}

export default TaxonomyLayout
