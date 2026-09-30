import type { PricingFeature } from "@/components/card-pricing/types"

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
