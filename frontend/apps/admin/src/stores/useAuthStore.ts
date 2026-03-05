import { create } from 'zustand'
import { RefreshResponse, User } from '@pitch-in/shared/types'

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
	},
	setLogoutStatus: () => {
		set({
			accessToken: '',
			refreshToken: '',
			isAuthenticated: false,
		})
	},
	setUser: (user: User) => {
		set({ user })
		localStorage.setItem(USER_DATA_KEY, JSON.stringify(user))
	},
}))

export default useAuthStore
