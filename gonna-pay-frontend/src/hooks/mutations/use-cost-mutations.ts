'use client'

import { useMutation, useQueryClient } from '@tanstack/react-query'
import { CostService } from '@/services/cost.service'
import { toast } from 'sonner'
import type { Cost, CostDetail, CreateCostInput, UpdateCostInput } from '@/types'
import { ApiError } from '@/lib/errors/api.error'

export function useCreateCost() {
  const queryClient = useQueryClient()

  return useMutation<Cost, ApiError, CreateCostInput>({
    mutationFn: (data) => CostService.create(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['costs'] })
      toast.success('Cost created successfully!')
    },
    onError: (error: ApiError) => {
      toast.error(error.message)
    },
  })
}

export function useUpdateCost() {
  const queryClient = useQueryClient()

  return useMutation<CostDetail, ApiError, { id: string; data: UpdateCostInput }>({
    mutationFn: ({ id, data }) => CostService.update(id, data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['costs'] })
      toast.success('Cost updated successfully!')
    },
    onError: (error: ApiError) => {
      toast.error(error.message)
    },
  })
}

export function useDeleteCost() {
  const queryClient = useQueryClient()

  return useMutation<void, ApiError, string>({
    mutationFn: (id) => CostService.delete(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['costs'] })
      toast.success('Cost deleted successfully!')
    },
    onError: (error: ApiError) => {
      toast.error(error.message)
    },
  })
}
