import { SearchInput } from '@/components/search-input'
import { FilterSelect } from '@/components/filter-select'
import { CONTACT_CATEGORY_OPTIONS } from '@/app/(dashboard)/my-contacts/constants'
import type { IPropsContactsFilter } from './types'

export function ContactsFilter({ category, search, onCategoryChange, onSearchChange }: IPropsContactsFilter) {
  return (
    <div className="flex items-center gap-2">
      <SearchInput value={search} onChange={onSearchChange} placeholder="Search by name..." />
      <FilterSelect
        value={category}
        onChange={onCategoryChange}
        options={CONTACT_CATEGORY_OPTIONS}
        placeholder="Category"
        className="w-32"
      />
    </div>
  )
}
