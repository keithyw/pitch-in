import type { Config } from 'tailwindcss'

const config: Config = {
	content: [
		'./src/**/*.{js,ts,jsx,tsx,mdx}',
		// This is the critical line for monorepos:
		'../../shared/components/**/*.{js,ts,jsx,tsx,mdx}',
	],
	theme: {
		extend: {},
	},
	plugins: [],
} satisfies Config

export default config
