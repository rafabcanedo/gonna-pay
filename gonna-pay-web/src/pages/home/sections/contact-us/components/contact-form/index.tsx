import { motion } from "framer-motion"

import { Field, FieldGroup, FieldTitle } from "@/components/ui/field"
import { Input } from "@/components/ui/input"
import { Button } from "@/components/ui/button"

export function ContactForm() {
  return (
    <motion.form
      className="flex flex-1 flex-col gap-8"
      initial={{ opacity: 0, y: 16 }}
      whileInView={{ opacity: 1, y: 0 }}
      viewport={{ once: true }}
      transition={{ duration: 0.5, ease: "easeOut", delay: 0.15 }}
    >
      <FieldGroup>
        <Field>
          <FieldTitle>Name</FieldTitle>
          <Input placeholder="Your name" />
        </Field>

        <Field>
          <FieldTitle>Email</FieldTitle>
          <Input type="email" placeholder="your@email.com" />
        </Field>

        <Field>
          <FieldTitle>Message</FieldTitle>
          <textarea
            rows={5}
            placeholder="Your message..."
            className="w-full rounded-lg border border-input bg-transparent px-2.5 py-2 text-sm outline-none transition-colors placeholder:text-muted-foreground focus-visible:border-ring focus-visible:ring-3 focus-visible:ring-ring/50 resize-none"
          />
        </Field>
      </FieldGroup>

      <Button variant="default" className="w-full h-10">
        Send Message
      </Button>
    </motion.form>
  )
}
