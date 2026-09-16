'use client'

import { useState } from 'react'
import { useGroupsWithMembersQuery } from '@/hooks/queries/use-group-query'
import { TablePagination } from '@/components/table-pagination'
import { GroupCards } from '../group-cards'

export const TableGroup = () => {
  const [page, setPage] = useState(1)
  const { data } = useGroupsWithMembersQuery(page)
  const groups = data?.groups ?? []
  const total = data?.total ?? 0
  const totalPages = data?.totalPages ?? 1

  return (
    <div className="flex flex-col gap-4">
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
