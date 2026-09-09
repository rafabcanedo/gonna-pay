import type { ChartConfig } from '@/components/ui/chart'
import { TimeRange } from '../../../types';

export const areaChartConfig = {
  income: {
    label: 'Income',
    color: '#22c55e',
  },
  spending: {
    label: 'Spending',
    color: '#3b82f6',
  },
} satisfies ChartConfig

export const descriptionMap: Record<TimeRange, string> = {
  '7d': 'Showing income and spending for the last 7 days',
  '30d': 'Showing income and spending for the last 30 days',
  '90d': 'Showing income and spending for the last 3 months',
}

export const timeRanges: { value: TimeRange; label: string }[] = [
  { value: '7d', label: 'Last 7 days' },
  { value: '30d', label: 'Last 30 days' },
  { value: '90d', label: 'Last 3 months' },
]
