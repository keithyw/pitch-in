'use client'

import { useCallback, useEffect, useMemo, useState } from 'react'
import { useParams } from 'next/navigation'
import {
	Button,
	DetailsLayout,
	DetailSectionRow,
	ViewDetailsTable,
} from '@pitch-in/shared/components'
import { FAILED_LOADING_USER_ERROR } from '@pitch-in/shared/constants'
import { Role, TableColumn, User } from '@pitch-in/shared/types'
import {
	useDetailsController,
	useDualListController,
} from '@pitch-in/shared/hooks'
import UserRolesDrawer from '@/components/ui/drawers/UserRolesDrawer'
import { USERS_URL } from '@/lib'
import { RoleAPI, UserAPI } from '@/lib/clients/api'

const COLUMNS: TableColumn<Role>[] = [
	{
		header: 'ID',
		accessor: 'id',
		sortable: true,
	},
	{
		header: 'Name',
		accessor: 'name',
		sortable: true,
	},
]

const UserDetailsPage = () => {
	const params = useParams()
	const [details, setDetails] = useState<DetailSectionRow[]>([])
	const [availableRoles, setAvailableRoles] = useState<Role[]>([])
	const [isRolesLoading, setIsRolesLoading] = useState<boolean>(false)
	const [isDrawerOpen, setIsDrawerOpen] = useState<boolean>(false)

	const cols = useMemo(() => COLUMNS, [])

	const detailsCallback = useCallback((res: User) => {
		setDetails([
			{
				label: 'Username',
				value: res.username,
			},
			{
				label: 'Email',
				value: res.email,
			},
			{
				label: 'First Name',
				value: res.first_name,
			},
			{
				label: 'Last Name',
				value: res.last_name,
			},
			{
				label: 'Is Active',
				value: res.is_active ? 'Yes' : 'No',
			},
		])
	}, [])

	const {
		data: user,
		isLoading,
		error,
		handleDeleteConfirm,
		handleEditClick,
		isConfirmationModalOpen,
		setIsConfirmationModalOpen,
		setData,
	} = useDetailsController({
		id: parseInt(params.id as string),
		deleteData: UserAPI.delete,
		getData: UserAPI.get,
		redirectUrl: USERS_URL,
		errorLoadingMessage: FAILED_LOADING_USER_ERROR,
		handleDetailsCallback: detailsCallback,
	})

	useEffect(() => {
		if (!user?.id || isLoading) return
		const fetchData = async () => {
			setIsRolesLoading(true)
			try {
				const roles = await RoleAPI.fetch({ page: 1 })
				setAvailableRoles(roles.results)
			} catch (e: unknown) {
				console.error('Failure loading data: ', e)
			} finally {
				setIsRolesLoading(false)
			}
		}
		fetchData()
	}, [user?.id, isLoading])

	const dualListController = useDualListController<Role, User>({
		isLoading,
		item: user as User,
		addItems: availableRoles,
		removeItems: user?.roles ?? [],
		getLabel: (item: Role) => {
			return item.name
		},
		attachFn: UserAPI.attach,
		detachFn: UserAPI.detach,
		onUpdate: (roles: Role[]) => {
			setData({ ...(user as User), roles })
		},
	})

	return (
		<DetailsLayout
			title='User Details'
			item={user as User}
			details={details}
			handleDeleteConfirm={handleDeleteConfirm}
			handleEditClick={handleEditClick}
			isLoading={isLoading}
			isConfirmationModalOpen={isConfirmationModalOpen}
			setIsConfirmationModalOpen={setIsConfirmationModalOpen}
			error={error}
			buttons={
				<Button actionType='view' onClick={() => setIsDrawerOpen(true)}>
					Manage Roles
				</Button>
			}
			popups={
				<UserRolesDrawer
					isLoading={isRolesLoading}
					isOpen={isDrawerOpen}
					onClose={() => setIsDrawerOpen(false)}
					availableRoles={dualListController.addOptions}
					assignedRoles={dualListController.removeOptions}
					user={user as User}
					onAttach={dualListController.handleAddItem}
					onDetach={dualListController.handleRemoveItem}
				/>
			}
		>
			<ViewDetailsTable
				data={(user as User)?.roles ?? []}
				columns={cols}
				rowKey='id'
			/>
		</DetailsLayout>
	)
}

export default UserDetailsPage
