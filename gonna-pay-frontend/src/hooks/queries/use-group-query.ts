'use client'

import { useQueries, useQuery } from '@tanstack/react-query'
import { GroupService } from '@/services/group.service'
import type { CostDetail, GetGroupsResponse, GroupDetail } from '@/types'
import { ApiError } from '@/lib/errors/api.error'
import { useCostsQuery } from './use-cost-query'
import { CostService } from '@/services'

export function useGroupsQuery() {
  return useQuery<GetGroupsResponse, ApiError>({
    queryKey: ['groups'],
    queryFn: () => GroupService.getAll(),
    staleTime: 1000 * 60 * 5,
  })
}

export function useGroupsWithMembersQuery() {
  const { data: listData, ...listQuery } = useGroupsQuery()
  const groups = listData?.groups ?? []

  const detailQueries = useQueries({
    queries: groups.map((group) => ({
      queryKey: ['groups', group.id] as const,
      queryFn: () => GroupService.getById(group.id),
      staleTime: 1000 * 60 * 5,
    })),
  })

  const groupsWithMembers: GroupDetail[] = groups.map((group, i) => ({
    ...group,
    members: detailQueries[i]?.data?.members ?? group.members,
  }))

  return {
    ...listQuery,
    data: listData
      ? { groups: groupsWithMembers, total: listData.total }
      : undefined,
  }
}

export function useGroupQuery(id: string) {
  return useQuery<GroupDetail, ApiError>({
    queryKey: ['groups', id],
    queryFn: () => GroupService.getById(id),
    staleTime: 1000 * 60 * 5,
    enabled: !!id,
  })
}

export function useCostsWithSplitsQuery() {
  const { data: costs = [], ...listQuery } = useCostsQuery()

  const detailQueries = useQueries({
    queries: costs.map((cost) => ({
      queryKey: ['costs', cost.id],
      queryFn: () => CostService.getById(cost.id),
      scaleTime: 1000 * 60 * 5,
    })),
  })

  const costsWithSplits = detailQueries
    .map((q) => q.data)
    .filter((d): d is CostDetail => Boolean(d))

  return { ...listQuery, data: costsWithSplits }
}
