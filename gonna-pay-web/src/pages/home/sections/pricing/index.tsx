import { motion } from "framer-motion"

import { Title } from "@/components/title"

import { freePlan, proPlan } from "./constants"
import { CardPlan } from "./components/card-plan"

export function HomePricing() {
  return (
    <section id="pricing" className="px-8 py-24">
      <motion.div
        initial={{ opacity: 0, y: 16 }}
        whileInView={{ opacity: 1, y: 0 }}
        viewport={{ once: true }}
        transition={{ duration: 0.5, ease: "easeOut" }}
      >
        <Title size="lg" variant="default">Pricing</Title>
        <p className="text-base text-muted-foreground">Simple plans for everyone</p>
      </motion.div>

      <motion.div
        className="mt-12 flex flex-col md:flex-row gap-8 justify-center items-center"
        initial="hidden"
        whileInView="visible"
        viewport={{ once: true }}
      >
        <motion.div variants={freePlan}>
          <CardPlan
            planName="Free"
            description="Everything you need to start splitting expenses — create groups, track costs, and settle up with friends and family at no cost."
          />
        </motion.div>

        <motion.div variants={proPlan}>
          <CardPlan
            planName="Pro"
            description="Unlock the full Gonna Pay experience with advanced analytics, unlimited groups, priority support, and more."
            highlighted
          />
        </motion.div>
      </motion.div>
    </section>
  )
}
