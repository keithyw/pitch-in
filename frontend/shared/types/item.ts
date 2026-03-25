export interface Item {
	id: number
	name: string
	slug: string
	description?: string
	created_at: string
	updated_at: string
	deleted_at: string
}

export interface CreateItemRequest {
	name: string
	description?: string
}
