import { z } from 'zod'

export const assetCreateSchema = z.object({
	object_key: z.string(),
	file: z.any().optional(),
})

export type AssetCreateFormData = z.infer<typeof assetCreateSchema>
