import {
	ConfirmationModal,
	CreateItemSection,
	DataTable,
	PageTitle,
	PermissionGuard,
} from '@pitch-in/shared/components'
import {
	MODAL_CONFIRMATION_BUTTON_STYLE,
	MODAL_CONFIRMATION_BUTTON_DELETING_STYLE,
	MODAL_CANCEL_BUTTON_STYLE,
	MODAL_CANCEL_DELETING_STYLE,
} from '@pitch-in/shared/constants'
import { TableColumn, TableRowAction } from '@pitch-in/shared/types'

interface ListLayoutProps<T> {
	title: string
	createPermission: string
	listPermission: string
	createText: string
	createUrl: string
	data: T[]
	columns: TableColumn<T>[]
	actions: TableRowAction<T>[]
	isLoading: boolean
	tableController: any
	deleteController: any
	deleteTitle: string
	deleteMessage: (item: T | null) => string
}

export const ListLayout = <T extends { id: number }>({
	title,
	createPermission,
	listPermission,
	createText,
	createUrl,
	data,
	columns,
	actions,
	isLoading,
	tableController,
	deleteController,
	deleteTitle,
	deleteMessage,
}: ListLayoutProps<T>) => {
	return (
		<PermissionGuard requiredPermission={listPermission}>
			<PageTitle>{title}</PageTitle>
			<CreateItemSection permission={createPermission} href={createUrl}>
				{createText}
			</CreateItemSection>
			<DataTable
				data={data}
				columns={columns}
				rowKey='id'
				actions={actions}
				searchTerm={tableController.searchTerm}
				onSearch={tableController.handleSearch}
				currentPage={tableController.currentPage}
				pageSize={tableController.pageSize}
				totalCount={tableController.totalCount}
				onPageChange={tableController.handlePageChange}
				onPageSizeChange={tableController.handlePageSizeChange}
				onSort={tableController.handleSort}
				currentSortField={tableController.sortField}
				currentSortDirection={tableController.sortDirection}
				isLoadingRows={isLoading}
			/>
			<ConfirmationModal
				isOpen={deleteController.isModalOpen}
				onClose={deleteController.closeDeleteModal}
				onConfirm={deleteController.handleDelete}
				title={deleteTitle}
				message={deleteMessage(deleteController.deleteItem)}
				confirmButtonText={
					deleteController.isDeleting ? 'Deleting...' : 'Delete'
				}
				confirmButtonClass={
					deleteController.isDeleting
						? MODAL_CONFIRMATION_BUTTON_DELETING_STYLE
						: MODAL_CONFIRMATION_BUTTON_STYLE
				}
				cancelButtonClass={
					deleteController.isDeleting
						? MODAL_CANCEL_DELETING_STYLE
						: MODAL_CANCEL_BUTTON_STYLE
				}
			/>
		</PermissionGuard>
	)
}
