import type { Cost } from '@/types'
import type { GraphDataItem, TimeRange } from '../types'

const daysMap: Record<TimeRange, number> = {
  '7d': 7,
  '30d': 30,
  '90d': 90,
}

export function formatDate(dateString: string): string {
  const date = new Date(dateString)
  return date.toLocaleDateString('en-US', { month: 'short', day: 'numeric' })
}

export function buildPaymentsData(costs: Cost[], timeRange: TimeRange): GraphDataItem[] {
  const days = daysMap[timeRange]
  const now = new Date()

  return Array.from({ length: days }, (_, i) => {
    const date = new Date(now)
    date.setDate(now.getDate() - (days - 1 - i))
    const dateStr = date.toISOString().split('T')[0]

    const dayCosts = costs.filter(c => c.createdAt.startsWith(dateStr))
    const spending = dayCosts.reduce((sum, c) => sum + c.ownerValue, 0)
    const income = dayCosts
      .filter(c => Boolean(c.groupId))
      .reduce((sum, c) => sum + (c.totalValue - c.ownerValue), 0)

    return { date: dateStr, income, spending }
  })
}
