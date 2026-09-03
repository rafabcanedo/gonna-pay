"use client";

import { useQuery, useQueries } from "@tanstack/react-query";
import { CostService } from "@/services/cost.service";
import type { GetCostsResponse, CostDetail } from "@/types";
import { ApiError } from "@/lib/errors/api.error";

export function useCostsQuery() {
  return useQuery<GetCostsResponse, ApiError>({
    queryKey: ["costs"],
    queryFn: () => CostService.getAll(),
    staleTime: 1000 * 60 * 5,
  });
}

export function useCostsWithSplitsQuery() {
  const { data: costs = [], ...listQuery } = useCostsQuery()

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
