import { AxiosInstance } from 'axios'
import { API_TAGS_URL } from '@pitch-in/shared/constants'
import {
	FetchParams,
	ListResponse,
	CreateTagRequest,
	Tag,
} from '@pitch-in/shared/types'
import { prepareQueryParams } from '@pitch-in/shared/utils'

interface TagService {
	attach: (tagId: number, entityId: number, entity: string) => Promise<void>
	create: (data: CreateTagRequest) => Promise<Tag>
	delete: (id: number) => Promise<void>
	detach: (tagId: number, entityId: number, entity: string) => Promise<void>
	fetch: (params: FetchParams) => Promise<ListResponse<Tag>>
	get: (id: number) => Promise<Tag>
	patch: (id: number, data: Partial<CreateTagRequest>) => Promise<Tag>
}

export const tagService = (client: AxiosInstance): TagService => ({
	attach: async (
		tagId: number,
		entityId: number,
		entity: string,
	): Promise<void> => {
		await client.post(`${API_TAGS_URL}/${tagId}/entity`, {
			entity,
			entity_id: entityId,
		})
		return
	},
	create: async (data: CreateTagRequest): Promise<Tag> => {
		const r = await client.post<Tag>(API_TAGS_URL, data)
		return r.data || ({} as Tag)
	},
	delete: async (id: number): Promise<void> => {
		await client.delete(`${API_TAGS_URL}/${id}`)
		return
	},
	detach: async (
		tagId: number,
		entityId: number,
		entity: string,
	): Promise<void> => {
		await client.delete(`${API_TAGS_URL}/${tagId}/entity`, {
			data: {
				entity,
				entity_id: entityId,
			},
		})
		return
	},
	fetch: async (params: FetchParams): Promise<ListResponse<Tag>> => {
		const p = prepareQueryParams(params)
		const r = await client.get<ListResponse<Tag>>(API_TAGS_URL, {
			params: p,
		})
		return r.data || { results: [], count: 0 }
	},
	get: async (id: number): Promise<Tag> => {
		const r = await client.get<Tag>(`${API_TAGS_URL}/${id}`)
		return r.data || ({} as Tag)
	},
	patch: async (id: number, data: Partial<CreateTagRequest>): Promise<Tag> => {
		const r = await client.patch<Tag>(`${API_TAGS_URL}/${id}`, data)
		return r.data || ({} as Tag)
	},
})
