import { SearchInput } from '@/components/search-input'
import { FilterSelect } from '@/components/filter-select'
import { GROUP_CATEGORY_OPTIONS } from '@/app/(dashboard)/groups/constants'
import type { IPropsGroupsFilter } from './types'

export function GroupsFilter({ category, search, onCategoryChange, onSearchChange }: IPropsGroupsFilter) {
  return (
    <div className="flex items-center gap-2">
      <SearchInput value={search} onChange={onSearchChange} placeholder="Search by name..." />
      <FilterSelect
        value={category}
        onChange={onCategoryChange}
        options={GROUP_CATEGORY_OPTIONS}
        placeholder="Category"
        className="w-32"
      />
    </div>
  )
}
