import { z } from 'zod'

export const userCreateSchema = z.object({
	username: z.string().min(3, 'Username must be at least characters long'),
	email: z.email('Invalid email address'),
	first_name: z.string().max(100),
	last_name: z.string().max(100),
	is_active: z.boolean(),
})

export type UserCreateFormData = z.infer<typeof userCreateSchema>
