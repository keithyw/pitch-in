import { AxiosInstance } from 'axios'
import { API_INGESTION_URL } from '@pitch-in/shared/constants'
import { IngestionRequest, IngestionResponse } from '@pitch-in/shared/types'

interface IngestionService {
	ingest: (item: IngestionRequest, file: File) => Promise<IngestionResponse>
}

export const ingestionService = (client: AxiosInstance): IngestionService => ({
	ingest: async (
		item: IngestionRequest,
		file: File,
	): Promise<IngestionResponse> => {
		const formData = new FormData()
		formData.append('file', file)
		formData.append('name', item.name)
		formData.append('description', item.description || '')
		item.tags.forEach((tag) => {
			formData.append('tags', tag)
		})
		const res = await client.post<IngestionResponse>(
			API_INGESTION_URL,
			formData,
			{
				headers: {
					'Content-Type': 'multipart/form-data',
				},
			},
		)
		return res.data || ({} as IngestionResponse)
	},
})
