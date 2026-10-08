import { AddMoney } from "@/components/add-money";
import { InitialHelper } from "@/components/initial-helper"
import { User, Coins, ScanBarcode } from 'lucide-react';
import { DashboardCards } from "./components/dashboard-cards";
import { RecentContacts } from "./components/recent-contacts";
import { RecentCosts } from "./components/recent-costs";
import { RemindersWidget } from "./components/reminders-widget";
import { SpendingChart } from "./components/spending-chart";
import { CategoryChart } from "./components/category-chart";
import { CostService, ContactService } from "@/services";
import { buildSpendingData, buildCategoryData, buildTotals } from "./lib/utils";

export default async function Dashboard() {
  const costsResponse = await CostService.getAll()
  const contactsResponse = await ContactService.getAll()

  const spendingData = buildSpendingData(costsResponse.data)
  const categoryData = buildCategoryData(costsResponse.data)
  const { amount, income, spending } = buildTotals(costsResponse.data)

  return (
    <div className="flex flex-col gap-8 py-8">
      <div className="flex flex-row gap-8">
        <DashboardCards title="Amount" value={amount} />
        <DashboardCards title="Income" value={income} />
        <DashboardCards title="Spending" value={spending} />
      </div>

      <div className="flex flex-row items-center justify-between gap-16 h-14">
        <InitialHelper name="My Costs" icon={Coins} link="/costs" />
        <AddMoney />
        <InitialHelper name="Payments" icon={ScanBarcode} link="/payments" />
        <InitialHelper name="My contacts" icon={User} link="/my-contacts" />
      </div>

      <div className="grid grid-cols-2 gap-8">
        <SpendingChart data={spendingData} />
        <CategoryChart data={categoryData} />
      </div>

      <div className="grid grid-cols-2 gap-8">
        <RecentCosts costs={costsResponse.data} total={costsResponse.total} />
        <RemindersWidget />
      </div>

      <RecentContacts contacts={contactsResponse.data} total={contactsResponse.total} />
    </div>
  )
}
