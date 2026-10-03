import { motion } from "framer-motion"

import { Title } from "@/components/title"
import { ContactForm } from "./components/contact-form"
import { ContactArt } from "./components/contact-art"

export function ContactUs() {
  return (
    <section id="contact" className="px-8 py-24">
      <motion.div
        initial={{ opacity: 0, y: 16 }}
        whileInView={{ opacity: 1, y: 0 }}
        viewport={{ once: true }}
        transition={{ duration: 0.5, ease: "easeOut" }}
      >
        <Title size="lg" variant="default">Contact Us</Title>

        <div>
          <h2 className="text-xl text-zinc-700">Connect with us</h2>
          <span className="text-base text-zinc-700">Let's to talk about your costs</span>
        </div>
      </motion.div>

      <div className="mt-12 flex flex-col md:flex-row gap-12 rounded-2xl bg-card p-6 md:p-10 shadow-xl ring-1 ring-foreground/5 max-w-4xl mx-auto">
        <ContactForm />
        <ContactArt />
      </div>
    </section>
  )
}
