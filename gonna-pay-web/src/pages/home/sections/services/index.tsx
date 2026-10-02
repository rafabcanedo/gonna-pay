import { Title } from "@/components/title"

import { SERVICES } from "./constants"
import { ServiceCard } from "./components/service-card"

export function Services() {
  return (
    <section id="services" className="px-8 py-24 bg-zinc-950">
      <Title size="lg" variant="default">What you can do</Title>
      <p className="text-zinc-400 text-base mt-3">Everything you need to split costs the right way.</p>

      <div className="grid grid-cols-2 gap-6 mt-12 max-w-4xl mx-auto">
        {SERVICES.map((service) => (
          <ServiceCard key={service.artKey} {...service} />
        ))}
      </div>
    </section>
  )
}
