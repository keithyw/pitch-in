export interface Tag {
	id: number
	tag: string
	slug: string
	created_at: string
	updated_at: string
	deleted_at: string
}

export interface CreateTagRequest {
	tag: string
}
