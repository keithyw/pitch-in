import { AxiosInstance } from 'axios'
import { API_LOGIN_URL, API_REFRESH_URL } from '@pitch-in/shared/constants'
import { LoginResponse, RefreshResponse } from '@pitch-in/shared/types'

interface AuthService {
	login: (email: string, pass: string) => Promise<LoginResponse>
	refresh: (refreshToken: string) => Promise<RefreshResponse>
}

export const authService = (client: AxiosInstance): AuthService => ({
	login: async (email: string, pass: string): Promise<LoginResponse> => {
		const r = await client.post<LoginResponse>(API_LOGIN_URL, {
			email,
			password: pass,
		})
		return r.data
	},
	refresh: async (refreshToken: string): Promise<RefreshResponse> => {
		const r = await client.post<RefreshResponse>(API_REFRESH_URL, {
			refreshToken,
		})
		return r.data
	},
})
