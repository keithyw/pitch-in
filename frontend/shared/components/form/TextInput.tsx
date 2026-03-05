import { forwardRef } from 'react'

interface TextInputProps extends React.InputHTMLAttributes<HTMLInputElement> {
	id: string
	label: string
	type?: string
	className?: string
	readOnly?: boolean
}

export const TextInput = forwardRef<HTMLInputElement, TextInputProps>(
	({ id, label, type = 'text', className = '', readOnly, ...rest }, ref) => {
		return (
			<div className='relative mb-4'>
				<label
					htmlFor={id}
					className='mb-2 block text-sm font-bold text-gray-700'
				>
					{label}
				</label>
				<input
					type={type}
					id={id}
					ref={ref}
					readOnly={readOnly}
					{...rest}
					className={`focus:shadow-outline w-full appearance-none rounded border px-3 py-2 leading-tight text-gray-700 shadow focus:border-blue-500 focus:outline-none ${readOnly ? 'cursor-not-allowed bg-gray-100' : ''} ${className || ''} `}
				/>
			</div>
		)
	},
)

TextInput.displayName = 'TextInput'
