interface ImageDimensions {
	width: number
	height: number
}

export const getImageDimensionsFromFile = (
	file: File,
): Promise<ImageDimensions> => {
	return new Promise((resolve, reject) => {
		const objectUrl: string = URL.createObjectURL(file)
		const img: HTMLImageElement = new Image()

		img.onload = (): void => {
			const dim = {
				width: img.naturalWidth,
				height: img.naturalHeight,
			}
			URL.revokeObjectURL(objectUrl)
			resolve(dim)
		}

		img.onerror = (): void => {
			URL.revokeObjectURL(objectUrl)
			reject(new Error('Failed to load image for dimension detection'))
		}

		img.src = objectUrl
	})
}

export const isImage = (mimeType: string): boolean => {
	return mimeType.startsWith('image/')
}
