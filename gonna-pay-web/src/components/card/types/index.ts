import type { LucideIcon } from "lucide-react"

export interface CardProps {
  title: string
  description: string
  icon: LucideIcon
  href: string
  size?: "default" | "sm"
  className?: string
}
