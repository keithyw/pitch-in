import { z } from 'zod'

export const permissionCreateSchema = z.object({
	code: z
		.string()
		.min(3, 'Code must be at least 3 characters long')
		.max(150, 'Code cannot exceed 150 charactrs'),
	display_name: z.string().optional().or(z.literal('')),
	path: z.string().optional().or(z.literal('')),
	method: z.string().optional().or(z.literal('')),
})

export type PermissionCreateFormData = z.infer<typeof permissionCreateSchema>
