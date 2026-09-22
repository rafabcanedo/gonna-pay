import { useState } from 'react'
import { Search } from 'lucide-react'
import { cn } from '@/lib/utils'
import { Input } from '@/components/ui/input'
import type { IPropsSearchInput } from './types'

export function SearchInput({ value, onChange, placeholder = 'Search...', className }: IPropsSearchInput) {
  const [isOpen, setIsOpen] = useState(false)

  const handleBlur = () => {
    if (value === '') setIsOpen(false)
  }

  const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key === 'Escape') {
      onChange('')
      setIsOpen(false)
    }
  }

  return (
    <div className="flex items-center">
      <div className={cn('transition-all duration-200 overflow-hidden', isOpen ? (className ?? 'w-36') : 'w-0')}>
        <Input
          autoFocus={isOpen}
          value={value}
          onChange={(e) => onChange(e.target.value)}
          onBlur={handleBlur}
          onKeyDown={handleKeyDown}
          placeholder={placeholder}
          className="h-8 text-xs placeholder:text-xs focus-visible:ring-0"
        />
      </div>
      <button
        type="button"
        onClick={() => setIsOpen(true)}
        className="p-1.5 rounded-md hover:bg-zinc-100 text-zinc-400 hover:text-zinc-600 transition-colors"
      >
        <Search className="w-4 h-4" />
      </button>
    </div>
  )
}
