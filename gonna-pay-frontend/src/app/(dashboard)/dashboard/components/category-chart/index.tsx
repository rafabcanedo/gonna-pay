'use client'

import { Pie, PieChart, Cell } from 'recharts'
import {
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
} from '@/components/ui/chart'
import type { IPropsCategoryChart } from './interfaces'
import { CATEGORY_COLORS, chartConfig } from './constants'

export function CategoryChart({ data }: IPropsCategoryChart) {
  return (
    <div className="rounded-xl border border-zinc-200 bg-white shadow-sm p-6">
      <h2 className="font-mono text-lg font-semibold text-zinc-600 mb-4">
        By Category
      </h2>
      <ChartContainer config={chartConfig} className="h-[220px] w-full">
        <PieChart>
          <ChartTooltip content={<ChartTooltipContent nameKey="category" />} />
          <Pie
            data={data}
            dataKey="total"
            nameKey="category"
            innerRadius={55}
            outerRadius={85}
          >
            {data.map((entry) => (
              <Cell
                key={entry.category}
                fill={CATEGORY_COLORS[entry.category] ?? '#6b7280'}
              />
            ))}
          </Pie>
        </PieChart>
      </ChartContainer>
      <div className="flex flex-wrap gap-3 mt-2">
        {data.map((entry) => (
          <div key={entry.category} className="flex items-center gap-1.5 text-xs text-zinc-600">
            <div
              className="h-2 w-2 rounded-full shrink-0"
              style={{ backgroundColor: CATEGORY_COLORS[entry.category] ?? '#6b7280' }}
            />
            {entry.category}
          </div>
        ))}
      </div>
    </div>
  )
}
