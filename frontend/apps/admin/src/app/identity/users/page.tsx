'use client'

import { PageTitle } from '@pitch-in/shared/components'
import { CreateItemSection } from '@pitch-in/shared/components'

const UsersPage = () => {
	return (
		<>
			<PageTitle>Users</PageTitle>
			<CreateItemSection permission='' href=''>
				Create New User
			</CreateItemSection>
		</>
	)
}

export default UsersPage
