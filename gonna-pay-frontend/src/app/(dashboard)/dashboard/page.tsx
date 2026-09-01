import { AddMoney } from "@/components/add-money";
import { InitialHelper } from "@/components/initial-helper"
import { User, WalletMinimal, ScanBarcode } from 'lucide-react';
import { DashboardCards } from "./components/dashboard-cards";
import { RecentContacts } from "./components/recent-contacts";
import { RecentCosts } from "./components/recent-costs";
import { Reminders } from "./components/reminders";
import { GetContactsResponse, GetCostsResponse } from "@/types";
import { apiCall } from "@/lib/api-client";

export default async function Dashboard() {

    const costs = await apiCall<GetCostsResponse>('/costs')
    const contacts = await apiCall<GetContactsResponse>('/contacts')

    return (
        <div className="flex flex-col">
            <div className="flex flex-row mt-8 gap-8">
                <DashboardCards title="Amout" value={3200} />
                <DashboardCards title="Money box" value={800} />
            </div>
            <div className="flex flex-row items-center justify-between gap-16 h-14 mt-12">
                <InitialHelper name="My wallet" icon={WalletMinimal} link="/my-wallet" />
                <AddMoney />
                <InitialHelper name="Payments" icon={ScanBarcode} link="/payments" />
                <InitialHelper name="My contacts" icon={User} link="/my-contacts" />
            </div>

            <div className="flex flex-col mt-10">
                <div>
                    <Reminders />
                </div>

                <div>
                    <RecentCosts costs={costs} total={costs.length} />
                </div>

                <div>
                    <RecentContacts contacts={contacts} total={contacts.length} />
                </div>
            </div>
        </div>
    )
}
