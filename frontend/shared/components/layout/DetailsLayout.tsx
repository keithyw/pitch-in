import {
	Button,
	ConfirmationModal,
	DetailSection,
	DetailSectionRow,
	PageTitle,
	PermissionGuard,
	ServerErrorMessages,
	SpinnerSection,
} from '@pitch-in/shared/components'

interface HasId {
	id: string | number
}

interface DetailsLayoutProps<T> {
	title: string
	item: T
	details: DetailSectionRow[]
	handleDeleteConfirm: () => Promise<void>
	handleEditClick: () => void
	isLoading: boolean
	isConfirmationModalOpen: boolean
	setIsConfirmationModalOpen: (isOpen: boolean) => void
	error: string | null
	children?: React.ReactNode
}

export const DetailsLayout = <T extends HasId>({
	title,
	item,
	details,
	handleDeleteConfirm,
	handleEditClick,
	isLoading,
	isConfirmationModalOpen,
	setIsConfirmationModalOpen,
	error,
	children,
}: DetailsLayoutProps<T>) => {
	if (isLoading) {
		return <SpinnerSection spinnerMessage='Loading details...' />
	}

	if (error) {
		return <ServerErrorMessages message={error} />
	}

	return (
		<div className='p-4'>
			<div className='mx-auto max-w-2xl bg-white p-8 shadow-md'>
				<PageTitle>{title}</PageTitle>
				{item && <DetailSection rows={details} />}
				{children}
				<PermissionGuard>
					<div className='mt-6 flex justify-end space-x-3'>
						<Button actionType='edit' onClick={handleEditClick}>
							Edit
						</Button>
						<Button
							actionType='delete'
							onClick={() => setIsConfirmationModalOpen(true)}
						>
							Delete
						</Button>
					</div>
				</PermissionGuard>
			</div>
			<ConfirmationModal
				isOpen={isConfirmationModalOpen}
				onClose={() => setIsConfirmationModalOpen(false)}
				onConfirm={handleDeleteConfirm}
				title='Confirm Deletion'
				message='Are you sure you want to delete this item?'
			/>
		</div>
	)
}
