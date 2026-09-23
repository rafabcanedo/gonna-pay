'use client'

import { useState, useEffect } from 'react'
import { useGroupsWithMembersQuery } from '@/hooks/queries/use-group-query'
import type { GroupQueryFilters } from '@/hooks/queries/types'
import { useDebounce } from '@/hooks/use-debounce'
import { TablePagination } from '@/components/table-pagination'
import { GroupCards } from '../group-cards'
import { GroupsFilter } from './components/groups-filter'

export const TableGroup = () => {
  const [page, setPage] = useState(1)
  const [category, setCategory] = useState('')
  const [search, setSearch] = useState('')
  const debouncedSearch = useDebounce(search)

  const filters: GroupQueryFilters = {
    ...(category ? { category } : {}),
    ...(debouncedSearch ? { search: debouncedSearch } : {}),
  }

  useEffect(() => { setPage(1) }, [category, debouncedSearch])

  const { data } = useGroupsWithMembersQuery(page, 20, filters)
  const groups = data?.data ?? []
  const total = data?.total ?? 0
  const totalPages = data?.totalPages ?? 1

  return (
    <div className="flex flex-col gap-4">
      <div>
        <GroupsFilter
          category={category}
          search={search}
          onCategoryChange={setCategory}
          onSearchChange={setSearch}
        />
      </div>
      <span className="font-poppins text-sm text-zinc-400">{total} {total === 1 ? 'group' : 'groups'}</span>
      <div className="flex flex-col gap-4 max-w-xl">
        {groups.map((group) => (
          <GroupCards key={group.id} group={group} />
        ))}
      </div>

      {totalPages > 1 && (
        <TablePagination page={page} totalPages={totalPages} onPageChange={setPage} />
      )}
    </div>
  )
}
