import { create } from 'zustand'
import {
	RefreshResponse,
	AuthUser,
	PermissionCheck,
} from '@pitch-in/shared/types'
import { ACCESS_TOKEN_KEY, REFRESH_TOKEN_KEY } from '@/lib/constants'

const USER_DATA_KEY = 'user_data'

interface AuthStore {
	accessToken: string
	refreshToken: string
	isAuthenticated: boolean
	isLoading: boolean
	user: AuthUser
	checkAccess: (checks: PermissionCheck) => boolean
	checkIfAuthenticated: () => void
	hasPermission: (code: string) => boolean
	hasAnyPermission: (codes: string[]) => boolean
	hasRole: (role: string) => boolean
	hasAnyRole: (roles: string[]) => boolean
	setLoginStatus: (res: RefreshResponse) => void
	setLogoutStatus: () => void
	setUser: (user: AuthUser) => void
}

const useAuthStore = create<AuthStore>((set, get) => ({
	accessToken: '',
	refreshToken: '',
	isAuthenticated: false,
	isLoading: true,
	user: {
		id: 0,
		username: '',
		roles: [],
		permissions: [],
	},
	checkAccess: (checks: PermissionCheck): boolean => {
		const { user, isAuthenticated } = get()
		if (!isAuthenticated || !user) return false
		if (user.roles?.includes('admin')) return true
		if (
			checks.requiredPermission &&
			!user.permissions?.includes(checks.requiredPermission)
		)
			return false
		if (checks.requiredRole && !user.roles?.includes(checks.requiredRole))
			return false
		return true
	},
	checkIfAuthenticated: (): void => {
		set({ isLoading: true })
		const accessToken = localStorage.getItem(ACCESS_TOKEN_KEY)
		const refreshToken = localStorage.getItem(REFRESH_TOKEN_KEY)

		if (accessToken && refreshToken) {
			const user = localStorage.getItem(USER_DATA_KEY)
			set({
				accessToken,
				refreshToken,
				isAuthenticated: true,
				isLoading: false,
				user: user ? JSON.parse(user) : null,
			})
		} else {
			set({
				isAuthenticated: false,
				isLoading: false,
			})
		}
	},
	hasPermission: (code: string): boolean => {
		const { user } = get()
		return user.permissions.some((p) => p === code)
	},
	hasAnyPermission: (codes: string[]): boolean => {
		const { user } = get()
		return codes.some((code) => user.permissions.some((p) => p === code))
	},
	hasRole: (role: string): boolean => {
		const { user } = get()
		return user.roles.some((r) => r === role)
	},
	hasAnyRole: (roles: string[]): boolean => {
		const { user } = get()
		return roles.some((role) => user.roles.some((r) => r === role))
	},
	setLoginStatus: (res: RefreshResponse) => {
		set({
			accessToken: res.token,
			refreshToken: res.refresh,
			isAuthenticated: true,
			isLoading: false,
		})
		localStorage.setItem(ACCESS_TOKEN_KEY, res.token)
		localStorage.setItem(REFRESH_TOKEN_KEY, res.refresh)
	},
	setLogoutStatus: () => {
		set({
			accessToken: '',
			refreshToken: '',
			isAuthenticated: false,
			isLoading: false,
			user: {
				id: 0,
				username: '',
				roles: [],
				permissions: [],
			},
		})
		localStorage.removeItem(ACCESS_TOKEN_KEY)
		localStorage.removeItem(REFRESH_TOKEN_KEY)
		localStorage.removeItem(USER_DATA_KEY)
	},
	setUser: (user: AuthUser) => {
		set({ user })
		localStorage.setItem(USER_DATA_KEY, JSON.stringify(user))
	},
}))

export default useAuthStore
