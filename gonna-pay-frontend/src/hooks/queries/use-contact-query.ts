'use client'

import { useQuery } from '@tanstack/react-query'
import { ContactService } from '@/services/contact.service'
import type { Contact, GetContactsResponse, GetContactFrequencyResponse } from '@/types'
import { ApiError } from '@/lib/errors/api.error'
import { QUERY_STALE_TIME } from '@/lib/constants'
import type { ContactQueryFilters } from './types'

export function useContactsQuery(page = 1, limit = 20, filters?: ContactQueryFilters) {
  return useQuery<GetContactsResponse, ApiError>({
    queryKey: ['contacts', { page, limit, ...filters }],
    queryFn: () => {
      const params = new URLSearchParams({ page: String(page), limit: String(limit) })
      if (filters?.category) params.set('category', filters.category)
      if (filters?.search) params.set('search', filters.search)
      return ContactService.getAll(params)
    },
    staleTime: QUERY_STALE_TIME,
  })
}

export function useContactQuery(id: string) {
  return useQuery<Contact, ApiError>({
    queryKey: ['contacts', id],
    queryFn: () => ContactService.getById(id),
    staleTime: QUERY_STALE_TIME,
    enabled: !!id,
  })
}

export function useContactFrequencyQuery(limit = 5) {
  return useQuery<GetContactFrequencyResponse, ApiError>({
    queryKey: ['contacts', 'frequency', limit],
    queryFn: () => ContactService.getFrequency(limit),
    staleTime: QUERY_STALE_TIME,
  })
}
