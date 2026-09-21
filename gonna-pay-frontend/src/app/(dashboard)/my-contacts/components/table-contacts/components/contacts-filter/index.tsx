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
import { CONTACT_CATEGORIES } from '@/app/(dashboard)/my-contacts/components/category-cards/constants'
import type { IPropsContactsFilter } from './types'

const CATEGORY_OPTIONS = [{ label: 'All', value: '' }, ...CONTACT_CATEGORIES]

export function ContactsFilter({ category, search, onCategoryChange, onSearchChange }: IPropsContactsFilter) {
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
          {CATEGORY_OPTIONS.map((opt) => (
            <SelectItem key={opt.value} value={opt.value}>
              {opt.label}
            </SelectItem>
          ))}
        </SelectContent>
      </Select>
    </div>
  )
}
