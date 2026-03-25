import { AxiosInstance } from 'axios'
import { API_ITEMS_URL } from '@pitch-in/shared/constants'
import {
	FetchParams,
	ListResponse,
	CreateItemRequest,
	Item,
} from '@pitch-in/shared/types'
import { prepareQueryParams } from '@pitch-in/shared/utils'

interface ItemService {
	create: (data: CreateItemRequest) => Promise<Item>
	delete: (id: number) => Promise<void>
	fetch: (params: FetchParams) => Promise<ListResponse<Item>>
	get: (id: number) => Promise<Item>
	patch: (id: number, data: Partial<CreateItemRequest>) => Promise<Item>
}

export const itemService = (client: AxiosInstance): ItemService => ({
	create: async (data: CreateItemRequest): Promise<Item> => {
		const r = await client.post<Item>(API_ITEMS_URL, data)
		return r.data || ({} as Item)
	},
	delete: async (id: number): Promise<void> => {
		await client.delete(`${API_ITEMS_URL}/${id}`)
		return
	},
	fetch: async (params: FetchParams): Promise<ListResponse<Item>> => {
		const p = prepareQueryParams(params)
		const r = await client.get<ListResponse<Item>>(API_ITEMS_URL, {
			params: p,
		})
		return r.data || { results: [], count: 0 }
	},
	get: async (id: number): Promise<Item> => {
		const r = await client.get<Item>(`${API_ITEMS_URL}/${id}`)
		return r.data || ({} as Item)
	},
	patch: async (
		id: number,
		data: Partial<CreateItemRequest>,
	): Promise<Item> => {
		const r = await client.patch<Item>(`${API_ITEMS_URL}/${id}`, data)
		return r.data || ({} as Item)
	},
})
