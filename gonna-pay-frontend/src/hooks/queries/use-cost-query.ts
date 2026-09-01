"use client";

import { useQuery } from "@tanstack/react-query";
import { CostService } from "@/services/cost.service";
import type { GetCostsResponse } from "@/types";
import { ApiError } from "@/lib/errors/api.error";

export function useCostsQuery() {
  return useQuery<GetCostsResponse, ApiError>({
    queryKey: ["costs"],
    queryFn: () => CostService.getAll(),
    staleTime: 1000 * 60 * 5,
  });
}
