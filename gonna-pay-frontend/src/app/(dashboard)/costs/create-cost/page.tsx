import { SideBar } from "@/components/side-bar"
import { SidebarProvider } from "@/components/ui/sidebar"
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
        <SidebarProvider className="px-8">
            <SideBar />
            <div className="flex-1 flex items-center justify-center min-h-screen">
                <HydrationBoundary state={dehydrate(queryClient)}>
                    <CreateCostForm />
                </HydrationBoundary>
            </div>
        </SidebarProvider>
    )
}