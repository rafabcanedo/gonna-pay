import { motion } from "framer-motion"

import { CardPricing } from "@/components/card-pricing"
import { Title } from "@/components/title"

import { FREE_FEATURES, PRO_FEATURES, FREE_PLAN, PRO_PLAN } from "./constants"
import { PricingIntro } from "./components/pricing-intro"

export function Pricing() {
  return (
    <div className="flex flex-col items-center gap-12 py-24 px-8">

      <PricingIntro />

      <motion.div
        className="w-full flex justify-start mt-12"
        initial={{ opacity: 0, y: 16 }}
        whileInView={{ opacity: 1, y: 0 }}
        viewport={{ once: true }}
        transition={{ duration: 0.5, ease: "easeOut" }}
      >
        <Title size="lg" variant="default">Our Pricing</Title>
      </motion.div>

      <motion.div
        className="flex flex-col md:flex-row gap-8"
        initial="hidden"
        whileInView="visible"
        viewport={{ once: true }}
      >
        <motion.div variants={FREE_PLAN}>
          <CardPricing
            planName="Free"
            price="$0"
            description="Get started splitting expenses with your group."
            features={FREE_FEATURES}
            buttonLabel="Get started"
          />
        </motion.div>

        <motion.div variants={PRO_PLAN}>
          <CardPricing
            planName="Pro"
            price="$6"
            priceNote="/month"
            description="For those who want the full experience."
            features={PRO_FEATURES}
            buttonLabel="Get Pro"
            highlighted
          />
        </motion.div>
      </motion.div>
    </div>
  )
}
