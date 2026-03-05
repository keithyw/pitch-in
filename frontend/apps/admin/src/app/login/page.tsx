'use client'

import axios, { AxiosError } from 'axios'
import { useState } from 'react'
import { useRouter } from 'next/navigation'
import { useForm, FieldValues } from 'react-hook-form'
import { CardContainer } from '@pitch-in/shared/components/layout'
import { PageTitle } from '@pitch-in/shared/'
import {
	Button,
	InputErrorMessage,
	TextInput,
} from '@pitch-in/shared/components'
import { LoginResponse } from '@pitch-in/shared/types'
import { AuthAPI } from '@/lib/clients/api'
import { DASHBOARD_URL } from '@/lib'
import useAuthStore from '@/stores/useAuthStore'

const LoginPage = () => {
	const router = useRouter()
	const [isLoading, setIsLoading] = useState(false)
	const setLoginStatus = useAuthStore((state) => state.setLoginStatus)
	const {
		register,
		handleSubmit,
		formState: { errors },
		setError,
		clearErrors,
	} = useForm<FieldValues>({
		defaultValues: {
			email: '',
			password: '',
		},
	})

	const login = async (data: FieldValues) => {
		clearErrors()
		setIsLoading(true)
		try {
			const res: LoginResponse = await AuthAPI.login(data.email, data.password)
			setLoginStatus({ token: res.token, refresh: res.refresh })
			useAuthStore.getState().setUser(res.user)
			// need groups/permissions next
			router.push(process.env.NEXT_PUBLIC_AUTH_REDIRECT_URL || DASHBOARD_URL)
		} catch (e: unknown) {
			if (axios.isAxiosError(e)) {
				const axiosError = e as AxiosError<{ detail?: string }>
				if (axiosError.response && axiosError.response.status === 401) {
					setError('serverError', {
						message: axiosError.response.data.detail || 'Login failed',
					})
				}
			} else if (e instanceof Error) {
				setError('serverError', { message: e.message })
				console.error(e.message)
			} else {
				setError('serverError', { message: 'An unknown error occurred' })
			}
		} finally {
			setIsLoading(false)
		}
	}

	return (
		<div className='flex min-h-screen items-center justify-center bg-gray-100'>
			<CardContainer className='w-full max-w-md'>
				<PageTitle>Login</PageTitle>
				<form onSubmit={handleSubmit(login)}>
					<TextInput
						id='email'
						label='Email'
						type='email'
						placeholder='Enter your email'
						{...register('email', { required: 'Email is required' })}
					/>
					{errors.email && (
						<InputErrorMessage errorMessage={errors.email.message as string} />
					)}
					<TextInput
						id='password'
						label='Password'
						type='password'
						placeholder='Enter your password'
						{...register('password', { required: 'Password is required' })}
					/>
					{errors.password && (
						<InputErrorMessage
							errorMessage={errors.password.message as string}
						/>
					)}

					{errors.serverError && (
						<InputErrorMessage
							errorMessage={errors.serverError.message as string}
						/>
					)}
					<div className='flex items-center justify-between'>
						<Button type='submit' actionType='submit' isLoading={isLoading}>
							{isLoading ? 'Logging in...' : 'Login'}
						</Button>
					</div>
				</form>
			</CardContainer>
		</div>
	)
}

export default LoginPage
