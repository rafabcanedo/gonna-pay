import type { ServiceItem } from "../types"

export const SERVICES: ServiceItem[] = [
  {
    artKey: "tracking",
    title: "Expense tracking",
    description: "Track personal or shared costs, organized by category and period.",
    variant: "default",
  },
  {
    artKey: "group",
    title: "Group splitting",
    description: "Create groups for any context: a trip, a dinner, a shared apartment. Split costs among all members.",
    variant: "primary",
  },
  {
    artKey: "math",
    title: "Automatic math",
    description: "Set your share. Gonna Pay divides the rest equally among everyone in the group.",
    variant: "dark",
  },
  {
    artKey: "contacts",
    title: "Contact management",
    description: "Keep your people organized by family, friends, and coworkers. See who you split with most.",
    variant: "default",
  },
]
