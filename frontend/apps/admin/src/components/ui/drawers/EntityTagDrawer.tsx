'use client'

import { useEffect, useState } from 'react'
import {
	AutocompleteInput,
	ChipContainer,
	SlideOverDrawer,
} from '@pitch-in/shared/components'
import { OptionType, Tag } from '@pitch-in/shared/types'
import { TagAPI } from '@/lib/clients/api'

interface Entity {
	id: number
	name: string
}

interface EntityTagDrawerProps<T extends Entity> {
	isLoading: boolean
	isOpen: boolean
	item: T
	itemTags: Tag[]
	entityType: 'items' // will expand later
	onClose: () => void
}

const EntityTagDrawer = <T extends Entity>({
	isLoading,
	isOpen,
	onClose,
	item,
	itemTags,
	entityType,
}: EntityTagDrawerProps<T>) => {
	const [addedItemTags, setAddedItemTags] = useState<Tag[]>([])
	const [searchTags, setSearchTags] = useState<Tag[]>([])
	const [resetKey, setResetkey] = useState(0)

	useEffect(() => {
		if (!item || isLoading) return
		// local copy
		setAddedItemTags(itemTags)
	}, [])

	const handleSearch = async (query: string): Promise<OptionType[]> => {
		try {
			const res = await TagAPI.fetch({
				page: 1,
				pageSize: 10,
				fields: [{ field: 'tag', value: query, operator: '~=' }],
			})
			if (res.count > 0) {
				setSearchTags(res.results)
				return res.results.map((t) => ({ value: t.id, label: t.tag }))
			}
		} catch (e: unknown) {
			console.error(e)
		}
		return []
	}

	const handleSelect = async (
		val: string | number | null,
		textValue?: string,
	) => {
		let addTag: Tag | undefined
		try {
			if (val) {
				await TagAPI.attach(val as number, item.id, entityType)
				addTag = searchTags.find((t) => t.id === val)
			} else if (textValue) {
				const newTag = await TagAPI.create({ tag: textValue })
				await TagAPI.attach(newTag.id, item.id, entityType)
				addTag = newTag
			}

			if (addTag) {
				setAddedItemTags((prev) => [...prev, addTag!])
				setResetkey((prev) => prev + 1)
			}
		} catch (e: unknown) {
			console.error(e)
		}
	}

	const handleRemove = async (t: Tag) => {
		try {
			await TagAPI.detach(t.id, item.id, entityType)
			setAddedItemTags((prev) =>
				prev.filter((existingTag) => existingTag.id !== t.id),
			)
		} catch (e: unknown) {
			console.error(e)
		}
	}

	if (!item || isLoading) return null

	return (
		<SlideOverDrawer
			title={`Manage Tags for "${item.name}"`}
			isOpen={isOpen}
			onClose={onClose}
			panelWidthClass='max-w-2xl'
		>
			<AutocompleteInput
				id='tag'
				label='Add Tags to Item. Hit return to add.'
				onSearch={(q) => handleSearch(q)}
				onSelect={(v, text) => handleSelect(v, text)}
				resetTriggger={resetKey}
			/>
			<ChipContainer
				itemName='tags'
				fieldName='tag'
				isLoadingData={isLoading}
				data={addedItemTags}
				errors=''
				onRemove={(t) => handleRemove(t)}
			/>
		</SlideOverDrawer>
	)
}

export default EntityTagDrawer
