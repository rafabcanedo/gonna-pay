'use client'

import { useQuery } from '@tanstack/react-query'
import { ContactService } from '@/services/contact.service'
import type { Contact, GetContactsResponse, GetContactFrequencyResponse } from '@/types'
import { ApiError } from '@/lib/errors/api.error'

export function useContactsQuery(page = 1, limit = 20) {
  return useQuery<GetContactsResponse, ApiError>({
    queryKey: ['contacts', { page, limit }],
    queryFn: () => ContactService.getAll(new URLSearchParams({ page: String(page), limit: String(limit) })),
    staleTime: 1000 * 60 * 5,
  })
}

export function useContactQuery(id: string) {
  return useQuery<Contact, ApiError>({
    queryKey: ['contacts', id],
    queryFn: () => ContactService.getById(id),
    staleTime: 1000 * 60 * 5,
    enabled: !!id,
  })
}

export function useContactFrequencyQuery(limit = 5) {
  return useQuery<GetContactFrequencyResponse, ApiError>({
    queryKey: ['contacts', 'frequency', limit],
    queryFn: () => ContactService.getFrequency(limit),
    staleTime: 1000 * 60 * 5,
  })
}
