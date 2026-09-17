import { TransactionCategory } from '@/types'

export const CATEGORY_BAR_COLORS: Record<string, string> = {
  [TransactionCategory.DINNER]:        'bg-blue-500',
  [TransactionCategory.LUNCH]:         'bg-green-500',
  [TransactionCategory.ENTERTAINMENT]: 'bg-emerald-500',
  [TransactionCategory.TRAVEL]:        'bg-orange-500',
  [TransactionCategory.OTHERS]:        'bg-sky-500',
}

export const DEFAULT_BAR_COLOR = 'bg-zinc-400'
