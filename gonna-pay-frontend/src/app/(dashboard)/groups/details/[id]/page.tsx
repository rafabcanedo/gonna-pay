import { dehydrate, HydrationBoundary, QueryClient } from '@tanstack/react-query'
import { GroupService } from '@/services'
import { ContactService } from '@/services'
import { GroupDetails } from './components/group-details'

export default async function GroupDetailPage({ params }: { params: Promise<{ id: string }> }) {
  const { id } = await params
  const queryClient = new QueryClient()

  await Promise.all([
    queryClient.prefetchQuery({
      queryKey: ['groups', id],
      queryFn: () => GroupService.getById(id),
    }),
    queryClient.prefetchQuery({
      queryKey: ['contacts'],
      queryFn: () => ContactService.getAll(),
    }),
  ])

  return (
    <HydrationBoundary state={dehydrate(queryClient)}>
      <GroupDetails groupId={id} />
    </HydrationBoundary>
  )
}
