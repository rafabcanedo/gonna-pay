'use client'

import { useQueries, useQuery } from '@tanstack/react-query'
import { GroupService } from '@/services/group.service'
import type { GetGroupsResponse, GroupDetail } from '@/types'
import { ApiError } from '@/lib/errors/api.error'

export type GroupQueryFilters = {
  category?: string
  search?: string
}

export function useGroupsQuery(page = 1, limit = 20, filters?: GroupQueryFilters) {
  return useQuery<GetGroupsResponse, ApiError>({
    queryKey: ['groups', { page, limit, ...filters }],
    queryFn: () => {
      const params = new URLSearchParams({ page: String(page), limit: String(limit) })
      if (filters?.category) params.set('category', filters.category)
      if (filters?.search) params.set('search', filters.search)
      return GroupService.getAll(params)
    },
    staleTime: 1000 * 60 * 5,
  })
}

export function useGroupsWithMembersQuery(page = 1, limit = 20, filters?: GroupQueryFilters) {
  const { data: listData, ...listQuery } = useGroupsQuery(page, limit, filters)
  const groups = listData?.data ?? []

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
      ? { groups: groupsWithMembers, total: listData.total, totalPages: listData.totalPages }
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
