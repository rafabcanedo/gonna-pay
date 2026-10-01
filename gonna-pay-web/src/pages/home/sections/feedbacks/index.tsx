import { Title } from "@/components/title"
import { FEEDBACKS } from "./constants"
import type { FeedbackData } from "./types"

function FeedbackCard({ name, role, text }: FeedbackData) {
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

export function Feedbacks() {
  return (
    <section id="feedbacks" className="py-24 bg-zinc-950 overflow-hidden">
      <div className="px-8 mb-12">
        <Title size="lg" variant="default">Feedbacks</Title>
        <p className="text-zinc-400 text-base mt-3">What people are saying about Gonna Pay</p>
      </div>

      <div className="flex gap-8 animate-marquee w-max">
        {[...FEEDBACKS, ...FEEDBACKS].map((feedback, i) => (
          <FeedbackCard key={i} {...feedback} />
        ))}
      </div>
    </section>
  )
}
