import workInProgress from "@/assets/work-in-progress.png"
import { Button } from "@/components/ui/button"
import { useNavigate } from "react-router"

export function WarningDevelopment() {
  const navigate = useNavigate()

  return (
    <div className="flex flex-col items-center justify-center gap-8 py-12">
      <div>
        <img src={workInProgress} alt="Work in progress" className="w-100 h-100" />
      </div>

      <div className="flex flex-row items-center justify-center gap-4">
        <span className="text-xl">🚧</span>
        <p className="text-lg font-semibold text-foreground">This page is work in progress</p>
        <span className="text-xl">🚧</span>
      </div>

      <div>
        <Button variant="ghost" size="xs" onClick={() => navigate("/")}>
          Back to home
        </Button>
      </div>
    </div>
  )
}
