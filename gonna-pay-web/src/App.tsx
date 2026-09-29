import { Wallet } from "lucide-react"
import { AppCard } from "@/components/card"
import './App.css'

function App() {
  return (
    <div className="flex gap-8 p-8">
      <AppCard
        title="Split Costs"
        description="Divide expenses automatically among your group members."
        icon={Wallet}
        href="#"
        size="default"
      />
      <AppCard
        title="Split Costs"
        description="Divide expenses automatically among your group members."
        icon={Wallet}
        href="#"
        size="sm"
      />
    </div>
  )
}

export default App
