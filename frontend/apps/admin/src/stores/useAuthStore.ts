import { create } from 'zustand'
import { RefreshResponse, User } from '@pitch-in/shared/types'
import { ACCESS_TOKEN_KEY, REFRESH_TOKEN_KEY } from '@/lib'

const USER_DATA_KEY = 'user_data'

interface AuthStore {
	accessToken: string
	refreshToken: string
	isAuthenticated: boolean
	user: User | null
	setLoginStatus: (res: RefreshResponse) => void
	setLogoutStatus: () => void
	setUser: (user: User) => void
}

const useAuthStore = create<AuthStore>((set, get) => ({
	accessToken: '',
	refreshToken: '',
	isAuthenticated: false,
	user: null,
	setLoginStatus: (res: RefreshResponse) => {
		set({
			accessToken: res.token,
			refreshToken: res.refresh,
			isAuthenticated: true,
		})
		localStorage.setItem(ACCESS_TOKEN_KEY, res.token)
		localStorage.setItem(REFRESH_TOKEN_KEY, res.refresh)
	},
	setLogoutStatus: () => {
		set({
			accessToken: '',
			refreshToken: '',
			isAuthenticated: false,
			user: null,
		})
		localStorage.removeItem(ACCESS_TOKEN_KEY)
		localStorage.removeItem(REFRESH_TOKEN_KEY)
		localStorage.removeItem(USER_DATA_KEY)
	},
	setUser: (user: User) => {
		set({ user })
		localStorage.setItem(USER_DATA_KEY, JSON.stringify(user))
	},
}))

export default useAuthStore
