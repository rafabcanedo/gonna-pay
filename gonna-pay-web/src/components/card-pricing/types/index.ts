export interface PricingFeature {
  label: string
  included: boolean
}

export interface CardPricingProps {
  planName: string
  price: string
  priceNote?: string
  description: string
  features: PricingFeature[]
  buttonLabel: string
  highlighted?: boolean
  onButtonClick?: () => void
  className?: string
}
