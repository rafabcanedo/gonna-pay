import type { Step } from "../../types"

export function StepItem({ number, title, description }: Step) {
  return (
    <div className="flex gap-4">
      <span className="text-primary font-mono text-sm font-medium shrink-0 pt-0.5">{number}</span>
      <div className="flex flex-col gap-1">
        <span className="text-foreground text-sm font-semibold">{title}</span>
        <span className="text-zinc-400 text-sm leading-relaxed">{description}</span>
      </div>
    </div>
  )
}
