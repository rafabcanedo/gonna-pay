import type { Variants } from "framer-motion"

import type { PricingFeature } from "@/components/card-pricing/types"

export const FREE_PLAN: Variants = {
  hidden: { opacity: 0, x: -30 },
  visible: { opacity: 1, x: 0, transition: { duration: 0.6, ease: "easeOut" } },
}

export const PRO_PLAN: Variants = {
  hidden: { opacity: 0, x: 30, scale: 0.97 },
  visible: { opacity: 1, x: 0, scale: 1, transition: { duration: 0.65, ease: "easeOut", delay: 0.1 } },
}

export const FREE_FEATURES: PricingFeature[] = [
  { label: "Basic expense splitting", included: true },
  { label: "Up to 5 contacts", included: true },
  { label: "Dashboard overview", included: true },
  { label: "Advanced reports", included: false },
  { label: "Priority support", included: false },
]

export const PRO_FEATURES: PricingFeature[] = [
  { label: "Basic expense splitting", included: true },
  { label: "Unlimited contacts", included: true },
  { label: "Dashboard overview", included: true },
  { label: "Advanced reports", included: true },
  { label: "Priority support", included: true },
]
