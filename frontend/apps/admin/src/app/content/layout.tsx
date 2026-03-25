import { CrudLayout, SubnavBarLink } from '@pitch-in/shared/components'
import { CONTENT_WRITE } from '@pitch-in/shared/constants'
import { CONTENT_URL, ITEMS_URL } from '@/lib/constants'

const ContentLayout = ({
	children,
}: Readonly<{ children: React.ReactNode }>) => {
	const links: SubnavBarLink[] = [
		{
			href: ITEMS_URL,
			label: 'Items',
			permission: CONTENT_WRITE,
		},
	]

	return (
		<CrudLayout title='Content Dashboard' links={links}>
			{children}
		</CrudLayout>
	)
}

export default ContentLayout
