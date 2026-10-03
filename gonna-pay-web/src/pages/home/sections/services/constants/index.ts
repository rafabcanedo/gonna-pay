import type { Variants } from "framer-motion"

import type { ServiceItem } from "../types"

export const cardContainer: Variants = {
  hidden: {},
  visible: { transition: { staggerChildren: 0.18 } }
}

export const cardItem: Variants = {
  hidden: { opacity: 0, y: 24 },
  visible: { opacity: 1, y: 0, transition: { duration: 0.65, ease: "easeOut" } }
}

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
