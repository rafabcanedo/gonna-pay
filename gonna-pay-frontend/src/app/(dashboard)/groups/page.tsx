import { dehydrate, HydrationBoundary, QueryClient } from '@tanstack/react-query'
import { GroupService } from '@/services'
import { CreateGroup } from './components/create-group'
import { TableGroup } from './components/table-groups'

export default async function Groups() {
  const queryClient = new QueryClient()

  await queryClient.prefetchQuery({
    queryKey: ['groups'],
    queryFn: () => GroupService.getAll(),
  })

  return (
    <div className="flex flex-col gap-4 px-8 w-full">
      <header className="flex flex-row items-center justify-between h-12 mt-4">
        <h1 className="font-montserrat text-xl text-zinc-600">Your groups</h1>
        <CreateGroup />
      </header>

      <HydrationBoundary state={dehydrate(queryClient)}>
        <TableGroup />
      </HydrationBoundary>
    </div>
  )
}
