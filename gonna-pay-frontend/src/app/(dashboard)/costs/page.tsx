import { Button } from "@/components/ui/button";
import { CostsTable } from "./components/costs-table";
import Link from "next/link";
import {
  dehydrate,
  HydrationBoundary,
  QueryClient,
} from "@tanstack/react-query";
import { CostService } from "@/services";
import { AvatarCost } from "./components/avatar-cost";

export default async function Costs() {
  const queryClient = new QueryClient();

  await queryClient.prefetchQuery({
    queryKey: ["costs"],
    queryFn: () => CostService.getAll(),
  });

  const avatarList = [
    { id: 1, name: "Isa" },
    { id: 2, name: "Artur" },
    { id: 3, name: "Joao" },
    { id: 1, name: "Gabriela" },
    { id: 2, name: "Rafael" },
    { id: 3, name: "Daniela" },
  ]

  return (
    <div className="flex flex-col px-8 w-full">
      <header className="flex flex-row items-center justify-between h-12 mt-4">
        <h1 className="font-montserrat text-xl text-zinc-600">Your costs</h1>
        <Link href="/costs/create-cost">
          <Button variant="outline">Add Cost</Button>
        </Link>
      </header>

      <div className="flex justify-center items-center gap-8">
        {avatarList.map((item) => (
          <AvatarCost key={item.id} name={item.name} />
        ))}
      </div>

      <div className="mt-12">
        <HydrationBoundary state={dehydrate(queryClient)}>
          <CostsTable />
        </HydrationBoundary>
      </div>
    </div>
  );
}
