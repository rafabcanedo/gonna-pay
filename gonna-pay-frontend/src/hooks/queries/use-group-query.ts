'use client'

import { useQuery } from '@tanstack/react-query'
import { GroupService } from '@/services/group.service'
import type { GetGroupsResponse, GroupDetail } from '@/types'
import { ApiError } from '@/lib/errors/api.error'

export function useGroupsQuery() {
  return useQuery<GetGroupsResponse, ApiError>({
    queryKey: ['groups'],
    queryFn: () => GroupService.getAll(),
    staleTime: 1000 * 60 * 5,
  })
}

export function useGroupQuery(id: string) {
  return useQuery<GroupDetail, ApiError>({
    queryKey: ['groups', id],
    queryFn: () => GroupService.getById(id),
    staleTime: 1000 * 60 * 5,
    enabled: !!id,
  })
}
