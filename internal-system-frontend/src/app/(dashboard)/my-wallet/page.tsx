import { chartConfigs } from "./constants";
import { WalletCards } from "./components/wallet-cards";
import { WalletCharts } from "./components/wallet-charts";

export default function MyWallet() {
  return (
    <div>
      <div className="flex flex-row mt-8 gap-8">
        <WalletCards title="Amout" value={3200} />
        <WalletCards title="Investiments" value={3200} />
        <WalletCards title="Costs" value={3200} />
      </div>

      <div className="flex flex-row gap-8 mt-8">
        <WalletCharts config={chartConfigs.investments} />
        <WalletCharts config={chartConfigs.business} />
        <WalletCharts config={chartConfigs.spending} />
      </div>
    </div>
  );
}
