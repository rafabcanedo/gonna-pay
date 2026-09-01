'use client'

import { useGroupsQuery } from '@/hooks/queries/use-group-query'
import { GroupCards } from '../group-cards'

export const TableGroup = () => {
  const { data } = useGroupsQuery()
  const groups = data?.groups ?? []
  const total = data?.total ?? 0

  return (
    <div className="flex flex-col gap-4">
      <span className="font-poppins text-sm text-zinc-400">{total} {total === 1 ? 'group' : 'groups'}</span>
      <div className="flex flex-col gap-4 max-w-xl">
        {groups.map((group) => (
          <GroupCards key={group.id} group={group} />
        ))}
      </div>
    </div>
  )
}
