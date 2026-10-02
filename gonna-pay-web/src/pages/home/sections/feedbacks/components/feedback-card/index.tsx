import type { FeedbackData } from "@/mocks/feedbacks/types"

export function FeedbackCard({ name, role, text }: FeedbackData) {
  return (
    <div className="w-72 shrink-0 bg-zinc-900 border border-primary rounded-2xl p-6 flex flex-col gap-4">
      <p className="text-sm text-zinc-300 leading-relaxed">"{text}"</p>
      <div className="flex flex-row gap-3 items-center">
        <div className="w-8 h-8 rounded-full bg-primary flex items-center justify-center shrink-0">
          <span className="text-xs font-bold text-primary-foreground">{name[0]}</span>
        </div>
        <div className="flex flex-col gap-0.5">
          <span className="text-sm font-semibold text-primary">{name}</span>
          <span className="text-xs text-zinc-500">{role}</span>
        </div>
      </div>
    </div>
  )
}
