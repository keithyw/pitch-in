import React, { Fragment } from 'react'
import {
	Menu,
	MenuButton,
	MenuItem,
	MenuItems,
	Transition,
} from '@headlessui/react'
import {
	Bars3Icon,
	PencilIcon,
	EyeIcon,
	ShieldCheckIcon,
	TrashIcon,
	UserGroupIcon,
} from '@heroicons/react/24/outline'
import type { TableRowAction, TableRowActionType } from '@pitch-in/shared/types'
import { useSharedPermissions } from '@pitch-in/shared/contexts'

const ICON_MAP: Record<TableRowActionType, React.ReactNode> = {
	view: <EyeIcon className='mr-2 h-5 w-5' />,
	edit: <PencilIcon className='mr-2 h-5 w-5' />,
	delete: <TrashIcon className='mr-2 h-5 w-5 text-red-500' />,
	permissionGroup: <ShieldCheckIcon className='mr-2 h-5 w-5' />,
	userGroup: <UserGroupIcon className='mr-2 h-5 w-5' />,
	custom: null,
}

export function RowActionsMenu<T>({
	actions,
	row,
}: {
	actions: TableRowAction<T>[]
	row: T
}) {
	const [openUpwards, setOpenUpwards] = React.useState(false)
	const buttonRef = React.useRef<HTMLButtonElement>(null)
	const { checkAccess } = useSharedPermissions()

	const handleOpen = () => {
		if (buttonRef.current) {
			const rect = buttonRef.current.getBoundingClientRect()
			const spaceBelow = window.innerHeight - rect.bottom
			const menuHeight = 240
			setOpenUpwards(spaceBelow < menuHeight)
		}
	}

	// Filter actions based on permissions
	const filteredActions = actions.filter((action) => {
		if (action.canDisplay && !action.canDisplay(row)) {
			return false
		}
		// If no permission requirements, show the action
		if (
			!action.requiredPermission &&
			!action.requiredPermissions &&
			!action.anyPermission &&
			!action.requiredRole &&
			!action.requiredRoles &&
			!action.anyRole
		) {
			return true
		}

		// Check permissions using the checkAccess function
		return checkAccess({
			requiredPermission: action.requiredPermission,
			requiredPermissions: action.requiredPermissions,
			anyPermission: action.anyPermission,
			requiredRole: action.requiredRole,
			requiredRoles: action.requiredRoles,
			anyRole: action.anyRole,
		})
	})

	if (!filteredActions || filteredActions.length === 0) return null

	return (
		<Menu as='div' className='relative inline-block text-left'>
			<MenuButton
				ref={buttonRef}
				onClick={handleOpen}
				className='rounded p-2 hover:bg-gray-100'
			>
				<Bars3Icon className='h-5 w-5' aria-hidden='true' />
				<span className='sr-only'>Open actions menu</span>
			</MenuButton>
			<Transition
				as={Fragment}
				enter='transition ease-out duration-100'
				enterFrom='transform opacity-0 scale-95'
				enterTo='transform opacity-100 scale-100'
				leave='transition ease-in duration-75'
				leaveFrom='transform opacity-100 scale-100'
				leaveTo='transform opacity-0 scale-95'
			>
				<MenuItems
					className={`ring-opacity-5 absolute right-0 z-50 max-h-60 w-44 origin-top-right overflow-auto rounded-md bg-white shadow-lg ring-1 ring-black focus:outline-none ${openUpwards ? 'bottom-full mb-2' : 'mt-2'} `}
				>
					{filteredActions.map((action, idx) => (
						<MenuItem
							key={idx}
							as='button'
							className={({ active }) =>
								`${active ? 'bg-gray-100' : ''} flex w-full items-center px-4 py-2 text-sm text-gray-700 transition-colors duration-200 hover:bg-gray-100 focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2 focus:outline-none`
							}
							onClick={() => action.onClick(row)}
						>
							{action.icon ||
								(action.actionType &&
									ICON_MAP[action.actionType as TableRowActionType])}
							{action.label}
						</MenuItem>
					))}
				</MenuItems>
			</Transition>
		</Menu>
	)
}
