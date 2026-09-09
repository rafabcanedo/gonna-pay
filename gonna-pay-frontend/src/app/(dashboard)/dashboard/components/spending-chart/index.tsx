'use client'

import { Bar, BarChart, CartesianGrid, XAxis, YAxis } from 'recharts'
import {
  ChartContainer,
  ChartLegend,
  ChartLegendContent,
  ChartTooltip,
  ChartTooltipContent,
} from '@/components/ui/chart'
import type { ChartConfig } from '@/components/ui/chart'
import type { IPropsSpendingChart } from './interfaces'

const chartConfig = {
  spending:   { label: 'Spending',   color: '#3b82f6' },
  receivable: { label: 'Receivable', color: '#22c55e' },
} satisfies ChartConfig

export function SpendingChart({ data }: IPropsSpendingChart) {
  return (
    <div className="rounded-xl border border-zinc-200 bg-white shadow-sm p-6">
      <h2 className="font-mono text-lg font-semibold text-zinc-600 mb-4">
        Last 7 Days
      </h2>
      <ChartContainer config={chartConfig} className="h-[220px] w-full">
        <BarChart accessibilityLayer data={data}>
          <CartesianGrid vertical={false} />
          <XAxis dataKey="day" tickLine={false} axisLine={false} tickMargin={10} />
          <YAxis
            tickLine={false}
            axisLine={false}
            tickFormatter={(v) => `$${v}`}
          />
          <ChartTooltip content={<ChartTooltipContent />} />
          <ChartLegend content={<ChartLegendContent />} />
          <Bar dataKey="spending"   fill="var(--color-spending)"   radius={4} />
          <Bar dataKey="receivable" fill="var(--color-receivable)" radius={4} />
        </BarChart>
      </ChartContainer>
    </div>
  )
}
