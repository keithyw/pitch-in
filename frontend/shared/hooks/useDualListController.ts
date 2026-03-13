'use client'

import { useCallback, useEffect, useMemo, useRef, useState } from 'react'
import toast from 'react-hot-toast'

interface UseDualListControllerProps<T, U> {
	isLoading: boolean
	item: U
	addItems: T[]
	removeItems: T[]
	getLabel: (item: T) => string
	attachFn: (attachId: number, toId: number) => Promise<void>
	detachFn: (detachId: number, fromId: number) => Promise<void>
	onUpdate?: (items: T[]) => void
}

// sets up the functionality for dealing with
// m2m style relations where you need to add/remove items from
// one list to another using the DualListSelector component
// and manages the add/remove items state
export function useDualListController<
	T extends { id: number },
	U extends { id: number },
>({
	isLoading,
	item,
	addItems,
	removeItems,
	getLabel,
	attachFn,
	detachFn,
	onUpdate,
}: UseDualListControllerProps<T, U>) {
	const [isSaving, setIsSaving] = useState<boolean>(false)
	const [localAddItems, setLocalAddItems] = useState<T[]>([])
	const [localRemoveItems, setLocalRemoveItems] = useState<T[]>([])

	// Track the last IDs we initialized with so we only re-init when
	// the actual data changes, not just when array references change.
	const initializedAddIds = useRef<string>('')
	const initializedRemoveIds = useRef<string>('')

	useEffect(() => {
		if (isLoading || isSaving) return

		const addIds = addItems.map((i) => i.id).join(',')
		const removeIds = removeItems.map((i) => i.id).join(',')

		// Skip if the data hasn't actually changed
		if (
			addIds === initializedAddIds.current &&
			removeIds === initializedRemoveIds.current
		) {
			return
		}

		initializedAddIds.current = addIds
		initializedRemoveIds.current = removeIds

		const attachedItemIds = new Set(removeItems.map((i) => i.id))
		const availableItems = addItems.filter((i) => !attachedItemIds.has(i.id))
		setLocalAddItems(availableItems)
		setLocalRemoveItems(removeItems)
	}, [addItems, isLoading, removeItems, isSaving])

	const addOptions = useMemo(() => {
		return localAddItems.map((m) => ({
			value: m.id,
			label: getLabel(m),
		}))
	}, [localAddItems, getLabel])

	const removeOptions = useMemo(() => {
		return localRemoveItems.map((m) => ({
			value: m.id,
			label: getLabel(m),
		}))
	}, [localRemoveItems, getLabel])

	const handleAddItem = useCallback(
		async (id: number) => {
			if (!item?.id) return
			setIsSaving(true)
			try {
				await attachFn(id, item.id)
				const addItem = localAddItems.find((m) => m.id === id)
				if (!addItem) return
				const updatedRemoveItems = [...localRemoveItems, addItem]

				setLocalAddItems((prev) => prev.filter((m) => m.id !== id))
				setLocalRemoveItems(updatedRemoveItems)
				if (onUpdate) onUpdate(updatedRemoveItems)
				toast.success('Success Attaching Item')
			} catch (e: unknown) {
				console.error('Failed attaching item: ', e)
				toast.error('Attaching item failed')
			} finally {
				setIsSaving(false)
			}
			// eslint-disable-next-line react-hooks/exhaustive-deps
		},
		[attachFn, item?.id, localAddItems, localRemoveItems, onUpdate],
	)

	const handleRemoveItem = useCallback(
		async (id: number) => {
			if (!item?.id) return
			setIsSaving(true)
			try {
				await detachFn(id, item.id)
				const updatedRemoveItems = localRemoveItems.filter((m) => m.id !== id)
				setLocalRemoveItems(updatedRemoveItems)
				setLocalAddItems((prev) => {
					const removeItem = localRemoveItems.find((m) => m.id === id)
					return removeItem ? [...prev, removeItem] : prev
				})
				if (onUpdate) onUpdate(updatedRemoveItems)
				toast.success('Success Detaching Item')
			} catch (e: unknown) {
				console.error('Failed detaching item: ', e)
				toast.error('Detaching item failed')
			} finally {
				setIsSaving(false)
			}
			// eslint-disable-next-line react-hooks/exhaustive-deps
		},
		[detachFn, item?.id, localRemoveItems],
	)

	return {
		addOptions,
		removeOptions,
		handleAddItem,
		handleRemoveItem,
	}
}
