import { AxiosInstance } from 'axios'
import { API_PERMISSIONS_URL } from '@pitch-in/shared/constants'
import {
	FetchParams,
	ListResponse,
	CreatePermissionRequest,
	Permission,
} from '@pitch-in/shared/types'
import { prepareQueryParams } from '@pitch-in/shared/utils'

interface PermissionService {
	create: (data: CreatePermissionRequest) => Promise<Permission>
	delete: (id: number) => Promise<void>
	fetch: (params: FetchParams) => Promise<ListResponse<Permission>>
	get: (id: number) => Promise<Permission>
	patch: (
		id: number,
		data: Partial<CreatePermissionRequest>,
	) => Promise<Permission>
}

export const permissionService = (
	client: AxiosInstance,
): PermissionService => ({
	create: async (data: CreatePermissionRequest): Promise<Permission> => {
		const r = await client.post<Permission>(API_PERMISSIONS_URL, data)
		return r.data || ({} as Permission)
	},
	delete: async (id: number): Promise<void> => {
		await client.delete(`${API_PERMISSIONS_URL}/${id}`)
		return
	},
	fetch: async (params: FetchParams): Promise<ListResponse<Permission>> => {
		const p = prepareQueryParams(params)
		const r = await client.get<ListResponse<Permission>>(API_PERMISSIONS_URL, {
			params: p,
		})
		return r.data || { results: [], count: 0 }
	},
	get: async (id: number): Promise<Permission> => {
		const r = await client.get<Permission>(`${API_PERMISSIONS_URL}/${id}`)
		return r.data || ({} as Permission)
	},
	patch: async (
		id: number,
		data: Partial<CreatePermissionRequest>,
	): Promise<Permission> => {
		const r = await client.patch<Permission>(
			`${API_PERMISSIONS_URL}/${id}`,
			data,
		)
		return r.data || ({} as Permission)
	},
})
