import { ArrowUpRight } from "lucide-react"
import { Link } from "react-router"

import { Card, CardContent } from "@/components/ui/card"
import { cn } from "@/lib/utils"

interface Props {
  planName: string
  description: string
  highlighted?: boolean
  className?: string
}

export function CardPlan({ planName, description, highlighted = false, className }: Props) {
  return (
    <Card className={cn("flex flex-col w-80", highlighted && "ring-2 ring-primary", className)}>
      <CardContent className="flex flex-col gap-6 p-6 h-full">
        <div className="flex flex-col gap-3">
          <span className="text-2xl font-semibold text-foreground">{planName}</span>
          <p className="text-sm text-muted-foreground">{description}</p>
        </div>

        <Link to="/pricing" className="group/learn flex items-center gap-2 w-fit">
          <span className="flex items-center justify-center rounded-full bg-foreground p-2 transition-opacity group-hover/learn:opacity-80">
            <ArrowUpRight size={16} className="text-primary" />
          </span>
          <span className="text-sm font-medium text-foreground transition-colors group-hover/learn:text-primary">
            Learn more
          </span>
        </Link>
      </CardContent>
    </Card>
  )
}
