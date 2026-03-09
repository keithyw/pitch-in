'use client'

import { useCallback, useState } from 'react'
import toast from 'react-hot-toast'

interface UseDeleteRecordProps<T> {
	deleteFn: (id: number) => Promise<void>
	onSuccess?: () => void
	itemNameProp: keyof T
}

export function useDeleteRecord<T extends { id: number }>({
	deleteFn,
	onSuccess,
	itemNameProp,
}: UseDeleteRecordProps<T>) {
	const [isModalOpen, setIsModalOpen] = useState<boolean>(false)
	const [isDeleting, setIsDeleting] = useState<boolean>(false)
	const [deleteItem, setDeleteItem] = useState<T | null>(null)

	const openDeleteModal = useCallback((item: T) => {
		setDeleteItem(item)
		setIsModalOpen(true)
	}, [])

	const closeDeleteModal = useCallback(() => {
		setIsModalOpen(false)
		setDeleteItem(null)
	}, [])

	const handleDelete = async () => {
		if (!deleteItem) return
		setIsDeleting(true)
		try {
			await deleteFn(deleteItem?.id)
			toast.success(`${String(deleteItem[itemNameProp])} has been deleted`)
			if (onSuccess) onSuccess()
			closeDeleteModal()
		} catch (e: unknown) {
			console.error('Failed deleting item: ', e)
			toast.error('Failed to deletee item')
		} finally {
			setIsDeleting(false)
		}
	}

	return {
		isModalOpen,
		deleteItem,
		isDeleting,
		openDeleteModal,
		closeDeleteModal,
		handleDelete,
	}
}
