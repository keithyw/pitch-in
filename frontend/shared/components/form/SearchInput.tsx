interface SearchInputProps {
	value: string
	onChange: (term: string) => void
	placeholder?: string
	className?: string
	background?: string
}

export const SearchInput = ({
	value,
	onChange,
	placeholder = 'Search…',
	className = '',
	background = 'dark:bg-gray-700 dark:border-gray-600 dark:text-white dark:focus:ring-blue-400',
}: SearchInputProps) => {
	const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
		onChange(e.target.value)
	}

	return (
		<div className='relative'>
			<input
				type='text'
				className={`w-full rounded-md border border-gray-300 py-2 pr-4 pl-10 shadow-sm focus:border-blue-500 focus:ring-2 focus:ring-blue-500 focus:outline-none ${background} ${className}`}
				placeholder={placeholder}
				value={value}
				onChange={handleChange}
			/>
			<svg
				className='absolute top-1/2 left-3 -translate-y-1/2 text-gray-400'
				xmlns='http://www.w3.org/2000/svg'
				width='18'
				height='18'
				viewBox='0 0 24 24'
				fill='none'
				stroke='currentColor'
				strokeWidth='2'
				strokeLinecap='round'
				strokeLinejoin='round'
			>
				<circle cx='11' cy='11' r='8'></circle>
				<line x1='21' y1='21' x2='16.65' y2='16.65'></line>
			</svg>
		</div>
	)
}

export default SearchInput
