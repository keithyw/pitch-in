import { AxiosInstance } from 'axios'
import { API_USERS_URL } from '@pitch-in/shared/constants'
import { FetchParams, ListResponse, User } from '@pitch-in/shared/types'
import { prepareQueryParams } from '@pitch-in/shared/utils'

interface UserService {
	fetch: (params: FetchParams) => Promise<ListResponse<User>>
}

// export const authService = (client: AxiosInstance): AuthService => ({

export const userService = (client: AxiosInstance): UserService => ({
	fetch: async (params: FetchParams): Promise<ListResponse<User>> => {
		const p = prepareQueryParams(params)
		const r = await client.get<ListResponse<User>>(API_USERS_URL, {
			params: p,
		})
		return r.data || { results: [], count: 0 }
	},
})

// r.Delete("/{userID}", h.Delete)
// 	r.Get("/{userID}", h.Get)
// 	r.Get("/", h.FindBy)
// 	r.Post("/", middleware.DecodeAndValidate(h.Post))
// 	r.Patch("/{userID}", middleware.DecodeAndValidate(h.Patch))
// login: async (email: string, pass: string): Promise<LoginResponse> => {
// 		const r = await client.post<LoginResponse>(API_LOGIN_URL, {
// 			email,
// 			password: pass,
// 		})
// 		return r.data
// 	},
