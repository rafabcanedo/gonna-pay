import { Check, X } from "lucide-react"
import { cn } from "@/lib/utils"
import { Button } from "@/components/ui/button"
import { Card, CardContent } from "@/components/ui/card"
import type { CardPricingProps } from "./types"

export function CardPricing({
  planName,
  price,
  priceNote,
  description,
  features,
  buttonLabel,
  highlighted = false,
  onButtonClick,
  className,
}: CardPricingProps) {
  return (
    <Card className={cn("flex flex-col w-full md:w-72", highlighted && "ring-2 ring-primary", className)}>
      <CardContent className="flex flex-col gap-6 p-6 h-full">
        <div className="flex flex-col gap-1">
          <span className="text-sm text-muted-foreground">{planName}</span>

          <div className="flex items-baseline gap-1">
            <span className="text-4xl font-bold text-foreground">{price}</span>
            {priceNote && (
              <span className="text-sm text-muted-foreground">{priceNote}</span>
            )}
          </div>

          <p className="text-sm text-muted-foreground">{description}</p>
        </div>

        <ul className="flex flex-col gap-3 flex-1">
          {features.map((feature, index) => (
            <li key={index} className="flex items-center gap-3">
              {feature.included ? (
                <Check size={16} className="text-primary shrink-0" />
              ) : (
                <X size={16} className="text-muted-foreground shrink-0" />
              )}
              <span className={cn("text-sm", !feature.included && "text-muted-foreground")}>
                {feature.label}
              </span>
            </li>
          ))}
        </ul>

        <Button
          variant={highlighted ? "default" : "outline"}
          className="w-full"
          onClick={onButtonClick}
        >
          {buttonLabel}
        </Button>
      </CardContent>
    </Card>
  )
}
