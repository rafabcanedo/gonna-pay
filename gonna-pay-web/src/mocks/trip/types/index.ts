import type { LucideIcon } from "lucide-react"

export interface TripCost {
  icon: LucideIcon
  name: string
  total: number
}

export interface TripSplit {
  name: string
  owes: number
  isOwner: boolean
}
