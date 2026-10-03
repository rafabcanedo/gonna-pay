import { Title } from "@/components/title"

import { STEPS } from "./constants"
import { TripMockCard } from "./components/trip-mock-card"
import { StepItem } from "./components/step-item"

export function UseCases() {
  return (
    <section id="how-it-works" className="px-8 py-24">
      <div className="flex flex-col md:flex-row gap-8 md:gap-16 items-center">
        <TripMockCard />

        <div className="flex flex-col">
          <Title size="lg" variant="default">How It Works</Title>
          <p className="text-zinc-400 text-base mt-3">
            From a shared hotel to a dinner bill — log it once, split it instantly.
          </p>
          <div className="flex flex-col gap-8 mt-10">
            {STEPS.map(step => (
              <StepItem key={step.number} {...step} />
            ))}
          </div>
        </div>
      </div>
    </section>
  )
}
