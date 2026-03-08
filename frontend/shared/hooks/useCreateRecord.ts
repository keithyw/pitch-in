'use client'

import toast from 'react-hot-toast'
import { useRouter } from 'next/navigation'
import { useForm, DefaultValues, FieldValues } from 'react-hook-form'
import { z } from 'zod'
import { zodResolver } from '@hookform/resolvers/zod'
import { handleFormErrors } from '@pitch-in/shared/utils'

interface UseCreateOptions<T extends z.ZodTypeAny, TResponse> {
	schema: T
	defaultValues: DefaultValues<z.infer<T>>
	createFn: (data: z.infer<T>) => Promise<TResponse>
	redirectUrl?: string
	successMessage?: (res: TResponse) => string
	onSuccess?: (res: TResponse) => void
}

export function useCreateRecord<T extends z.ZodTypeAny, TResponse>({
	schema,
	defaultValues,
	createFn,
	redirectUrl,
	successMessage,
	onSuccess,
}: UseCreateOptions<T, TResponse>) {
	type TFieldValues = z.infer<T> & FieldValues

	const router = useRouter()
	const formMethods = useForm<TFieldValues>({
		resolver: zodResolver(schema as any),
		defaultValues: defaultValues as DefaultValues<TFieldValues>,
	})

	const { handleSubmit, setError, reset } = formMethods

	const onSubmit = async (data: TFieldValues) => {
		try {
			const res = await createFn(data)
			const msg = successMessage ? successMessage(res) : 'Item created'
			toast.success(msg)
			reset()
			if (onSuccess) onSuccess(res)
			if (redirectUrl) router.push(redirectUrl)
		} catch (e: unknown) {
			handleFormErrors<TFieldValues>(
				e,
				setError,
				'Failed to create item. Please review your input.',
			)
		}
	}

	return {
		...formMethods,
		onSubmit: handleSubmit(onSubmit),
	}
}
