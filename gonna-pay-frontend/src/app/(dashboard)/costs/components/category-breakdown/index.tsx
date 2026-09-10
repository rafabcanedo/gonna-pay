import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import type { IPropsCategoryBreakdown } from './interfaces'

export function CategoryBreakdown({ data }: IPropsCategoryBreakdown) {
  return (
    <Card className="w-full">
      <CardHeader>
        <CardTitle className="text-sm text-zinc-500 font-light">Spending by category</CardTitle>
      </CardHeader>
      <CardContent className="flex flex-col gap-3">
        {data.map(({ category, percentage }) => (
          <div key={category} className="flex items-center gap-3">
            <span className="text-sm text-zinc-600 w-28 shrink-0">{category}</span>
            <div className="flex-1 h-2 rounded-full bg-zinc-100">
              <div
                className="h-2 rounded-full bg-zinc-700"
                style={{ width: `${percentage}%` }}
              />
            </div>
            <span className="text-xs text-zinc-400 w-8 text-right">{percentage}%</span>
          </div>
        ))}
      </CardContent>
    </Card>
  )
}
