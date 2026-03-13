'use client'

import { useCallback, useEffect, useMemo, useState } from 'react'
import { useParams } from 'next/navigation'
import {
	Button,
	DetailsLayout,
	DetailSectionRow,
	ViewDetailsTable,
} from '@pitch-in/shared/components'
import { Permission, Role, TableColumn } from '@pitch-in/shared/types'
import { failedLoadingError } from '@pitch-in/shared/utils'
import { useDetailsController, useDualListController } from '@pitch-in/shared'
import RolePermissionsDrawer from '@/components/ui/drawers/RolePermissionsDrawer'
import { ROLES_URL } from '@/lib'
import { PermissionAPI, RoleAPI } from '@/lib/clients/api'

const COLS: TableColumn<Permission>[] = [
	{
		header: 'ID',
		accessor: 'id',
		sortable: true,
	},
	{
		header: 'Code',
		accessor: 'code',
		sortable: true,
	},
	{
		header: 'Display Name',
		accessor: 'display_name',
		sortable: true,
	},
]

const RoleDetailsPage = () => {
	const params = useParams()
	const [details, setDetails] = useState<DetailSectionRow[]>([])
	const [addPermissions, setAddPermissions] = useState<Permission[]>([])
	const [isPermissionsLoading, setIsPermissionsLoading] =
		useState<boolean>(false)
	const [isDrawerOpen, setIsDrawerOpen] = useState<boolean>(false)

	const cols = useMemo(() => COLS, [])

	const detailsCallback = useCallback((r: Role) => {
		setDetails([
			{
				label: 'Name',
				value: r.name,
			},
			{
				label: 'Description',
				value: r.description,
			},
		])
	}, [])

	const detailsController = useDetailsController({
		id: parseInt(params.id as string),
		deleteData: RoleAPI.delete,
		getData: RoleAPI.get,
		redirectUrl: ROLES_URL,
		errorLoadingMessage: failedLoadingError('role'),
		handleDetailsCallback: detailsCallback,
	})

	const { data, isLoading, setData } = detailsController

	useEffect(() => {
		if (!data?.id || isLoading) return
		const fetchData = async () => {
			setIsPermissionsLoading(true)
			try {
				const perms = await PermissionAPI.fetch({ page: 1 })
				setAddPermissions(perms.results)
			} catch (e: unknown) {
				console.error('Failure loading data: ', e)
			} finally {
				setIsPermissionsLoading(false)
			}
		}
		fetchData()
	}, [data?.id, isLoading])

	const dualListController = useDualListController<Permission, Role>({
		isLoading,
		item: data as Role,
		addItems: addPermissions,
		removeItems: (data as Role)?.permissions ?? [],
		getLabel: (item: Permission) => {
			return item.display_name
		},
		attachFn: RoleAPI.attach,
		detachFn: RoleAPI.detach,
		onUpdate: (perms: Permission[]) => {
			setData({ ...(data as Role), permissions: perms })
		},
	})

	return (
		<DetailsLayout
			title='Role Details'
			item={data as Role}
			details={details}
			handleDeleteConfirm={detailsController.handleDeleteConfirm}
			handleEditClick={detailsController.handleEditClick}
			isLoading={isLoading}
			isConfirmationModalOpen={detailsController.isConfirmationModalOpen}
			setIsConfirmationModalOpen={detailsController.setIsConfirmationModalOpen}
			error={detailsController.error}
			buttons={
				<Button actionType='view' onClick={() => setIsDrawerOpen(true)}>
					Manage Permissions
				</Button>
			}
			popups={
				<RolePermissionsDrawer
					isLoading={isPermissionsLoading}
					isOpen={isDrawerOpen}
					onClose={() => {
						setIsDrawerOpen(false)
					}}
					availablePermissions={dualListController.addOptions}
					assignedPermissions={dualListController.removeOptions}
					role={detailsController.data as Role}
					onAttach={dualListController.handleAddItem}
					onDetach={dualListController.handleRemoveItem}
				/>
			}
		>
			<ViewDetailsTable
				data={(data as Role)?.permissions ?? []}
				columns={cols}
				rowKey='id'
			/>
		</DetailsLayout>
	)
}

export default RoleDetailsPage
