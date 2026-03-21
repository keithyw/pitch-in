'use client'

import { useState } from 'react'
import { File, FileText, ImageIcon, Music, Video } from 'lucide-react'
import { cn } from '@pitch-in/shared/utils'
interface AssetPreviewProps {
	url: string | null | undefined
	alt?: string | null | undefined
	type: string | null | undefined
	className?: string
	size?: 'sm' | 'lg' | 'full'
}

export const AssetPreview = ({
	url,
	alt,
	type,
	className,
	size = 'sm',
}: AssetPreviewProps) => {
	const [hasError, setHasError] = useState(false)

	const sizeClasses = {
		sm: 'h-12 w-12',
		lg: 'h-64 w-64',
		full: 'w-full max-h-[600px]',
	}
	// const commonContainerClasses = `flex h-12 w-12 items-center justify-center rounded-md bg-gray-100 dark:bg-gray-700 ${
	// 	className || ''
	// }`

	const commonContainerClasses = cn(
		'flex items-center justify-center rounded-md bg-gray-100 dark:bg-gray-700',
		sizeClasses[size],
		className,
	)

	const iconSize = size === 'sm' ? 'h-6 w-6' : 'h-16 w-16'
	const commonIconClasses = cn('text-gray-500', iconSize)

	const getIconForType = (assetType: string) => {
		switch (assetType.toLowerCase()) {
			case 'video':
				return <Video className={commonIconClasses} />
			case 'audio':
				return <Music className={commonIconClasses} />
			case 'pdf':
				return <FileText className={commonIconClasses} />
			case 'document':
				return <File className={commonIconClasses} />
			default:
				return <ImageIcon className={commonIconClasses} />
		}
	}

	if (type?.toLowerCase() !== 'image' || !url || hasError) {
		return (
			<div className={commonContainerClasses}>
				{getIconForType(type || 'unknown')}
			</div>
		)
	}

	return (
		<a href={url} target='_blank' rel='noopener noreferrer'>
			<img
				src={url}
				alt={alt || 'Asset preview'}
				className={cn(
					'h-12 w-12 rounded-md object-cover transition-transform duration-200 ease-in-out hover:scale-150',
					sizeClasses[size],
					className,
				)}
				onError={() => setHasError(true)}
				loading='lazy'
			/>
		</a>
	)
}

export default AssetPreview
