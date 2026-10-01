import { cn } from "@/lib/utils"
import { Title } from "@/components/title"

import { SERVICES } from "./constants"
import type { ServiceItem } from "./types"

function ServiceCard({ icon: Icon, title, description, variant }: ServiceItem) {
  const isDark = variant === "dark"
  const isPrimary = variant === "primary"

  return (
    <div className={cn(
      "rounded-2xl p-8 flex flex-col gap-6",
      isPrimary && "bg-primary",
      isDark && "bg-zinc-900",
      !isPrimary && !isDark && "border border-border"
    )}>
      <Icon
        size={24}
        className={cn(isPrimary ? "text-primary-foreground" : "text-primary")}
      />

      <div className="flex flex-col gap-2">
        <span className={cn(
          "text-base font-semibold",
          isPrimary ? "text-primary-foreground" : isDark ? "text-white" : "text-foreground"
        )}>
          {title}
        </span>
        <span className={cn(
          "text-sm leading-relaxed",
          isPrimary ? "text-primary-foreground/70" : isDark ? "text-zinc-400" : "text-muted-foreground"
        )}>
          {description}
        </span>
      </div>
    </div>
  )
}

export function Services() {
  return (
    <section id="services" className="px-8 py-24">
      <Title size="lg" variant="default">What you can do</Title>
      <p className="text-muted-foreground text-base mt-3">Everything you need to split costs the right way.</p>

      <div className="grid grid-cols-2 gap-6 mt-12">
        {SERVICES.map((service: ServiceItem) => (
          <ServiceCard key={service.title} {...service} />
        ))}
      </div>
    </section>
  )
}
