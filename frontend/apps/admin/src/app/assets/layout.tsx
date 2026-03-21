import { CrudLayout, SubnavBarLink } from '@pitch-in/shared/components'
import { ASSETS_READ } from '@pitch-in/shared/constants'
import { ASSETS_URL } from '@/lib/constants'

const AssetLayout = ({ children }: Readonly<{ children: React.ReactNode }>) => {
	const links: SubnavBarLink[] = [
		{
			href: ASSETS_URL,
			label: 'Assets',
			permission: ASSETS_READ,
		},
	]

	return (
		<CrudLayout title='Assets Management' links={links}>
			{children}
		</CrudLayout>
	)
}

export default AssetLayout
