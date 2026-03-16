'use client'

import React, { createContext, useContext } from 'react'
import { PermissionCheck } from '@pitch-in/shared/types'

interface PermissionContextType {
	checkAccess: (checks: PermissionCheck) => boolean
}

const PermissionContext = createContext<PermissionContextType | null>(null)

export const PermissionProvider = ({
	children,
	checkAccess,
}: {
	children: React.ReactNode
	checkAccess: (checks: PermissionCheck) => boolean
}) => (
	<PermissionContext.Provider value={{ checkAccess }}>
		{children}
	</PermissionContext.Provider>
)

export const useSharedPermissions = () => {
	const ctx = useContext(PermissionContext)
	if (!ctx)
		throw new Error(
			'useSharedPermissions must be used within PermissionProvider',
		)
	return ctx
}
