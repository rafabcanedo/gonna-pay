import type { LucideIcon } from "lucide-react"

export interface ServiceItem {
  icon: LucideIcon
  title: string
  description: string
  variant: "default" | "primary" | "dark"
}
