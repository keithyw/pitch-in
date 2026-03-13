import { DualListSelector, SlideOverDrawer } from '@pitch-in/shared/components'
import { OptionType, Role } from '@pitch-in/shared/types'

interface RolePermissionsDrawerProps {
	isLoading: boolean
	isOpen: boolean
	onClose: () => void
	availablePermissions: OptionType[]
	assignedPermissions: OptionType[]
	role: Role
	onAttach: (id: number) => Promise<void>
	onDetach: (id: number) => Promise<void>
}

const RolePermissionsDrawer = ({
	isLoading,
	isOpen,
	onClose,
	availablePermissions,
	assignedPermissions,
	role,
	onAttach,
	onDetach,
}: RolePermissionsDrawerProps) => {
	if (!role || isLoading) return null
	return (
		<SlideOverDrawer
			title={`Manage Permissions for ${role.name}`}
			isOpen={isOpen}
			onClose={onClose}
			panelWidthClass='max-w-2xl'
		>
			<DualListSelector
				addTitle='Available Permissions'
				removeTitle='Included Permissions'
				addItems={availablePermissions}
				removeItems={assignedPermissions}
				onAdd={(id) => onAttach(parseInt(id))}
				onRemove={(id) => onDetach(parseInt(id))}
			/>
		</SlideOverDrawer>
	)
}

export default RolePermissionsDrawer
