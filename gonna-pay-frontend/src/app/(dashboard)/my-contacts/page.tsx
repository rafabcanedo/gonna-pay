import { AddContact } from "./components/add-contacts";
import { ContactService } from "@/services";
import { dehydrate, HydrationBoundary, QueryClient } from "@tanstack/react-query";
import { TableContact } from "./components/table-contacts";

export default async function MyContacts() {

  const queryClient = new QueryClient()

  await queryClient.prefetchQuery({
    queryKey: ['contacts'],
    queryFn: () => ContactService.getAll(),
  })

  return (
    <div className="px-8 w-full">
      <header className="flex flex-row items-center justify-between h-12 mt-4">
        <h1 className="font-montserrat text-xl text-zinc-600">
          Your contact list
        </h1>
        <AddContact />
      </header>

      <div className="mt-12">
        <HydrationBoundary state={dehydrate(queryClient)}>
          <TableContact />
        </HydrationBoundary>
      </div>
    </div>
  );
}
