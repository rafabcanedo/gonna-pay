import { TransactionCategory } from '@/types'

export const GROUP_CATEGORY_OPTIONS = [
  { label: 'All', value: 'all' },
  { label: 'Dinner', value: TransactionCategory.DINNER },
  { label: 'Lunch', value: TransactionCategory.LUNCH },
  { label: 'Entertainment', value: TransactionCategory.ENTERTAINMENT },
  { label: 'Travel', value: TransactionCategory.TRAVEL },
  { label: 'Others', value: TransactionCategory.OTHERS },
]
