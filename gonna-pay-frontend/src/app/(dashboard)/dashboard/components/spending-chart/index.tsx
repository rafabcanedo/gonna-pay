'use client'

import { Bar, BarChart, XAxis, YAxis } from 'recharts'
import {
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
} from '@/components/ui/chart'
import type { ChartConfig } from '@/components/ui/chart'
import type { IPropsSpendingChart } from './interfaces'

const chartConfig = {
  total: { label: 'Total', color: '#3b82f6' },
} satisfies ChartConfig

export function SpendingChart({ data }: IPropsSpendingChart) {
  return (
    <div className="rounded-xl border border-zinc-200 bg-white shadow-sm p-6">
      <h2 className="font-mono text-lg font-semibold text-zinc-600 mb-4">
        Spending by Month
      </h2>
      <ChartContainer config={chartConfig} className="h-[220px] w-full">
        <BarChart data={data}>
          <XAxis dataKey="month" tickLine={false} axisLine={false} />
          <YAxis
            tickLine={false}
            axisLine={false}
            tickFormatter={(v) => `$${v}`}
          />
          <ChartTooltip content={<ChartTooltipContent />} />
          <Bar dataKey="total" fill="var(--color-total)" radius={4} />
        </BarChart>
      </ChartContainer>
    </div>
  )
}
