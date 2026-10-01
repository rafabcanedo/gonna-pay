import { ArrowUpRight } from "lucide-react"

import { Card, CardContent } from "@/components/ui/card"
import { cn } from "@/lib/utils"
import type { CardProps } from "./types"

export function AppCard({ title, description, icon: Icon, href, size = "default", className }: CardProps) {
  return (
    <Card
      className={cn(
        size === "sm" && "w-72",
        size === "default" && "w-56",
        className
      )}
    >
      <CardContent className={cn("flex flex-col gap-4", size === "sm" ? "p-5" : "p-6")}>
        <Icon
          className="text-primary"
          size={size === "sm" ? 20 : 24}
        />

        <div className="flex flex-col gap-1">
          <h3 className={cn("font-semibold text-foreground", size === "sm" ? "text-sm" : "text-base")}>
            {title}
          </h3>
          <p className={cn("text-muted-foreground", size === "sm" ? "text-xs" : "text-sm")}>
            {description}
          </p>
        </div>

        <a
          href={href}
          className="group/learn flex items-center gap-2 w-fit"
        >
          <span className="flex items-center justify-center rounded-full bg-foreground p-2 transition-opacity group-hover/learn:opacity-80">
            <ArrowUpRight size={16} className="text-primary" />
          </span>
          <span className="text-sm font-medium text-foreground transition-colors group-hover/learn:text-primary">
            Learn more
          </span>
        </a>
      </CardContent>
    </Card>
  )
}
