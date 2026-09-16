import { AddContact } from "./components/add-contacts";
import { ContactService } from "@/services";
import { dehydrate, HydrationBoundary, QueryClient } from "@tanstack/react-query";
import { TableContact } from "./components/table-contacts";
import { CategoryCards } from "./components/category-cards";
import { TopContacts } from "./components/top-contacts";
import type { GetContactsResponse } from "@/types";

export default async function MyContacts() {
  const queryClient = new QueryClient()

  await queryClient.prefetchQuery({
    queryKey: ['contacts', { page: 1, limit: 20 }],
    queryFn: () => ContactService.getAll(new URLSearchParams({ page: "1", limit: "20" })),
  })

  await queryClient.prefetchQuery({
    queryKey: ['contacts', 'frequency', 5],
    queryFn: () => ContactService.getFrequency(5),
  })

  const contactsResponse = queryClient.getQueryData<GetContactsResponse>(['contacts', { page: 1, limit: 20 }])
  const stats = contactsResponse?.stats

  return (
    <div className="px-8 w-full">
      <header className="flex flex-row items-center justify-between h-12 mt-4">
        <h1 className="font-montserrat text-xl text-zinc-600">
          Your contact list
        </h1>
        <AddContact />
      </header>

      <div className="flex flex-row items-start gap-4 mt-8">
        {stats && <CategoryCards stats={stats} />}
        <HydrationBoundary state={dehydrate(queryClient)}>
          <TopContacts />
        </HydrationBoundary>
      </div>

      <div className="mt-8">
        <HydrationBoundary state={dehydrate(queryClient)}>
          <TableContact />
        </HydrationBoundary>
      </div>
    </div>
  );
}
