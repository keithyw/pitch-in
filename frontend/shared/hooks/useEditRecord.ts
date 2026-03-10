'use client'

import { useCallback, useEffect, useState } from 'react'
import { useForm, Resolver, DefaultValues, FieldValues } from 'react-hook-form'
import { useRouter } from 'next/navigation'
import toast from 'react-hot-toast'
import { zodResolver } from '@hookform/resolvers/zod'
import { z } from 'zod'
import { handleFormErrors } from '@pitch-in/shared/utils'

interface UseEditRecordProps<T extends z.ZodTypeAny, TResponse> {
	id: number
	defaultValues?: Partial<z.infer<T>>
	getData: (id: number) => Promise<TResponse>
	updateData: (id: number, data: z.infer<T>) => Promise<TResponse>
	errorLoadingMessage: string
	redirectUrl: string
	schema: T
	handleFetchCallback: (data: TResponse) => Partial<z.infer<T>>
	transformData: (data: z.infer<T>) => Promise<z.infer<T>>
}

export function useEditRecord<T extends z.ZodTypeAny, TResponse>({
	id,
	defaultValues,
	getData,
	updateData,
	errorLoadingMessage,
	redirectUrl,
	schema,
	handleFetchCallback,
	transformData,
}: UseEditRecordProps<T, TResponse>) {
	type TFieldValues = z.infer<T> & FieldValues
	const router = useRouter()
	const [isLoading, setIsLoading] = useState(false)
	const [loadingError, setLoadingError] = useState<string | null>(null)
	const [data, setData] = useState<TResponse | null>(null)

	const formMethods = useForm<TFieldValues>({
		resolver: zodResolver(schema as any),
		defaultValues: defaultValues as DefaultValues<TFieldValues>,
	})

	const { handleSubmit, setError, reset } = formMethods

	const onSubmit = async (data: TFieldValues) => {
		try {
			const payload = await transformData(data)
			const res = await updateData(id, payload)
			setData(res)
			toast.success(`Item updated successfully!`)
			router.push(redirectUrl)
		} catch (e: unknown) {
			handleFormErrors<TFieldValues>(
				e,
				setError,
				'Failed to edit item. Please review your input.',
			)
		}
	}

	const fetchItem = useCallback(async () => {
		setIsLoading(true)
		setLoadingError(null)
		try {
			const res = await getData(id)
			if (res) {
				setData(res)
				const vals = handleFetchCallback(res)
				reset(vals as DefaultValues<TFieldValues>)
			}
		} catch (e: unknown) {
			if (e instanceof Error) {
				setLoadingError(errorLoadingMessage)
				toast.error(errorLoadingMessage)
				console.error(e.message)
				router.push(redirectUrl)
			}
		} finally {
			setIsLoading(false)
		}
	}, [
		errorLoadingMessage,
		getData,
		handleFetchCallback,
		id,
		redirectUrl,
		reset,
		router,
	])

	useEffect(() => {
		void fetchItem()
	}, [fetchItem])

	return {
		data,
		isLoading,
		fieldErrors: formMethods.formState.errors,
		isSubmitting: formMethods.formState.isSubmitting,
		loadingError,
		register: formMethods.register,
		control: formMethods.control,
		reset,
		onSubmit: handleSubmit(onSubmit),
	}
}
