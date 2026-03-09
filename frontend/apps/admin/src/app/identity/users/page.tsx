'use client'

import { useMemo, useState } from 'react'
import toast from 'react-hot-toast'
import { useRouter } from 'next/navigation'
import {
	ConfirmationModal,
	CreateItemSection,
	PageTitle,
	PermissionGuard,
} from '@pitch-in/shared/components'
import { DataTable } from '@pitch-in/shared/components/ui/table'
import {
	DEFAULT_PAGE_SIZE,
	MODAL_CONFIRMATION_BUTTON_DELETING_STYLE,
	MODAL_CONFIRMATION_BUTTON_STYLE,
	MODAL_CANCEL_BUTTON_STYLE,
	MODAL_CANCEL_DELETING_STYLE,
} from '@pitch-in/shared/constants'
import { useDataTableController } from '@pitch-in/shared/hooks'
import { TableColumn, TableRowAction, User } from '@pitch-in/shared/types'
import { UserAPI } from '@/lib/clients/api'
import { CREATE_USERS_URL, USERS_URL } from '@/lib'

const USER_COLUMNS: TableColumn<User>[] = [
	{
		header: 'ID',
		accessor: 'id',
		sortable: true,
	},
	{
		header: 'Username',
		accessor: 'username',
		sortable: true,
	},
	{
		header: 'Full Name',
		render: (u: User) => `${u.first_name} ${u.last_name}`,
		sortable: true,
		sortField: 'last_name',
	},
	{
		header: 'Is Active',
		render: (u: User) => `${u.is_active ? 'Yes' : 'No'}`,
	},
]

const UsersPage = () => {
	const router = useRouter()
	const [isModalOpen, setIsModalOpen] = useState(false)
	const [deleteUser, setDeleteUser] = useState<User | null>(null)
	const [isDeleting, setIsDeleting] = useState(false)

	const {
		data: users,
		isLoading,
		searchTerm,
		totalCount,
		currentPage,
		pageSize,
		sortField,
		sortDirection,
		handleSearch,
		handlePageChange,
		handlePageSizeChange,
		handleSort,
		loadData,
	} = useDataTableController({
		initialSortField: 'username',
		defaultPageSize: DEFAULT_PAGE_SIZE,
		fetchData: UserAPI.fetch,
	})

	const userColumns = useMemo(() => USER_COLUMNS, [])

	const openModal = (user: User) => {
		setDeleteUser(user)
		setIsModalOpen(true)
	}

	const closeModal = () => {
		setIsModalOpen(false)
		setDeleteUser(null)
	}

	const handleDelete = async () => {
		if (deleteUser) {
			setIsDeleting(true)
			try {
				await UserAPI.delete(deleteUser.id)
				toast.success(`User ${deleteUser.username} has been deleted`)
				loadData()
				closeModal()
			} catch (e: unknown) {
			} finally {
				setIsDeleting(false)
			}
		}
	}

	const actions: TableRowAction<User>[] = [
		{
			label: 'Details',
			onClick: (u) => {
				router.push(`${USERS_URL}/${u.id}`)
			},
			actionType: 'view',
			requiredPermission: '',
		},
		{
			label: 'Edit',
			onClick: (u) => {
				router.push(`${USERS_URL}/${u.id}/edit`)
			},
			actionType: 'edit',
			requiredPermission: '',
		},
		{
			label: 'Delete',
			onClick: openModal,
			actionType: 'delete',
			requiredPermission: '',
		},
	]

	return (
		<PermissionGuard requiredPermission=''>
			<PageTitle>Users</PageTitle>
			<CreateItemSection permission='' href={CREATE_USERS_URL}>
				Create New User
			</CreateItemSection>

			<DataTable
				data={users}
				columns={userColumns}
				rowKey='id'
				actions={actions}
				searchTerm={searchTerm}
				onSearch={handleSearch}
				currentPage={currentPage}
				pageSize={pageSize}
				totalCount={totalCount}
				onPageChange={handlePageChange}
				onPageSizeChange={handlePageSizeChange}
				onSort={handleSort}
				currentSortField={sortField}
				currentSortDirection={sortDirection}
				isLoadingRows={isLoading}
			/>
			<ConfirmationModal
				isOpen={isModalOpen}
				onClose={closeModal}
				onConfirm={handleDelete}
				title='Confirm Delete User'
				message={`Are you sure you want to delete ${deleteUser?.username}?`}
				confirmButtonText={isDeleting ? 'Deleting...' : 'Delete'}
				confirmButtonClass={
					isDeleting
						? MODAL_CONFIRMATION_BUTTON_DELETING_STYLE
						: MODAL_CONFIRMATION_BUTTON_STYLE
				}
				cancelButtonClass={
					isDeleting ? MODAL_CANCEL_DELETING_STYLE : MODAL_CANCEL_BUTTON_STYLE
				}
			/>
		</PermissionGuard>
	)
}

export default UsersPage
