export interface ChartDataItem {
  name: string;
  value: number;
  color: string;
}

export type TrendDirection = "up" | "down";

export interface ChartFooter {
  trend: TrendDirection;
  trendValue: number;
  subtitle: string;
}

export interface ChartConfig {
  title: string;
  description: string;
  data: ChartDataItem[];
  footer?: ChartFooter;
}
