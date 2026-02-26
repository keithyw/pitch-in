'use client'

import {
	Menu,
	MenuButton,
	MenuItem,
	MenuItems,
	Transition,
} from '@headlessui/react'
import {
	ArrowLeftStartOnRectangleIcon,
	UserCircleIcon,
	ChevronDownIcon,
	Cog6ToothIcon,
} from '@heroicons/react/24/outline'
import { useRouter } from 'next/navigation'

const ProfileDropdown = () => {
	const router = useRouter()
	return (
		<Menu as='div' className='relative z-10 inline-block text-left'>
			<div>
				<MenuButton className='inline-flex items-center justify-center gap-x-1.5 rounded-full bg-blue-600 px-4 py-2 text-sm font-semibold text-white shadow-sm transition-colors duration-200 hover:bg-blue-700 focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 focus:ring-offset-gray-800 focus:outline-none'>
					<UserCircleIcon className='-ml-0.5 h-5 w-5' aria-hidden='true' />
					Profile
					<ChevronDownIcon
						className='-mr-1 h-5 text-gray-500'
						aria-hidden='true'
					/>
				</MenuButton>
			</div>
		</Menu>
	)
}

export default ProfileDropdown
