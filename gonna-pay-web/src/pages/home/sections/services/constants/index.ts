import { UserPlus, Users, Receipt, LayoutList } from "lucide-react"

import type { ServiceItem } from "../types"

export const SERVICES: ServiceItem[] = [
  {
    icon: UserPlus,
    title: "Add contacts",
    description: "Register the people you share expenses with — friends, family or coworkers.",
    variant: "default",
  },
  {
    icon: Users,
    title: "Create a group",
    description: "Organize contacts by context: a trip, a dinner, a shared apartment.",
    variant: "primary",
  },
  {
    icon: Receipt,
    title: "Log a cost",
    description: "Register an expense solo or linked to a group. Set how much you cover.",
    variant: "dark",
  },
  {
    icon: LayoutList,
    title: "See splits instantly",
    description: "Gonna Pay calculates and assigns each person's share automatically.",
    variant: "default",
  },
]
