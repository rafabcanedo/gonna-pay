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
import { CostsSummary } from "./components/costs-summary";
import { CategoryBreakdown } from "./components/category-breakdown";
import type { Contact, GetCostsResponse } from "@/types";

export default async function Costs() {
  const queryClient = new QueryClient();

  await queryClient.prefetchQuery({
    queryKey: ["costs"],
    queryFn: () => CostService.getAll(),
  });

  await queryClient.prefetchQuery({
    queryKey: ["contacts"],
    queryFn: () => ContactService.getAll(),
  });

  const contacts = queryClient.getQueryData<Contact[]>(["contacts"]) ?? []
  const costsResponse = queryClient.getQueryData<GetCostsResponse>(["costs"])
  const stats = costsResponse?.stats

  return (
    <div className="flex flex-col px-8 w-full">
      <header className="flex flex-row items-center justify-between h-12 mt-4">
        <h1 className="font-montserrat text-xl text-zinc-600">Your costs</h1>
        <Link href="/costs/create-cost">
          <Button variant="outline">Add Cost</Button>
        </Link>
      </header>

      <div className="flex flex-row gap-4 mt-8">
        {stats && <CostsSummary thisMonth={stats.thisMonth} inSplits={stats.inSplits} solo={stats.solo} />}
      </div>

      {stats && stats.byCategory.length > 0 && (
        <div className="mt-4">
          <CategoryBreakdown data={stats.byCategory} />
        </div>
      )}

      <HydrationBoundary state={dehydrate(queryClient)}>
        <div className="flex justify-center items-center gap-8 mt-4">
          {contacts.map((item) => (
            <AvatarCost key={item.id} name={item.name} />
          ))}
        </div>

        <div className="mt-8">
          <CostsTable />
        </div>
      </HydrationBoundary>
    </div>
  );
}
