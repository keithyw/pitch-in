export interface Permission {
	id: number
	code: string
	display_name: string
	path: string
	method: string
	created_at: string
	updated_at: string
}

export interface CreatePermissionRequest {
	code: string
	display_name?: string
	path?: string
	method?: string
}
