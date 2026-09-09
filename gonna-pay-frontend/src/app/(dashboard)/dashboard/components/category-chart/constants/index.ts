import type { ChartConfig } from '@/components/ui/chart'

export const CATEGORY_COLORS: Record<string, string> = {
  Dinner:        '#f97316',
  Lunch:         '#eab308',
  Entertainment: '#8b5cf6',
  Travel:        '#3b82f6',
  Others:        '#6b7280',
}

export const chartConfig = {
  total: { label: 'Total' },
} satisfies ChartConfig
