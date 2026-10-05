import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import type { IPropsCategoryBreakdown } from './interfaces'
import { CATEGORY_BAR_COLORS, DEFAULT_BAR_COLOR } from './constants'
import { getCategoryIcon } from '@/utils/category-icons'

export function CategoryBreakdown({ data }: IPropsCategoryBreakdown) {
  return (
    <Card className="w-full">
      <CardHeader>
        <CardTitle className="text-sm text-zinc-500 font-light">Spending by category</CardTitle>
      </CardHeader>
      <CardContent className="flex flex-col gap-3">
        {data.map(({ category, percentage }) => {
          const Icon = getCategoryIcon(category)

          return (
            <div key={category} className="flex items-center gap-3">
              <span className="flex items-center gap-1.5 text-sm text-zinc-600 w-28 shrink-0">
                {Icon && <Icon className="w-3.5 h-3.5" />}
                {category}
              </span>
              <div className="flex-1 h-2 rounded-full bg-zinc-100">
                <div
                  className={`h-2 rounded-full ${CATEGORY_BAR_COLORS[category] ?? DEFAULT_BAR_COLOR}`}
                  style={{ width: `${percentage}%` }}
                />
              </div>
              <span className="text-xs text-zinc-400 w-8 text-right">{percentage}%</span>
            </div>
          )
        })}
      </CardContent>
    </Card>
  )
}
