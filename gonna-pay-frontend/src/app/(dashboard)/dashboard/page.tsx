import { AddMoney } from "@/components/add-money";
import { InitialHelper } from "@/components/initial-helper"
import { User, WalletMinimal, ScanBarcode } from 'lucide-react';
import { DashboardCards } from "./components/dashboard-cards";
import { RecentContacts } from "./components/recent-contacts";
import { RecentCosts } from "./components/recent-costs";
import { RemindersWidget } from "./components/reminders-widget";
import { SpendingChart } from "./components/spending-chart";
import { CategoryChart } from "./components/category-chart";
import { GetContactsResponse, GetCostsResponse, Cost } from "@/types";
import { apiCall } from "@/lib/api-client";
import type { SpendingChartPoint } from "./components/spending-chart/types";
import type { CategoryChartPoint } from "./components/category-chart/types";

function buildSpendingData(costs: Cost[]): SpendingChartPoint[] {
  const map = new Map<string, number>()
  for (const cost of costs) {
    const month = new Date(cost.createdAt).toLocaleString('en-US', { month: 'short' })
    map.set(month, (map.get(month) ?? 0) + cost.totalValue)
  }
  return Array.from(map, ([month, total]) => ({ month, total }))
}

function buildCategoryData(costs: Cost[]): CategoryChartPoint[] {
  const map = new Map<string, number>()
  for (const cost of costs) {
    map.set(cost.category, (map.get(cost.category) ?? 0) + cost.totalValue)
  }
  return Array.from(map, ([category, total]) => ({ category, total }))
}

export default async function Dashboard() {
  const costs = await apiCall<GetCostsResponse>('/costs')
  const contacts = await apiCall<GetContactsResponse>('/contacts')

  const spendingData = buildSpendingData(costs)
  const categoryData = buildCategoryData(costs)

  return (
    <div className="flex flex-col gap-8 py-8">
      <div className="flex flex-row gap-8">
        <DashboardCards title="Amount" value={3200} />
        <DashboardCards title="Money box" value={800} />
      </div>

      <div className="flex flex-row items-center justify-between gap-16 h-14">
        <InitialHelper name="My wallet" icon={WalletMinimal} link="/my-wallet" />
        <AddMoney />
        <InitialHelper name="Payments" icon={ScanBarcode} link="/payments" />
        <InitialHelper name="My contacts" icon={User} link="/my-contacts" />
      </div>

      <div className="grid grid-cols-2 gap-8">
        <SpendingChart data={spendingData} />
        <CategoryChart data={categoryData} />
      </div>

      <div className="grid grid-cols-2 gap-8">
        <RecentCosts costs={costs} total={costs.length} />
        <RemindersWidget />
      </div>

      <RecentContacts contacts={contacts} total={contacts.length} />
    </div>
  )
}
