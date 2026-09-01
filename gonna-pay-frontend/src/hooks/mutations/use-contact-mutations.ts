'use client'

import { useMutation, useQueryClient } from '@tanstack/react-query'
import { ContactService } from '@/services/contact.service'
import { toast } from 'sonner'
import type { Contact, CreateContactInput } from '@/types'
import { ApiError } from '@/lib/errors/api.error'

export function useCreateContact() {
  const queryClient = useQueryClient()

  return useMutation<Contact, ApiError, CreateContactInput>({
    mutationFn: (data: CreateContactInput) => ContactService.create(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['contacts'] })
      toast.success('Contact created successfully!')
    },
    onError: (error: ApiError) => {
      if (error.status !== 409) {
    toast.error(error.message)
    }
    },
  })
}

export function useUpdateContact() {
  const queryClient = useQueryClient()

  return useMutation<Contact, ApiError, { id: string; data: Partial<CreateContactInput> }>({
    mutationFn: ({ id, data }) => ContactService.update(id, data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['contacts'] })
      toast.success('Contact updated successfully!')
    },
    onError: (error: ApiError) => {
      toast.error(error.message)
    },
  })
}

export function useDeleteContact() {
  const queryClient = useQueryClient()

  return useMutation<void, ApiError, string>({
    mutationFn: (id: string) => ContactService.delete(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['contacts'] })
      toast.success('Contact deleted successfully!')
    },
    onError: (error: ApiError) => {
      toast.error(error.message)
    },
  })
}
