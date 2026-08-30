import { Skeleton } from "@/components/ui/skeleton"
import { LoadingSkeletonProps } from './interfaces'

export const ApiLoading = ({
    type = 'table',
    rows = 5,
    columns = 4
}: LoadingSkeletonProps) => {
    if (type === 'card') {
        return (
            <div className="space-y-4 p-6">
                {Array.from({ length: rows }).map((_, i) => (
                    <div key={i} className="flex items-center gap-4 p-4 border rounded-lg">
                        <Skeleton className="h-12 w-12 rounded-full" />
                        <div className="space-y-2 flex-1">
                            <Skeleton className="h-4 w-[250px]" />
                            <Skeleton className="h-4 w-[200px]" />
                        </div>
                    </div>
                ))}
            </div>
        )
    }

    return (
        <div className="space-y-4 p-6">
            <div className="flex gap-4 pb-4 border-b">
                {Array.from({ length: columns }).map((_, i) => (
                    <Skeleton key={i} className="h-4 flex-1" />
                ))}
            </div>
            {Array.from({ length: rows }).map((_, rowIndex) => (
                <div key={rowIndex} className="flex gap-4 items-center">
                    {Array.from({ length: columns }).map((_, colIndex) => (
                        <Skeleton key={colIndex} className="h-8 flex-1" />
                    ))}
                </div>
            ))}
        </div>
    )
}