import { Button } from "@/components/ui/button";
import { CostsTable } from "./components/costs-table";
import Link from "next/link";
import {
  dehydrate,
  HydrationBoundary,
  QueryClient,
} from "@tanstack/react-query";
import { ContactService, CostService } from "@/services";
import { AvatarCost } from "./components/avatar-cost";
import { CostsStats } from "./components/costs-stats";
import type { GetContactsResponse } from "@/types";

export default async function Costs() {
  const queryClient = new QueryClient();

  await queryClient.prefetchQuery({
    queryKey: ["costs", { page: 1, limit: 20 }],
    queryFn: () => CostService.getAll(new URLSearchParams({ page: "1", limit: "20" })),
  });

  await queryClient.prefetchQuery({
    queryKey: ["contacts", { page: 1, limit: 20 }],
    queryFn: () => ContactService.getAll(new URLSearchParams({ page: "1", limit: "20" })),
  });

  const contacts = queryClient.getQueryData<GetContactsResponse>(["contacts", { page: 1, limit: 20 }])?.data ?? []

  return (
    <div className="flex flex-col px-8 w-full">
      <header className="flex flex-row items-center justify-between h-12 mt-4">
        <h1 className="font-montserrat text-xl text-zinc-600">Your costs</h1>
        <Link href="/costs/create-cost">
          <Button variant="outline">Add Cost</Button>
        </Link>
      </header>

      {contacts.length > 0 && (
        <div className="flex flex-row items-center justify-center gap-8 mt-4">
          {contacts.map((item) => (
            <AvatarCost key={item.id} name={item.name} />
          ))}
        </div>
      )}

      <HydrationBoundary state={dehydrate(queryClient)}>
        <CostsStats />

        <div className="mt-8">
          <CostsTable />
        </div>
      </HydrationBoundary>
    </div>
  );
}
