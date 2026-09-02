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
import { Contact } from "@/types";

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

  const contacts = queryClient.getQueryData<Contact[]>(["contacts"])

  return (
    <div className="flex flex-col px-8 w-full">
      <header className="flex flex-row items-center justify-between h-12 mt-4">
        <h1 className="font-montserrat text-xl text-zinc-600">Your costs</h1>
        <Link href="/costs/create-cost">
          <Button variant="outline">Add Cost</Button>
        </Link>
      </header>

      <HydrationBoundary state={dehydrate(queryClient)}>
        <div className="flex justify-center items-center gap-8">
          {contacts?.map((item) => (
            <AvatarCost key={item.id} name={item.name} />
          ))}
        </div>

        <div className="mt-12">
          <CostsTable />
        </div>
      </HydrationBoundary>
    </div>
  );
}
