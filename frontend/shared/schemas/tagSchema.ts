import { z } from 'zod'

export const tagCreateSchema = z.object({
	tag: z
		.string()
		.min(2, 'Tag must be at least 2 characters long')
		.max(255, 'Tag cannot exceed 255 characters'),
})

export type TagCreateFormData = z.infer<typeof tagCreateSchema>
