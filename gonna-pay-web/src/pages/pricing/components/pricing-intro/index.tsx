import { Check, Clock } from "lucide-react"
import { cn } from "@/lib/utils"
import { MOCK_CARDS } from "./constants"
import type { MockCardData } from "./types"

function MockCard({ card }: { card: MockCardData }) {
  return (
    <div className="bg-zinc-900 rounded-2xl p-5 w-52 flex flex-col gap-3 shadow-xl">
      <div className="flex flex-col gap-0.5">
        <span className="text-zinc-500 text-xs">Last expense</span>
        <div className="flex items-center gap-2">
          <span>{card.emoji}</span>
          <span className="text-white font-semibold text-sm">{card.title}</span>
        </div>
        <span className="text-2xl font-bold text-white">{card.total}</span>
      </div>

      <div className="h-px bg-zinc-800" />

      <div className="flex flex-col gap-2">
        {card.splits.map((split) => (
          <div key={split.name} className="flex items-center justify-between">
            <div className="flex items-center gap-2">
              {split.paid ? (
                <Check size={12} className="text-primary shrink-0" />
              ) : (
                <Clock size={12} className="text-zinc-500 shrink-0" />
              )}
              <span className={cn("text-xs", split.paid ? "text-white" : "text-zinc-500")}>
                {split.name}
              </span>
            </div>
            <span className={cn("text-xs font-medium", split.paid ? "text-white" : "text-zinc-500")}>
              {split.amount}
            </span>
          </div>
        ))}
      </div>
    </div>
  )
}

export function PricingIntro() {
  return (
    <div className="flex flex-col md:flex-row items-center gap-8 md:gap-16 w-full">
      <div className="flex flex-col gap-5 flex-1">
        <span className="text-sm font-medium text-primary">Why Gonna Pay?</span>

        <div className="flex flex-col gap-1">
          <h2 className="text-2xl md:text-4xl font-bold text-foreground leading-tight">
            A smarter way to manage your shared costs
          </h2>
          <span className="text-4xl font-bold text-primary">day by day.</span>
        </div>

        <p className="text-muted-foreground text-base max-w-sm">
          Split expenses, track who owes what, and keep things fair — automatically.
        </p>
      </div>

      <div className="flex-1 flex justify-center">
        <div className="flex items-end">
          <div className="relative -rotate-6 translate-y-4 -mr-5 hidden md:block" style={{ zIndex: 1 }}>
            <MockCard card={MOCK_CARDS[0]} />
          </div>

          <div className="relative" style={{ zIndex: 3 }}>
            <MockCard card={MOCK_CARDS[1]} />
          </div>

          <div className="relative rotate-6 translate-y-4 -ml-5 hidden md:block" style={{ zIndex: 2 }}>
            <MockCard card={MOCK_CARDS[2]} />
          </div>
        </div>
      </div>
    </div>
  )
}
