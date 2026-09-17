import { dehydrate, HydrationBoundary, QueryClient } from '@tanstack/react-query'
import { CostService, ContactService } from '@/services'
import { GraphPayments } from './components/graph-payments'
import { TablePayments } from './components/table-payments'
import { TableCostSplits } from './components/table-cost-splits'

export default async function Payments() {
  const queryClient = new QueryClient()

  await queryClient.prefetchQuery({
    queryKey: ['costs', { page: 1, limit: 20 }],
    queryFn: () => CostService.getAll(new URLSearchParams({ page: "1", limit: "20" })),
  })

  await queryClient.prefetchQuery({
    queryKey: ['contacts', { page: 1, limit: 20 }],
    queryFn: () => ContactService.getAll(new URLSearchParams({ page: "1", limit: "20" })),
  })

  return (
    <div className="flex flex-col p-8 w-full">
      <header>
        <h1 className="font-montserrat text-xl text-zinc-600">Your payments</h1>
      </header>

      <HydrationBoundary state={dehydrate(queryClient)}>
        <div className="mt-12">
          <GraphPayments />
        </div>

        <TablePayments />

        <TableCostSplits />
      </HydrationBoundary>
    </div>
  )
}
