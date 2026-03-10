import { z } from 'zod'

export const roleCreateSchema = z.object({
	name: z
		.string()
		.min(3, 'Name must be at least 3 characters long')
		.max(255, 'Namecannot exceed 255 characters'),
	description: z.string().optional().or(z.literal('')),
})

export type RoleCreateFormData = z.infer<typeof roleCreateSchema>
