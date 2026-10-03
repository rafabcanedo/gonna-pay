import { motion } from "framer-motion"

import { Title } from "@/components/title"

import { STEPS, slideFromLeft, slideFromRight, stepsContainer, stepItem } from "./constants"
import { TripMockCard } from "./components/trip-mock-card"
import { StepItem } from "./components/step-item"

export function UseCases() {
  return (
    <section id="how-it-works" className="px-8 py-24">
      <motion.div
        className="flex flex-col md:flex-row gap-8 md:gap-16 items-center max-w-4xl mx-auto"
        initial="hidden"
        whileInView="visible"
        viewport={{ once: true, margin: "-80px" }}
      >
        <motion.div variants={slideFromLeft}>
          <TripMockCard />
        </motion.div>

        <motion.div className="flex flex-col" variants={slideFromRight}>
          <Title size="lg" variant="default">How It Works</Title>
          <p className="text-zinc-400 text-base mt-3">
            From a shared hotel to a dinner bill — log it once, split it instantly.
          </p>
          <motion.div className="flex flex-col gap-8 mt-10" variants={stepsContainer}>
            {STEPS.map(step => (
              <motion.div key={step.number} variants={stepItem}>
                <StepItem {...step} />
              </motion.div>
            ))}
          </motion.div>
        </motion.div>
      </motion.div>
    </section>
  )
}
