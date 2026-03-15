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
