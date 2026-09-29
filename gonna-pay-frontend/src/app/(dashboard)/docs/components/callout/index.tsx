import { Lightbulb, TriangleAlert } from 'lucide-react'
import type { IPropsCallout } from './interfaces'

const config = {
  tip: {
    Icon: Lightbulb,
    wrapper: 'border-lime-400 bg-lime-50',
    iconColor: 'text-lime-600',
    label: 'Tip',
    labelColor: 'text-lime-700',
  },
  important: {
    Icon: TriangleAlert,
    wrapper: 'border-amber-400 bg-amber-50',
    iconColor: 'text-amber-500',
    label: 'Important',
    labelColor: 'text-amber-700',
  },
}

export function Callout({ type, children }: IPropsCallout) {
  const { Icon, wrapper, iconColor, label, labelColor } = config[type]

  return (
    <div className={`not-prose flex gap-3 rounded-lg border-l-4 px-4 py-3 my-4 ${wrapper}`}>
      <Icon className={`mt-0.5 h-4 w-4 shrink-0 ${iconColor}`} />
      <div className="flex flex-col gap-1">
        <span className={`text-sm font-semibold ${labelColor}`}>{label}</span>
        <div className="text-sm text-zinc-700">{children}</div>
      </div>
    </div>
  )
}
