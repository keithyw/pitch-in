'use client'

interface PermissionGuardProps {
	requiredPermission?: string
	children: React.ReactNode
}

export const PermissionGuard = ({
	requiredPermission,
	children,
}: PermissionGuardProps) => {
	return <>{children}</>
}
