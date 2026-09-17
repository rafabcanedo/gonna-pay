'use client'

import { useCostsQuery } from '@/hooks/queries/use-cost-query'
import { CostsSummary } from '../costs-summary'
import { CategoryBreakdown } from '../category-breakdown'

export const CostsStats = () => {
  const { data } = useCostsQuery(1)
  const stats = data?.stats

  if (!stats) return null

  return (
    <>
      <div className="flex flex-row gap-4 mt-8">
        <CostsSummary thisMonth={stats.thisMonth} inSplits={stats.inSplits} solo={stats.solo} />
      </div>

      {stats.byCategory.length > 0 && (
        <div className="mt-4">
          <CategoryBreakdown data={stats.byCategory} />
        </div>
      )}
    </>
  )
}
