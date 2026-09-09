import { TimeRange } from '../../../types';

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
