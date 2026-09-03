import type { TransactionCategory } from '@/types'

export type CostDetailsForm = {
  costName: string
  totalValue: string
  category: TransactionCategory
  ownerPercentage: string
}
