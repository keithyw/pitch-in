export interface User {
	id: number
	username: string
	email: string
	first_name: string
	last_name: string
	is_active: boolean
	created_at: string
	updated_at: string
}

export interface CreateUserRequest {
	username: string
	email: string
	first_name: string
	last_name: string
	is_active: boolean
}
