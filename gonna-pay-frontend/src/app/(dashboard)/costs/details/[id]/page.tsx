import { dehydrate, HydrationBoundary, QueryClient } from '@tanstack/react-query'
import { CostService } from '@/services'
import { CostDetails } from './components/cost-details'

export default async function CostDetailPage({ params }: { params: Promise<{ id: string }> }) {
  const { id } = await params
  const queryClient = new QueryClient()

  await queryClient.prefetchQuery({
    queryKey: ['costs', id],
    queryFn: () => CostService.getById(id),
  })

  return (
    <HydrationBoundary state={dehydrate(queryClient)}>
      <CostDetails costId={id} />
    </HydrationBoundary>
  )
}
