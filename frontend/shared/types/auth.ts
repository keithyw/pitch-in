export interface AuthUser {
	id: number
	username: string
	roles: string[]
	permissions: string[]
}
export interface LoginResponse {
	token: string
	refresh: string
	user: AuthUser
}

export interface RefreshResponse {
	token: string
	refresh: string
}

export interface PermissionCheck {
	requiredPermission?: string
	requiredPermissions?: string[] // ALL required
	anyPermission?: string[] // ANY required
	requiredRole?: string
	requiredRoles?: string[] // ALL required
	anyRole?: string[] // ANY required
}
