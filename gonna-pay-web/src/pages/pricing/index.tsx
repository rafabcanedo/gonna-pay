import { CardPricing } from "@/components/card-pricing"
import { Title } from "@/components/title"
import { FREE_FEATURES, PRO_FEATURES } from "./constants"
import { PricingIntro } from "./components/pricing-intro"

export function Pricing() {
  return (
    <div className="flex flex-col items-center gap-12 py-24 px-8">

      <PricingIntro />

      <div className="w-full flex justify-start mt-12">
        <Title size="lg" variant="default">Our Pricing</Title>
      </div>

      <div className="flex flex-row gap-8">
        <CardPricing
          planName="Free"
          price="$0"
          description="Get started splitting expenses with your group."
          features={FREE_FEATURES}
          buttonLabel="Get started"
        />

        <CardPricing
          planName="Pro"
          price="$6"
          priceNote="/month"
          description="For those who want the full experience."
          features={PRO_FEATURES}
          buttonLabel="Get Pro"
          highlighted
        />
      </div>
    </div>
  )
}
