import type { ComponentType } from "react"

import { cn } from "@/lib/utils"

import type { ServiceItem, ServiceArtKey } from "../../types"
import { TrackingArt } from "../tracking-art"
import { GroupArt } from "../group-art"
import { MathArt } from "../math-art"
import { ContactsArt } from "../contacts-art"

const artMap: Record<ServiceArtKey, ComponentType> = {
  tracking: TrackingArt,
  group: GroupArt,
  math: MathArt,
  contacts: ContactsArt,
}

export function ServiceCard({ artKey, title, description, variant }: ServiceItem) {
  const isDark = variant === "dark"
  const isPrimary = variant === "primary"
  const Art = artMap[artKey]

  return (
    <div className={cn(
      "rounded-2xl p-8 flex flex-row h-52",
      isPrimary && "bg-primary",
      isDark && "bg-zinc-900",
      !isPrimary && !isDark && "border border-border"
    )}>
      <div className="flex flex-col gap-3 flex-1 min-w-0">
        <span className={cn(
          "text-sm font-semibold w-fit px-2.5 py-1.5 rounded-md leading-tight",
          isDark ? "bg-primary text-primary-foreground" : "bg-zinc-900 text-white"
        )}>
          {title}
        </span>

        <p className={cn(
          "text-sm leading-relaxed max-w-[180px]",
          isPrimary ? "text-primary-foreground/80" : isDark ? "text-zinc-400" : "text-muted-foreground"
        )}>
          {description}
        </p>
      </div>

      <div className="flex items-center justify-end shrink-0">
        <Art />
      </div>
    </div>
  )
}
