import { PlusIcon } from '@heroicons/react/20/solid'
import { Button } from '@pitch-in/shared/components'
import { OptionType } from '@pitch-in/shared/types'

interface AvailableListProps {
	items: OptionType[]
	onClick: (itemId: string) => void
}

const AvailableList = ({ items, onClick }: AvailableListProps) => {
	console.log('items: ')
	console.log(items)
	return (
		<div className='max-h-full space-y-2 overflow-y-auto'>
			{items.map((i) => (
				<div
					key={i.value}
					className='flex items-center justify-between rounded-lg border-gray-200 bg-white p-3 shadow-sm transition duration-150 hover:bg-indigo-50'
				>
					<div className='min-w-0 pr-4'>
						<p className='truncate text-sm font-medium text-gray-900'>
							{i.label}
						</p>
					</div>
					<Button
						actionType='submit'
						type='button'
						onClick={() => onClick(i.value as string)}
					>
						<PlusIcon className='h-4 w-4' aria-hidden='true' />
					</Button>
				</div>
			))}
		</div>
	)
}

interface DualListSelectorProps {
	addTitle: string
	removeTitle: string
	addItems: OptionType[]
	removeItems: OptionType[]
	onAdd: (itemId: string) => void
	onRemove: (itemId: string) => void
}
export const DualListSelector = ({
	addTitle,
	removeTitle,
	addItems,
	removeItems,
	onAdd,
	onRemove,
}: DualListSelectorProps) => {
	return (
		<div className='grid h-full grid-cols-2 gap-8'>
			<div className='flex h-full flex-col overflow-y-auto'>
				<h3 className='mb-4 text-lg font-semibold text-gray-900'>{addTitle}</h3>
				<AvailableList items={addItems} onClick={onAdd} />
			</div>
			<div className='flex h-full flex-col overflow-y-auto'>
				<h3 className='mb-4 text-lg font-semibold text-gray-900'>
					{removeTitle}
				</h3>
				<AvailableList items={removeItems} onClick={onRemove} />
			</div>
		</div>
	)
}
