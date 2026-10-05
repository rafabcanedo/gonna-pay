import { motion } from "framer-motion"

import { Button } from "@/components/ui/button"
import { fadeUpContainer, fadeUpItem } from "@/lib/animations"

import { HeroArt } from "./components/hero-art"

export function Intro() {
  return (
    <section id="intro" className="px-8 py-24 flex flex-col items-center text-center">
      <motion.div
        className="flex flex-col items-center text-center w-full"
        variants={fadeUpContainer}
        initial="hidden"
        animate="visible"
      >
        <motion.div className="flex flex-col" variants={fadeUpItem}>
          <h1>You enjoy the moment</h1>
          <h1>We track the bill</h1>
        </motion.div>

        <motion.p className="text-muted-foreground text-base mt-4 max-w-md" variants={fadeUpItem}>
          Create a group, add your expenses and set each person's share.
          Gonna Pay figures out the rest.
        </motion.p>

        <motion.div variants={fadeUpItem}>
          <Button variant="default" size="default" className="md:flex w-40 h-12 border-2 border-primary bg-primary hover:border-lime-500 mt-8">
            Get Started
          </Button>
        </motion.div>

        <motion.div className="mt-16 w-full max-w-2xl mx-auto" variants={fadeUpItem}>
          <HeroArt />
        </motion.div>
      </motion.div>
    </section>
  )
}
