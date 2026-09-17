'use client'

import { useContactsQuery } from '@/hooks/queries/use-contact-query'
import { CategoryCards } from '../category-cards'

export const ContactsStats = () => {
  const { data } = useContactsQuery(1)
  const stats = data?.stats

  if (!stats) return null

  return <CategoryCards stats={stats} />
}
