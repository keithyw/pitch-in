interface FileUploadInputProps {
	label: string
	onChange: (file: File | null) => void
	currentFile: File | null
}

export const FileUploadInput = ({
	label,
	onChange,
	currentFile,
}: FileUploadInputProps) => {
	const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
		const file = e.target.files ? e.target.files[0] : null
		onChange(file)
	}
	return (
		<div className='flex flex-col space-y-2'>
			<label className='text- sm font-medium text-gray-700'>{label}</label>
			<input
				type='file'
				onChange={handleFileChange}
				className='block w-full text-sm text-gray-500 file:mr-4 file:rounded-full file:border-0 file:bg-indigo-50 file:px-4 file:py-2 file:text-sm file:font-semibold file:text-indigo-700 hover:file:bg-indigo-100'
			/>
			{currentFile && (
				<p className='mt-1 text-xs text-gray-500'>
					Selected: {currentFile.name}
				</p>
			)}
		</div>
	)
}
