import type { Cost } from '@/types'
import type { CategoryBreakdownPoint } from '../components/category-breakdown/types'

function currentMonthCosts(costs: Cost[]): Cost[] {
  const now = new Date()
  return costs.filter((c) => {
    const d = new Date(c.createdAt)
    return d.getFullYear() === now.getFullYear() && d.getMonth() === now.getMonth()
  })
}

export function buildCostsSummary(costs: Cost[]) {
  const monthly = currentMonthCosts(costs)
  const thisMonth = monthly.reduce((sum, c) => sum + c.totalValue, 0)
  const inSplits = monthly
    .filter((c) => c.groupId)
    .reduce((sum, c) => sum + c.totalValue, 0)
  const solo = monthly
    .filter((c) => !c.groupId)
    .reduce((sum, c) => sum + c.totalValue, 0)
  return { thisMonth, inSplits, solo }
}

export function buildCategoryBreakdown(costs: Cost[]): CategoryBreakdownPoint[] {
  const monthly = currentMonthCosts(costs)
  const map = new Map<string, number>()
  for (const cost of monthly) {
    map.set(cost.category, (map.get(cost.category) ?? 0) + cost.totalValue)
  }
  const grandTotal = Array.from(map.values()).reduce((sum, v) => sum + v, 0)
  return Array.from(map, ([category, total]) => ({
    category,
    total,
    percentage: grandTotal > 0 ? Math.round((total / grandTotal) * 100) : 0,
  })).sort((a, b) => b.total - a.total)
}
