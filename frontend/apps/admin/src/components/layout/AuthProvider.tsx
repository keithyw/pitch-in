'use client'

import { useEffect, ReactNode } from 'react'
import { useRouter, usePathname } from 'next/navigation'
import { LoadingSpinner } from '@pitch-in/shared/components'
import {
	DASHBOARD_URL,
	LOGIN_URL,
	PROTECTED_ROUTES,
	PUBLIC_ROUTES,
} from '@/lib/constants'
import useAuthStore from '@/stores/useAuthStore'

interface AuthProviderProps {
	children: ReactNode
}

const AuthProvider = ({ children }: AuthProviderProps) => {
	const router = useRouter()
	const pathname = usePathname()

	const { isAuthenticated, isLoading, checkIfAuthenticated } = useAuthStore()

	useEffect(() => {
		checkIfAuthenticated()
	}, [checkIfAuthenticated])

	useEffect(() => {
		if (isLoading) return
		const isProtected = PROTECTED_ROUTES.some((p) => pathname.startsWith(p))

		const isAuthPage = PUBLIC_ROUTES.includes(pathname)

		if (isProtected && !isAuthenticated) {
			router.replace(LOGIN_URL)
		} else if (isAuthPage && isAuthenticated) {
			router.replace(process.env.NEXT_PUBLIC_AUTH_REDIRECT_URL || DASHBOARD_URL)
		}
	}, [isAuthenticated, pathname, router])

	const isProtected = PROTECTED_ROUTES.some((p) => pathname.startsWith(p))

	if (isLoading && isProtected) {
		return <LoadingSpinner />
	}

	return <>{children}</>
}

export default AuthProvider
