export interface GraphDataItem {
  date: string;
  income: number;
  spending: number;
}

export type TimeRange = "7d" | "30d" | "90d";

export interface GraphConfig {
  title: string;
  description: string;
  data: GraphDataItem[];
  timeRanges: {
    value: TimeRange;
    label: string;
  }[];
}
