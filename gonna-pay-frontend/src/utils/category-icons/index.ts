import {
  UserGroup,
  User,
  BriefcaseBusiness,
  HandPlatter,
  Salad,
  Clapperboard,
  Plane,
  Tag,
} from 'lucide-react'
import type { LucideIcon } from 'lucide-react'
import { ContactCategory, TransactionCategory } from '@/types'

export const CATEGORY_ICONS: Record<string, LucideIcon> = {
  [ContactCategory.FAMILY]: UserGroup,
  [ContactCategory.FRIEND]: User,
  [ContactCategory.WORK]: BriefcaseBusiness,

  [TransactionCategory.DINNER]: HandPlatter,
  [TransactionCategory.LUNCH]: Salad,
  [TransactionCategory.ENTERTAINMENT]: Clapperboard,
  [TransactionCategory.TRAVEL]: Plane,
  [TransactionCategory.OTHERS]: Tag,
}

export function getCategoryIcon(category: string): LucideIcon | undefined {
  return CATEGORY_ICONS[category]
}
