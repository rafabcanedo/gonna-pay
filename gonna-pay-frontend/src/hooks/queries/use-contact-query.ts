'use client'

import { useQuery } from '@tanstack/react-query'
import { ContactService } from '@/services/contact.service'
import type { Contact, GetContactsResponse } from '@/types'
import { ApiError } from '@/lib/errors/api.error'

export function useContactsQuery() {
  return useQuery<GetContactsResponse, ApiError>({
    queryKey: ['contacts'],
    queryFn: () => ContactService.getAll(),
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
