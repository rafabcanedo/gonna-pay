import { TransactionCategory } from '@/types'

export const costCategoryOptions = [
  { label: 'Dinner', value: TransactionCategory.DINNER },
  { label: 'Lunch', value: TransactionCategory.LUNCH },
  { label: 'Entertainment', value: TransactionCategory.ENTERTAINMENT },
  { label: 'Travel', value: TransactionCategory.TRAVEL },
  { label: 'Others', value: TransactionCategory.OTHERS },
]

export const COST_CATEGORY_OPTIONS = [
  { label: 'All', value: 'all' },
  ...costCategoryOptions,
]

export const PERIOD_OPTIONS = [
  { label: 'All', value: 'all' },
  { label: 'This month', value: 'month' },
  { label: 'Last 7 days', value: '7days' },
]

export const TYPE_OPTIONS = [
  { label: 'All', value: 'all' },
  { label: 'Solo', value: 'solo' },
  { label: 'Group', value: 'group' },
]

export const MAX_VALUE = 10000
