import { Title } from "@/components/title"

import { CardPlan } from "./components/card-plan"

export function HomePricing() {
  return (
    <section id="pricing" className="px-8 py-24">
      <Title size="lg" variant="default">Pricing</Title>

      <p className="text-base text-muted-foreground">Simple plans for everyone</p>

      <div className="mt-12 flex flex-col md:flex-row gap-8 justify-center items-center">
        <CardPlan
          planName="Free"
          description="Everything you need to start splitting expenses — create groups, track costs, and settle up with friends and family at no cost."
        />

        <CardPlan
          planName="Pro"
          description="Unlock the full Gonna Pay experience with advanced analytics, unlimited groups, priority support, and more."
          highlighted
        />
      </div>
    </section>
  )
}
