import { Tag } from './tag'
export interface Item {
	id: number
	name: string
	slug: string
	description?: string
	created_at: string
	updated_at: string
	deleted_at: string
	tags: Tag[]
}

export interface CreateItemRequest {
	name: string
	description?: string
}
