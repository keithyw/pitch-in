import { User } from '@pitch-in/shared/types'

export interface LoginResponse {
	token: string
	refresh: string
	user: User
}

export interface RefreshResponse {
	token: string
	refresh: string
}
