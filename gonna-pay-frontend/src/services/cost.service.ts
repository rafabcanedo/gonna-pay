import { apiCall } from '@/lib/api-client'
import type { GetCostsResponse, Cost, CostDetail, CreateCostInput, UpdateCostInput } from '@/types'

export const CostService = {
  getAll: async (params?: URLSearchParams) => {
    const query = params ? `?${params.toString()}` : ''
    return apiCall<GetCostsResponse>(`/costs${query}`)
  },

  getById: async (id: string) => {
    return apiCall<CostDetail>(`/costs/${id}`)
  },

  create: async (data: CreateCostInput) => {
    return apiCall<Cost>('/costs', {
      method: 'POST',
      body: JSON.stringify(data),
    })
  },

  update: async (id: string, data: UpdateCostInput) => {
    return apiCall<CostDetail>(`/costs/${id}`, {
      method: 'PATCH',
      body: JSON.stringify(data),
    })
  },

  delete: async (id: string) => {
    return apiCall<void>(`/costs/${id}`, {
      method: 'DELETE',
    })
  },
}
