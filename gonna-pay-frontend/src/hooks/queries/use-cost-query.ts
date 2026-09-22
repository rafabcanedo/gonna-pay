"use client";

import { useQuery, useQueries } from "@tanstack/react-query";
import { CostService } from "@/services/cost.service";
import type { GetCostsResponse, CostDetail } from "@/types";
import { ApiError } from "@/lib/errors/api.error";
import type { CostQueryFilters } from './types'

export function useCostsQuery(page = 1, limit = 20, filters?: CostQueryFilters) {
  return useQuery<GetCostsResponse, ApiError>({
    queryKey: ["costs", { page, limit, ...filters }],
    queryFn: () => {
      const params = new URLSearchParams({ page: String(page), limit: String(limit) })
      if (filters?.category) params.set('category', filters.category)
      if (filters?.period) params.set('period', filters.period)
      if (filters?.type) params.set('type', filters.type)
      if (filters?.minValue !== undefined) params.set('minValue', String(filters.minValue))
      if (filters?.maxValue !== undefined) params.set('maxValue', String(filters.maxValue))
      return CostService.getAll(params)
    },
    staleTime: 1000 * 60 * 5,
  });
}

export function useCostsWithSplitsQuery() {
  const { data: response, ...listQuery } = useCostsQuery()
  const costs = response?.data ?? []

  const detailQueries = useQueries({
    queries: costs.map((cost) => ({
      queryKey: ["costs", cost.id],
      queryFn: () => CostService.getById(cost.id),
      staleTime: 1000 * 60 * 5,
    })),
  })

  const costsWithSplits = detailQueries
    .map((q) => q.data)
    .filter((d): d is CostDetail => Boolean(d))

  return { ...listQuery, data: costsWithSplits }
}

export function useCostQuery(id: string) {
  return useQuery<CostDetail, ApiError>({
    queryKey: ['costs', id],
    queryFn: () => CostService.getById(id),
    staleTime: 1000 * 60 * 5,
    enabled: !!id,
  })
}
