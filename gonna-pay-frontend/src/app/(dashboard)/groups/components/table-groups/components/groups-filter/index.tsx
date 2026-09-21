import { useState } from 'react'
import { Search } from 'lucide-react'
import { Input } from '@/components/ui/input'
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select'
import { TransactionCategory } from '@/types'
import type { IPropsGroupsFilter } from './types'

const GROUP_CATEGORY_OPTIONS = [
  { label: 'All', value: '' },
  { label: 'Dinner', value: TransactionCategory.DINNER },
  { label: 'Lunch', value: TransactionCategory.LUNCH },
  { label: 'Entertainment', value: TransactionCategory.ENTERTAINMENT },
  { label: 'Travel', value: TransactionCategory.TRAVEL },
  { label: 'Others', value: TransactionCategory.OTHERS },
]

export function GroupsFilter({ category, search, onCategoryChange, onSearchChange }: IPropsGroupsFilter) {
  const [isSearchOpen, setIsSearchOpen] = useState(false)

  const handleBlur = () => {
    if (search === '') setIsSearchOpen(false)
  }

  const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key === 'Escape') {
      onSearchChange('')
      setIsSearchOpen(false)
    }
  }

  return (
    <div className="flex items-center gap-2">
      <div className="flex items-center">
        <div className={`transition-all duration-200 overflow-hidden ${isSearchOpen ? 'w-36' : 'w-0'}`}>
          <Input
            autoFocus={isSearchOpen}
            value={search}
            onChange={(e) => onSearchChange(e.target.value)}
            onBlur={handleBlur}
            onKeyDown={handleKeyDown}
            placeholder="Search by name..."
            className="h-8 text-xs"
          />
        </div>
        <button
          type="button"
          onClick={() => setIsSearchOpen(true)}
          className="p-1.5 rounded-md hover:bg-zinc-100 text-zinc-400 hover:text-zinc-600 transition-colors"
        >
          <Search className="w-4 h-4" />
        </button>
      </div>

      <Select value={category} onValueChange={onCategoryChange}>
        <SelectTrigger className="h-8 w-32 text-xs">
          <SelectValue placeholder="Category" />
        </SelectTrigger>
        <SelectContent>
          {GROUP_CATEGORY_OPTIONS.map((opt) => (
            <SelectItem key={opt.value} value={opt.value}>
              {opt.label}
            </SelectItem>
          ))}
        </SelectContent>
      </Select>
    </div>
  )
}
