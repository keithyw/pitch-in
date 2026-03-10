import { AxiosInstance } from 'axios'
import { API_ROLES_URL } from '@pitch-in/shared/constants'
import {
	FetchParams,
	ListResponse,
	Role,
	CreateRoleRequest,
} from '@pitch-in/shared/types'
import { prepareQueryParams } from '@pitch-in/shared/utils'

interface RoleService {
	create: (data: CreateRoleRequest) => Promise<Role>
	delete: (id: number) => Promise<void>
	fetch: (params: FetchParams) => Promise<ListResponse<Role>>
	get: (id: number) => Promise<Role>
	patch: (id: number, data: Partial<CreateRoleRequest>) => Promise<Role>
}

export const roleService = (client: AxiosInstance): RoleService => ({
	create: async (data: CreateRoleRequest): Promise<Role> => {
		const r = await client.post<Role>(API_ROLES_URL, data)
		return r.data || ({} as Role)
	},
	delete: async (id: number): Promise<void> => {
		await client.delete(`${API_ROLES_URL}/${id}`)
		return
	},
	fetch: async (params: FetchParams): Promise<ListResponse<Role>> => {
		const p = prepareQueryParams(params)
		const r = await client.get<ListResponse<Role>>(API_ROLES_URL, {
			params: p,
		})
		return r.data || { results: [], count: 0 }
	},
	get: async (id: number): Promise<Role> => {
		const r = await client.get<Role>(`${API_ROLES_URL}/${id}`)
		return r.data || ({} as Role)
	},
	patch: async (
		id: number,
		data: Partial<CreateRoleRequest>,
	): Promise<Role> => {
		const r = await client.patch<Role>(`${API_ROLES_URL}/${id}`, data)
		return r.data || ({} as Role)
	},
})
