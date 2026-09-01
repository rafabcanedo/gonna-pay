"use client";

import * as React from "react";
import { Area, AreaChart, CartesianGrid, XAxis } from "recharts";

import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";

import {
  ChartContainer,
  ChartLegend,
  ChartLegendContent,
  ChartTooltip,
  ChartTooltipContent,
  type ChartConfig,
} from "@/components/ui/chart";

import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";

import { TimeRange } from "../../types";
import { graphConfig } from "../../lib/data";
import { filterDataByTimeRange, formatDate } from "../../lib/utils";

export const description = "Your Payments area";

const areaChartConfig = {
  income: {
    label: "Income",
    color: "#22c55e",
  },
  spending: {
    label: "Spending",
    color: "#3b82f6",
  },
} satisfies ChartConfig;

export const GraphPayments = () => {
  const [timeRange, setTimeRange] = React.useState<TimeRange>("30d");

  const filteredData = filterDataByTimeRange(
    graphConfig.data,
    timeRange,
    "2026-01-31"
  );

  return (
    <Card className="pt-0">
      <CardHeader className="flex items-center gap-2 space-y-0 border-b py-5 sm:flex-row">
        <div className="grid flex-1 gap-1">
          <CardTitle>{graphConfig.title}</CardTitle>
          <CardDescription>{graphConfig.description}</CardDescription>
        </div>
        <Select
          value={timeRange}
          onValueChange={(value) => setTimeRange(value as TimeRange)}
        >
          <SelectTrigger
            className="hidden w-[160px] rounded-lg sm:ml-auto sm:flex"
            aria-label="Select a value"
          >
            <SelectValue placeholder="Last 30 days" />
          </SelectTrigger>
          <SelectContent className="rounded-xl">
            {graphConfig.timeRanges.map((range) => (
              <SelectItem
                key={range.value}
                value={range.value}
                className="rounded-lg"
              >
                {range.label}
              </SelectItem>
            ))}
          </SelectContent>
        </Select>
      </CardHeader>

      <CardContent className="px-2 pt-4 sm:px-6 sm:pt-6">
        <ChartContainer
          config={areaChartConfig}
          className="aspect-auto h-[250px] w-full"
        >
          <AreaChart data={filteredData}>
            <defs>
              <linearGradient id="fillIncome" x1="0" y1="0" x2="0" y2="1">
                <stop
                  offset="5%"
                  stopColor={areaChartConfig.income.color}
                  stopOpacity={0.8}
                />
                <stop
                  offset="95%"
                  stopColor={areaChartConfig.income.color}
                  stopOpacity={0.1}
                />
              </linearGradient>
              <linearGradient id="fillSpending" x1="0" y1="0" x2="0" y2="1">
                <stop
                  offset="5%"
                  stopColor={areaChartConfig.spending.color}
                  stopOpacity={0.8}
                />
                <stop
                  offset="95%"
                  stopColor={areaChartConfig.spending.color}
                  stopOpacity={0.1}
                />
              </linearGradient>
            </defs>

            <CartesianGrid vertical={false} />

            <XAxis
              dataKey="date"
              tickLine={false}
              axisLine={false}
              tickMargin={8}
              minTickGap={32}
              tickFormatter={formatDate}
            />

            <ChartTooltip
              cursor={false}
              content={
                <ChartTooltipContent
                  labelFormatter={formatDate}
                  indicator="dot"
                />
              }
            />

            <Area
              dataKey="income"
              type="natural"
              fill="url(#fillIncome)"
              stroke={areaChartConfig.income.color}
              stackId={undefined}
            />

            <Area
              dataKey="spending"
              type="natural"
              fill="url(#fillSpending)"
              stroke={areaChartConfig.spending.color}
              stackId={undefined}
            />

            <ChartLegend content={<ChartLegendContent />} />
          </AreaChart>
        </ChartContainer>
      </CardContent>
    </Card>
  );
};
