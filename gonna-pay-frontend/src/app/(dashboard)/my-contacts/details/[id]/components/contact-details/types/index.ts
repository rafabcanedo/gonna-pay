import type { ContactCategory } from '@/types'

export type EditContactForm = {
  name?: string
  email?: string
  phone?: string
  category?: ContactCategory
}
