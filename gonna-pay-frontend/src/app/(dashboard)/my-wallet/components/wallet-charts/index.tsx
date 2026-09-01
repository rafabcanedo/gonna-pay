"use client";

import { TrendingUp, TrendingDown } from "lucide-react";
import { Pie, PieChart, Cell } from "recharts";

import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import {
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
} from "@/components/ui/chart";
import { WalletChartsProps } from './interfaces'

export const WalletCharts = ({ config }: WalletChartsProps) => {
  const chartData = config.data.map((item) => ({
    ...item,
    fill: item.color,
  }));

  const TrendIcon = config.footer?.trend === "up" ? TrendingUp : TrendingDown;

  return (
    <Card className="flex flex-col max-w-sm">
      <CardHeader className="items-center pb-0">
        <CardTitle>{config.title}</CardTitle>
        <CardDescription>{config.description}</CardDescription>
      </CardHeader>

      <CardContent className="flex-1 pb-0">
        <ChartContainer
          config={{}}
          className="mx-auto aspect-square max-h-[250px]"
        >
          <PieChart width={250} height={250}>
            <ChartTooltip content={<ChartTooltipContent hideLabel />} />
            <Pie
              data={chartData}
              dataKey="value"
              nameKey="name"
              cx="50%"
              cy="50%"
              outerRadius={80}
              label
            >
              {chartData.map((entry, index) => (
                <Cell key={`cell-${index}`} fill={entry.fill} />
              ))}
            </Pie>
          </PieChart>
        </ChartContainer>
      </CardContent>

      {config.footer && (
        <CardFooter className="flex-col gap-2 text-sm">
          <div className="flex items-center gap-2 leading-none font-medium">
            Trending {config.footer.trend} by {config.footer.trendValue}% this
            month <TrendIcon className="h-4 w-4" />
          </div>
          <div className="text-muted-foreground leading-none">
            {config.footer.subtitle}
          </div>
        </CardFooter>
      )}
    </Card>
  );
};
