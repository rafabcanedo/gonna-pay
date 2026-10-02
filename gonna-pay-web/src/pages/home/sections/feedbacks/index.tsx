import { Title } from "@/components/title"
import { FEEDBACKS } from "@/mocks/feedbacks"

import { FeedbackCard } from "./components/feedback-card"

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
