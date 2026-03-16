'use client'

import type { Metadata } from 'next'
import { Geist, Geist_Mono } from 'next/font/google'
import { Toaster } from 'react-hot-toast'
import { PermissionProvider } from '@pitch-in/shared/contexts'
import AuthProvider from '@/components/layout/AuthProvider'
import Navbar from '@/components/layout/Navbar'
import useAuthStore from '@/stores/useAuthStore'
import './globals.css'

const geistSans = Geist({
	variable: '--font-geist-sans',
	subsets: ['latin'],
})

const geistMono = Geist_Mono({
	variable: '--font-geist-mono',
	subsets: ['latin'],
})

export default function RootLayout({
	children,
}: Readonly<{
	children: React.ReactNode
}>) {
	const checkAccess = useAuthStore((state) => state.checkAccess)

	return (
		<html lang='en'>
			<body
				className={`${geistSans.variable} ${geistMono.variable} antialiased`}
			>
				<AuthProvider>
					<PermissionProvider checkAccess={checkAccess}>
						<Navbar />
						<div>{children}</div>
						<Toaster position='bottom-right' />
					</PermissionProvider>
				</AuthProvider>
			</body>
		</html>
	)
}
