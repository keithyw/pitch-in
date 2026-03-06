import { createAxiosClient } from '@pitch-in/shared/clients'
import { API_REFRESH_URL, API_LOGIN_URL } from '@pitch-in/shared/constants'
import { RefreshResponse } from '@pitch-in/shared/types'
import {
	ServiceFactory,
	authService,
	userService,
} from '@pitch-in/shared/services'
import { ACCESS_TOKEN_KEY, REFRESH_TOKEN_KEY } from '@/lib'
import useAuthStore from '@/stores/useAuthStore'

const api = createAxiosClient({
	baseUrl: process.env.NEXT_PUBLIC_API_URL,
	refreshUrl: API_REFRESH_URL,
	authUrls: [API_LOGIN_URL, API_REFRESH_URL],
	getToken: () => localStorage.getItem(ACCESS_TOKEN_KEY),
	getRefreshToken: () => localStorage.getItem(REFRESH_TOKEN_KEY),
	onRefreshSuccess: (res: RefreshResponse) => {
		useAuthStore.getState().setLoginStatus(res)
		// TRICK: Move the background user refresh here!
		// This removes the userService import from the client file.
		// triggerUserRefresh();
	},
	onLogout: () => useAuthStore.getState().setLogoutStatus(),
})

const factory = new ServiceFactory(api)

export const AuthAPI = factory.create(authService)
export const UserAPI = factory.create(userService)
