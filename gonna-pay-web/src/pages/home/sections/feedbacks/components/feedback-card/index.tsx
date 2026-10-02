import type { FeedbackData } from "@/mocks/feedbacks/types"

export function FeedbackCard({ name, role, text }: FeedbackData) {
  return (
    <div className="relative w-72 shrink-0">
      <div className="absolute bottom-[-7px] left-8 w-4 h-4 rotate-45 bg-zinc-900 border-b border-r border-primary z-0" />
      <div className="relative z-10 bg-zinc-900 border border-primary rounded-2xl p-6 flex flex-col gap-4">
        <p className="text-sm text-zinc-300 leading-relaxed">"{text}"</p>
        <div className="flex flex-col gap-0.5">
          <span className="text-sm font-semibold text-primary">{name}</span>
          <span className="text-xs text-zinc-500">{role}</span>
        </div>
      </div>
    </div>
  )
}
