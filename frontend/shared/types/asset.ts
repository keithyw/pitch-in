export interface Asset {
	id: number
	object_key: string
	mime_type: string
	width: number
	height: number
	size_bytes: number
	created_at: string
	updated_at: string
	url: string
}

export interface CreateAssetRequest {
	object_key: string
	width?: number
	height?: number
	file: File
}
