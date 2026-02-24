import { Toaster } from 'react-hot-toast'
import { Container } from '@pitch-in/shared'
import './globals.css'

export default function RootLayout({
	children,
}: Readonly<{
	children: React.ReactNode
}>) {
	return (
		<html lang='en'>
			<body className='antialiased'>
				<Container as='main'>{children}</Container>
				<Toaster position='bottom-right' />
			</body>
		</html>
	)
}
