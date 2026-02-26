'use client'

interface PermissionGuardProps {
	children: React.ReactNode
}

export const PermissionGuard = ({ children }: PermissionGuardProps) => {
	return <>{children}</>
}
