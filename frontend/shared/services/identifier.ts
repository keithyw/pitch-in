import { AxiosError } from 'axios'
import { AxiosInstance } from 'axios'
import { API_IDENTIFIER_URL } from '@pitch-in/shared/constants'
import { AIServiceException, IdentifierResponse } from '@pitch-in/shared/types'

interface IdentifierService {
	identify: (prompt: string, image: Blob) => Promise<IdentifierResponse>
}

export const identifierService = (
	client: AxiosInstance,
): IdentifierService => ({
	identify: async (
		prompt: string,
		image: Blob,
	): Promise<IdentifierResponse> => {
		const formData = new FormData()
		formData.append('prompt', prompt)
		formData.append('file', image)
		try {
			const res = await client.post<IdentifierResponse>(
				API_IDENTIFIER_URL,
				formData,
				{
					headers: {
						'Content-Type': 'multipart/form-data',
					},
					timeout: 60000,
				},
			)
			return res.data || ({} as IdentifierResponse)
		} catch (e: unknown) {
			if (e instanceof AxiosError) {
				if (e.response) {
					throw new AIServiceException(
						e.response.data.message,
						e.response.status,
					)
				}
			}
		}
		throw new AIServiceException('Unknown error', 500)
	},
})
