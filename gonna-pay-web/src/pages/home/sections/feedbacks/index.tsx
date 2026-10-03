import { motion } from "framer-motion"

import { Title } from "@/components/title"
import { FEEDBACKS } from "@/mocks/feedbacks"

import { FeedbackCard } from "./components/feedback-card"

export function Feedbacks() {
  return (
    <section id="feedbacks" className="py-24 bg-zinc-950 overflow-hidden">
      <motion.div
        className="px-8 mb-12"
        initial={{ opacity: 0, y: 16 }}
        whileInView={{ opacity: 1, y: 0 }}
        viewport={{ once: true }}
        transition={{ duration: 0.5, ease: "easeOut" }}
      >
        <Title size="lg" variant="default">Feedbacks</Title>
        <p className="text-zinc-400 text-base mt-3">What people are saying about Gonna Pay</p>
      </motion.div>

      <div
        className="overflow-hidden"
        style={{ maskImage: "linear-gradient(to right, transparent, black 8%, black 92%, transparent)" }}
      >
        <div className="flex gap-8 animate-marquee w-max hover:[animation-play-state:paused]">
          {[...FEEDBACKS, ...FEEDBACKS].map((feedback, i) => (
            <FeedbackCard key={i} {...feedback} />
          ))}
        </div>
      </div>
    </section>
  )
}
