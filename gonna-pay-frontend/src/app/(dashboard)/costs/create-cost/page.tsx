import CreateCostForm from "./components/create-cost-form"
import { dehydrate, HydrationBoundary, QueryClient } from "@tanstack/react-query"
import { GroupService } from "@/services"

export default async function CreateCost() {
    const queryClient = new QueryClient()

    await queryClient.prefetchQuery({
        queryKey: ['groups'],
        queryFn: () => GroupService.getAll(),
    })

    return (
        <div className="flex-1 flex items-center justify-center min-h-screen">
            <HydrationBoundary state={dehydrate(queryClient)}>
                <CreateCostForm />
            </HydrationBoundary>
        </div>
    )
}
