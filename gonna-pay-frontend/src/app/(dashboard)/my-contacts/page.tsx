import { AddContact } from "./components/add-contacts";
import { ContactService } from "@/services";
import { dehydrate, HydrationBoundary, QueryClient } from "@tanstack/react-query";
import { TableContact } from "./components/table-contacts";
import { ContactsStats } from "./components/contacts-stats";
import { TopContacts } from "./components/top-contacts";

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

  return (
    <div className="px-8 w-full">
      <header className="flex flex-row items-center justify-between h-12 mt-4">
        <h1 className="font-montserrat text-xl text-zinc-600">
          Your contact list
        </h1>
        <AddContact />
      </header>

      <HydrationBoundary state={dehydrate(queryClient)}>
        <div className="flex flex-row items-start gap-4 mt-8">
          <ContactsStats />
          <TopContacts />
        </div>

        <div className="mt-8">
          <TableContact />
        </div>
      </HydrationBoundary>
    </div>
  );
}
