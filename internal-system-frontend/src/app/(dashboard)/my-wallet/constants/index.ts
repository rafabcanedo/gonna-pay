import { ChartConfig } from "../types";

export const chartConfigs: Record<string, ChartConfig> = {
  investments: {
    title: "Investments",
    description: "Current Portfolio Distribution",
    data: [
      { name: "Stocks", value: 5000, color: "#22c55e" },
      { name: "Fixed Income", value: 3000, color: "#3b82f6" },
      { name: "Crypto", value: 1000, color: "#f59e0b" },
    ],
    footer: {
      trend: "up",
      trendValue: 12.5,
      subtitle: "Total portfolio value for last month",
    },
  },

  business: {
    title: "Business Expenses",
    description: "Your business",
    data: [
      { name: "Marketing", value: 1200, color: "#16a34a" },
      { name: "Software/SaaS", value: 800, color: "#0d9488" },
      { name: "Operations", value: 500, color: "#0891b2" },
      { name: "Consulting", value: 300, color: "#0284c7" },
      { name: "Others", value: 200, color: "#64748b" },
    ],
    footer: {
      trend: "up",
      trendValue: 3.2,
      subtitle: "Total business expenses for the last month",
    },
  },

  spending: {
    title: "Personal Spending",
    description: "Monthly expense breakdown",
    data: [
      { name: "Food & Dining", value: 800, color: "#3b82f6" },
      { name: "Transportation", value: 400, color: "#06b6d4" },
      { name: "Entertainment", value: 300, color: "#10b981" },
      { name: "Shopping", value: 500, color: "#2dd4bf" },
      { name: "Bills & Utilities", value: 600, color: "#60a5fa" },
      { name: "Others", value: 200, color: "#94a3b8" },
    ],
    footer: {
      trend: "down",
      trendValue: 5.8,
      subtitle: "Total personal spending for last month",
    },
  },
};
