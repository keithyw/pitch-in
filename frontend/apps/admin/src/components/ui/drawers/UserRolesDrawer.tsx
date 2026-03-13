import { DualListSelector, SlideOverDrawer } from '@pitch-in/shared/components'
import { OptionType, User } from '@pitch-in/shared/types'

interface UserRolesDrawerProps {
	isLoading: boolean
	isOpen: boolean
	onClose: () => void
	availableRoles: OptionType[]
	assignedRoles: OptionType[]
	user: User
	onAttach: (id: number) => Promise<void>
	onDetach: (id: number) => Promise<void>
}

const UserRolesDrawer = ({
	isLoading,
	isOpen,
	onClose,
	availableRoles,
	assignedRoles,
	user,
	onAttach,
	onDetach,
}: UserRolesDrawerProps) => {
	if (!user || isLoading) return null
	return (
		<SlideOverDrawer
			title={`Manage Roles for ${user.first_name} ${user.last_name}`}
			isOpen={isOpen}
			onClose={onClose}
			panelWidthClass='max-w-2xl'
		>
			<DualListSelector
				addTitle='Available Roles'
				removeTitle='Included Roles'
				addItems={availableRoles}
				removeItems={assignedRoles}
				onAdd={(id) => onAttach(parseInt(id))}
				onRemove={(id) => onDetach(parseInt(id))}
			/>
		</SlideOverDrawer>
	)
}

export default UserRolesDrawer
