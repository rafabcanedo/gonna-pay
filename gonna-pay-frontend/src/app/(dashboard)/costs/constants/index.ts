import { TransactionCategory } from '@/types'
import { getCategoryIcon } from '@/utils/category-icons'

export const COST_CATEGORY_FORM_OPTIONS = [
  { label: 'Dinner', value: TransactionCategory.DINNER, icon: getCategoryIcon(TransactionCategory.DINNER) },
  { label: 'Lunch', value: TransactionCategory.LUNCH, icon: getCategoryIcon(TransactionCategory.LUNCH) },
  { label: 'Entertainment', value: TransactionCategory.ENTERTAINMENT, icon: getCategoryIcon(TransactionCategory.ENTERTAINMENT) },
  { label: 'Travel', value: TransactionCategory.TRAVEL, icon: getCategoryIcon(TransactionCategory.TRAVEL) },
  { label: 'Others', value: TransactionCategory.OTHERS, icon: getCategoryIcon(TransactionCategory.OTHERS) },
]

export const COST_CATEGORY_OPTIONS = [
  { label: 'All', value: 'all' },
  ...COST_CATEGORY_FORM_OPTIONS,
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
