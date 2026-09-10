import { AddContact } from "./components/add-contacts";
import { ContactService } from "@/services";
import { dehydrate, HydrationBoundary, QueryClient } from "@tanstack/react-query";
import { TableContact } from "./components/table-contacts";
import { CategoryCards } from "./components/category-cards";
import { TopContacts } from "./components/top-contacts";
import type { Contact } from "@/types";

export default async function MyContacts() {
  const queryClient = new QueryClient()

  await queryClient.prefetchQuery({
    queryKey: ['contacts'],
    queryFn: () => ContactService.getAll(),
  })

  await queryClient.prefetchQuery({
    queryKey: ['contacts', 'frequency', 5],
    queryFn: () => ContactService.getFrequency(5),
  })

  const contacts = queryClient.getQueryData<Contact[]>(['contacts']) ?? []

  return (
    <div className="px-8 w-full">
      <header className="flex flex-row items-center justify-between h-12 mt-4">
        <h1 className="font-montserrat text-xl text-zinc-600">
          Your contact list
        </h1>
        <AddContact />
      </header>

      <div className="flex flex-row items-start gap-4 mt-8">
        <CategoryCards contacts={contacts} />
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
