import type { Cost } from '@/types'
import type { SpendingChartPoint } from '../components/spending-chart/types'
import type { CategoryChartPoint } from '../components/category-chart/types'

export function buildSpendingData(costs: Cost[]): SpendingChartPoint[] {
  const today = new Date()
  return Array.from({ length: 7 }, (_, i) => {
    const date = new Date(today)
    date.setDate(today.getDate() - (6 - i))
    const day = date.toLocaleString('en-US', { weekday: 'short' })
    const dateStr = date.toDateString()
    const dayCosts = costs.filter((c) => new Date(c.createdAt).toDateString() === dateStr)
    const spending = dayCosts.reduce((sum, c) => sum + c.ownerValue, 0)
    const income = dayCosts
      .filter((c) => c.groupId)
      .reduce((sum, c) => sum + (c.totalValue - c.ownerValue), 0)
    return { day, spending, income }
  })
}

export function buildCategoryData(costs: Cost[]): CategoryChartPoint[] {
  const map = new Map<string, number>()
  for (const cost of costs) {
    map.set(cost.category, (map.get(cost.category) ?? 0) + cost.totalValue)
  }
  return Array.from(map, ([category, total]) => ({ category, total }))
}

export function buildTotals(costs: Cost[]) {
  const now = new Date()
  const monthly = costs.filter((c) => {
    const d = new Date(c.createdAt)
    return d.getFullYear() === now.getFullYear() && d.getMonth() === now.getMonth()
  })
  const amount = monthly.reduce((sum, c) => sum + c.totalValue, 0)
  const income = monthly
    .filter((c) => c.groupId)
    .reduce((sum, c) => sum + (c.totalValue - c.ownerValue), 0)
  const spending = monthly
    .filter((c) => c.groupId)
    .reduce((sum, c) => sum + c.ownerValue, 0)
  return { amount, income, spending }
}
