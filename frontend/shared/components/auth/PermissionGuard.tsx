'use client'

import useAuthStore from '@admin/stores/useAuthStore'
import React from 'react'

interface PermissionGuardProps {
	requiredPermission?: string
	requiredPermissions?: string[]
	anyPermission?: string[]
	requiredRole?: string
	requiredRoles?: string[]
	anyRole?: string[]
	children: React.ReactNode
	fallback?: React.ReactNode
}

export const PermissionGuard = ({
	requiredPermission,
	requiredPermissions,
	anyPermission,
	requiredRole,
	requiredRoles,
	anyRole,
	children,
	fallback = null,
}: PermissionGuardProps) => {
	const {
		user,
		hasPermission,
		hasAnyPermission,
		hasRole,
		hasAnyRole,
		isAuthenticated,
	} = useAuthStore()

	if (!isAuthenticated || !user) {
		return <>{fallback}</>
	}

	let hasAccess = true

	if (requiredPermission && !hasPermission(requiredPermission)) {
		hasAccess = false
	}

	if (requiredPermissions && requiredPermissions.length > 0) {
		const hasPerms = requiredPermissions.every((p) => hasPermission(p))
		if (!hasPerms) {
			hasAccess = false
		}
	}

	if (anyPermission && anyPermission.length > 0) {
		if (!hasAnyPermission(anyPermission)) {
			hasAccess = false
		}
	}

	if (requiredRole && !hasRole(requiredRole)) {
		hasAccess = false
	}

	if (requiredRoles && requiredRoles.length > 0) {
		if (!requiredRoles.every((r) => hasRole(r))) {
			hasAccess = false
		}
	}

	if (anyRole && anyRole.length > 0) {
		if (!hasAnyRole(anyRole)) {
			hasAccess = false
		}
	}

	return hasAccess ? <>{children}</> : <>{fallback}</>
}
