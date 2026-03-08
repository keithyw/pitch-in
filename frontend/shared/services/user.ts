import { AxiosInstance } from 'axios'
import { API_USERS_URL } from '@pitch-in/shared/constants'
import {
	FetchParams,
	ListResponse,
	CreateUserRequest,
	User,
} from '@pitch-in/shared/types'
import { prepareQueryParams } from '@pitch-in/shared/utils'

interface UserService {
	create: (data: CreateUserRequest) => Promise<User>
	delete: (id: number) => Promise<void>
	fetch: (params: FetchParams) => Promise<ListResponse<User>>
	get: (id: number) => Promise<User>
	patch: (id: number, data: Partial<CreateUserRequest>) => Promise<User>
}

export const userService = (client: AxiosInstance): UserService => ({
	create: async (data: CreateUserRequest): Promise<User> => {
		const r = await client.post<User>(API_USERS_URL, data)
		return r.data || ({} as User)
	},
	delete: async (id: number): Promise<void> => {
		await client.delete(`${API_USERS_URL}/${id}`)
		return
	},
	fetch: async (params: FetchParams): Promise<ListResponse<User>> => {
		const p = prepareQueryParams(params)
		const r = await client.get<ListResponse<User>>(API_USERS_URL, {
			params: p,
		})
		return r.data || { results: [], count: 0 }
	},
	get: async (id: number): Promise<User> => {
		const r = await client.get<User>(`${API_USERS_URL}/${id}`)
		return r.data || ({} as User)
	},
	patch: async (
		id: number,
		data: Partial<CreateUserRequest>,
	): Promise<User> => {
		const r = await client.patch<User>(`${API_USERS_URL}/${id}`, data)
		return r.data || ({} as User)
	},
})

// 	r.Patch("/{userID}", middleware.DecodeAndValidate(h.Patch))
