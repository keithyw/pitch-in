import {
	DocumentIcon,
	PhotoIcon,
	VideoCameraIcon,
	MusicalNoteIcon,
	DocumentTextIcon,
	ArchiveBoxIcon,
	QuestionMarkCircleIcon,
} from '@heroicons/react/24/outline'

export interface FileMetadata {
	fileName: string
	mimeType: string
	width?: number
	height?: number
	previewUrl?: string
}

/**
 * Maps a MIME type string to a specific Heroicon component.
 */
export const getFileIcon = (mimeType: string | undefined | null) => {
	if (!mimeType) return QuestionMarkCircleIcon

	// Image types
	if (mimeType.startsWith('image/')) return PhotoIcon

	// Video types
	if (mimeType.startsWith('video/')) return VideoCameraIcon

	// Audio types
	if (mimeType.startsWith('audio/')) return MusicalNoteIcon

	// Specific Document types
	switch (mimeType) {
		case 'application/pdf':
			return DocumentIcon
		case 'application/msword':
		case 'application/vnd.openxmlformats-officedocument.wordprocessingml.document':
			return DocumentTextIcon
		case 'application/zip':
		case 'application/x-7z-compressed':
		case 'application/x-rar-compressed':
			return ArchiveBoxIcon
		default:
			return DocumentIcon // Default fallback for generic files
	}
}

/**
 * Processes a file to extract metadata and generate a preview if applicable.
 */
export const getFileMetadata = async (file: File): Promise<FileMetadata> => {
	const metadata: FileMetadata = {
		fileName: file.name,
		mimeType: file.type,
	}

	// Only perform dimension checks for images
	if (file.type.startsWith('image/')) {
		return new Promise((resolve) => {
			const objectUrl = URL.createObjectURL(file)
			const img = new Image()

			img.onload = () => {
				metadata.width = img.naturalWidth
				metadata.height = img.naturalHeight
				metadata.previewUrl = objectUrl // Keep this to display in UI
				resolve(metadata)
			}

			img.onerror = () => {
				URL.revokeObjectURL(objectUrl)
				resolve(metadata) // Resolve with basic name/mime if image load fails
			}

			img.src = objectUrl
		})
	}

	// For non-images, just return the name and mime type
	return metadata
}
