import { AxiosInstance } from 'axios'
import { API_ASSETS_URL } from '@pitch-in/shared/constants'
import {
	Asset,
	CreateAssetRequest,
	FetchParams,
	ListResponse,
} from '@pitch-in/shared/types'
import { prepareQueryParams } from '@pitch-in/shared/utils'

interface AssetService {
	createWithFile: (data: CreateAssetRequest) => Promise<Asset>
	delete: (id: number) => Promise<void>
	fetch: (params: FetchParams) => Promise<ListResponse<Asset>>
	get: (id: number) => Promise<Asset>
	patch: (id: number, data: Partial<CreateAssetRequest>) => Promise<Asset>
}

export const assetService = (client: AxiosInstance): AssetService => ({
	createWithFile: async (data: CreateAssetRequest): Promise<Asset> => {
		const formData = new FormData()
		formData.append('file', data.file, data.file.name)
		formData.append('object_key', data.object_key)
		if (data.width) formData.append('width', data.width.toString() as string)
		if (data.height) formData.append('height', data.height.toString() as string)
		const res = await client.post<Asset>(API_ASSETS_URL, formData, {
			headers: {
				'Content-Type': 'multipart/form-data',
			},
		})
		return res.data || ({} as Asset)
	},
	delete: async (id: number): Promise<void> => {
		await client.delete(`${API_ASSETS_URL}/${id}`)
		return
	},
	fetch: async (params: FetchParams): Promise<ListResponse<Asset>> => {
		const p = prepareQueryParams(params)
		const r = await client.get<ListResponse<Asset>>(API_ASSETS_URL, {
			params: p,
		})
		return r.data || { results: [], count: 0 }
	},
	get: async (id: number): Promise<Asset> => {
		const res = await client.get<Asset>(`${API_ASSETS_URL}/${id}`)
		return res.data || ({} as Asset)
	},
	patch: async (
		id: number,
		data: Partial<CreateAssetRequest>,
	): Promise<Asset> => {
		const formData = new FormData()
		if (data.file) formData.append('file', data.file, data.file.name)
		if (data.object_key) formData.append('object_key', data.object_key)
		if (data.width) formData.append('width', data.width.toString() as string)
		if (data.height) formData.append('height', data.height.toString() as string)
		const res = await client.patch<Asset>(`${API_ASSETS_URL}/${id}`, formData, {
			headers: {
				'Content-Type': 'multipart/form-data',
			},
		})
		return res.data || ({} as Asset)
	},
})
