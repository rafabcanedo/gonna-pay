import { Plane } from "lucide-react"

import { cn } from "@/lib/utils"
import { TRIP_NAME, TRIP_OWNER_PERCENTAGE, TRIP_COSTS, TRIP_SPLITS } from "@/mocks/trip"
import type { TripCost, TripSplit } from "@/mocks/trip/types"

export function TripMockCard() {
  return (
    <div className="bg-zinc-900 border border-zinc-800 rounded-3xl p-8 flex flex-col gap-6 shadow-[0_0_60px_rgba(161,237,111,0.08)] shrink-0">
      <div className="flex flex-col gap-2">
        <div className="flex items-center gap-2">
          <Plane className="w-4 h-4 text-primary" />
          <span className="text-white font-semibold">{TRIP_NAME}</span>
        </div>
        <div className="flex items-center gap-2 text-xs text-zinc-500">
          {TRIP_SPLITS.map((split, i) => (
            <>
              <span key={split.name}>{split.name}</span>
              {i < TRIP_SPLITS.length - 1 && <span key={`dot-${i}`}>·</span>}
            </>
          ))}
        </div>
      </div>

      <div className="border-t border-zinc-800" />

      <div className="flex flex-col gap-3">
        {TRIP_COSTS.map((cost: TripCost) => {
          const Icon = cost.icon
          return (
            <div key={cost.name} className="flex items-center justify-between">
              <div className="flex items-center gap-3">
                <div className="w-8 h-8 bg-zinc-800 rounded-lg flex items-center justify-center shrink-0">
                  <Icon className="w-4 h-4 text-zinc-400" />
                </div>
                <span className="text-sm text-zinc-300">{cost.name}</span>
              </div>
              <span className="text-sm font-medium text-white">${cost.total}</span>
            </div>
          )
        })}
      </div>

      <div className="border-t border-zinc-800" />

      <div className="flex flex-col gap-3">
        <span className="text-xs font-semibold text-zinc-500 uppercase tracking-wider">Splits</span>
        {TRIP_SPLITS.map((split: TripSplit) => {
          const percentage = split.isOwner
            ? TRIP_OWNER_PERCENTAGE
            : (100 - TRIP_OWNER_PERCENTAGE) / 2

          return (
            <div key={split.name} className="flex items-center justify-between">
              <span className={cn("text-sm", split.isOwner ? "text-primary font-medium" : "text-zinc-300")}>
                {split.name}
              </span>
              <div className="flex items-center gap-3">
                <span className="text-xs text-zinc-500">{percentage}%</span>
                <span className={cn("text-sm font-semibold w-12 text-right", split.isOwner ? "text-primary" : "text-white")}>
                  ${split.owes}
                </span>
              </div>
            </div>
          )
        })}
      </div>
    </div>
  )
}
