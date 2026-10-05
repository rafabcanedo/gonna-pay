import { ContactCategory } from '@/types'
import { getCategoryIcon } from '@/utils/category-icons'

export const CONTACT_CATEGORIES = [
  { label: 'Family', value: ContactCategory.FAMILY, icon: getCategoryIcon(ContactCategory.FAMILY) },
  { label: 'Friends', value: ContactCategory.FRIEND, icon: getCategoryIcon(ContactCategory.FRIEND) },
  { label: 'Work', value: ContactCategory.WORK, icon: getCategoryIcon(ContactCategory.WORK) },
]

export const CONTACT_CATEGORY_OPTIONS = [
  { label: 'All', value: 'all' },
  ...CONTACT_CATEGORIES,
]
