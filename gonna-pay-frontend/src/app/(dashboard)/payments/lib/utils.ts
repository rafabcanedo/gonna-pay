import { GraphDataItem, TimeRange } from "../types";

export const filterDataByTimeRange = (
  data: GraphDataItem[],
  timeRange: TimeRange,
  referenceDate: string = "2026-01-31"
): GraphDataItem[] => {
  const daysMap: Record<TimeRange, number> = {
    "7d": 7,
    "30d": 30,
    "90d": 90,
  };

  const daysToSubtract = daysMap[timeRange];
  const reference = new Date(referenceDate);
  const startDate = new Date(reference);
  startDate.setDate(startDate.getDate() - daysToSubtract);

  return data.filter((item) => {
    const date = new Date(item.date);
    return date >= startDate;
  });
};

export const formatDate = (dateString: string): string => {
  const date = new Date(dateString);
  return date.toLocaleDateString("en-US", {
    month: "short",
    day: "numeric",
  });
};
