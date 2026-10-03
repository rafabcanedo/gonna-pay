import { motion } from "framer-motion"

import { Title } from "@/components/title"

import { SERVICES, cardContainer, cardItem } from "./constants"
import { ServiceCard } from "./components/service-card"

export function Services() {
  return (
    <section id="features" className="px-8 py-24 bg-zinc-950">
      <motion.div
        initial={{ opacity: 0, y: 16 }}
        whileInView={{ opacity: 1, y: 0 }}
        viewport={{ once: true }}
        transition={{ duration: 0.5, ease: "easeOut" }}
      >
        <Title size="lg" variant="default">Features</Title>
        <p className="text-zinc-400 text-base mt-3">Everything you need to split costs the right way.</p>
      </motion.div>

      <motion.div
        className="grid grid-cols-1 md:grid-cols-2 gap-6 mt-12 max-w-4xl mx-auto"
        variants={cardContainer}
        initial="hidden"
        whileInView="visible"
        viewport={{ once: true, margin: "-80px" }}
      >
        {SERVICES.map((service) => (
          <motion.div key={service.artKey} variants={cardItem}>
            <ServiceCard {...service} />
          </motion.div>
        ))}
      </motion.div>
    </section>
  )
}
