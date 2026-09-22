import { ContactCategory } from '@/types'

export const CONTACT_CATEGORIES = [
  { label: 'Family', value: ContactCategory.FAMILY },
  { label: 'Friends', value: ContactCategory.FRIEND },
  { label: 'Work', value: ContactCategory.WORK },
]

export const CONTACT_CATEGORY_OPTIONS = [
  { label: 'All', value: 'all' },
  ...CONTACT_CATEGORIES,
]
