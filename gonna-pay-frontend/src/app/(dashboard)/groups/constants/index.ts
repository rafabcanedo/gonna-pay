import { TransactionCategory } from '@/types'
import { getCategoryIcon } from '@/utils/category-icons'

export const GROUP_CATEGORY_OPTIONS = [
  { label: 'All', value: 'all' },
  { label: 'Dinner', value: TransactionCategory.DINNER, icon: getCategoryIcon(TransactionCategory.DINNER) },
  { label: 'Lunch', value: TransactionCategory.LUNCH, icon: getCategoryIcon(TransactionCategory.LUNCH) },
  { label: 'Entertainment', value: TransactionCategory.ENTERTAINMENT, icon: getCategoryIcon(TransactionCategory.ENTERTAINMENT) },
  { label: 'Travel', value: TransactionCategory.TRAVEL, icon: getCategoryIcon(TransactionCategory.TRAVEL) },
  { label: 'Others', value: TransactionCategory.OTHERS, icon: getCategoryIcon(TransactionCategory.OTHERS) },
]
