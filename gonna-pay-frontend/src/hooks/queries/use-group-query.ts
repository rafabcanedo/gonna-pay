'use client'

import { useQueries, useQuery } from '@tanstack/react-query'
import { GroupService } from '@/services/group.service'
import type { GetGroupsResponse, GroupDetail } from '@/types'
import { ApiError } from '@/lib/errors/api.error'
import { QUERY_STALE_TIME } from '@/lib/constants'
import type { GroupQueryFilters } from './types'

export function useGroupsQuery(page = 1, limit = 20, filters?: GroupQueryFilters) {
  return useQuery<GetGroupsResponse, ApiError>({
    queryKey: ['groups', { page, limit, ...filters }],
    queryFn: () => {
      const params = new URLSearchParams({ page: String(page), limit: String(limit) })
      if (filters?.category) params.set('category', filters.category)
      if (filters?.search) params.set('search', filters.search)
      return GroupService.getAll(params)
    },
    staleTime: QUERY_STALE_TIME,
  })
}

export function useGroupsWithMembersQuery(page = 1, limit = 20, filters?: GroupQueryFilters) {
  const { data: listData, ...listQuery } = useGroupsQuery(page, limit, filters)
  const groups = listData?.data ?? []

  const detailQueries = useQueries({
    queries: groups.map((group) => ({
      queryKey: ['groups', group.id] as const,
      queryFn: () => GroupService.getById(group.id),
      staleTime: QUERY_STALE_TIME,
    })),
  })

  const groupsWithMembers: GroupDetail[] = groups.map((group, i) => ({
    ...group,
    members: detailQueries[i]?.data?.members ?? group.members,
  }))

  return {
    ...listQuery,
    data: listData
      ? {
          data: groupsWithMembers,
          page: listData.page,
          limit: listData.limit,
          total: listData.total,
          totalPages: listData.totalPages,
        }
      : undefined,
  }
}

export function useGroupQuery(id: string) {
  return useQuery<GroupDetail, ApiError>({
    queryKey: ['groups', id],
    queryFn: () => GroupService.getById(id),
    staleTime: QUERY_STALE_TIME,
    enabled: !!id,
  })
}
