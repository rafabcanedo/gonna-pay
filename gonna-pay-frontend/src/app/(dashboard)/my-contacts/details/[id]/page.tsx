import { dehydrate, HydrationBoundary, QueryClient } from '@tanstack/react-query'
import { ContactService } from '@/services'
import { ContactDetails } from './components/contact-details'

export default async function ContactDetailPage({ params }: { params: Promise<{ id: string }> }) {
  const { id } = await params
  const queryClient = new QueryClient()

  await queryClient.prefetchQuery({
    queryKey: ['contacts', id],
    queryFn: () => ContactService.getById(id),
  })

  return (
    <HydrationBoundary state={dehydrate(queryClient)}>
      <ContactDetails contactId={id} />
    </HydrationBoundary>
  )
}
