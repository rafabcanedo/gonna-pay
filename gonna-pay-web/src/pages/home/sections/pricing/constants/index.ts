import type { Variants } from "framer-motion"

export const freePlan: Variants = {
  hidden: { opacity: 0, x: -30 },
  visible: { opacity: 1, x: 0, transition: { duration: 0.6, ease: "easeOut" } }
}

export const proPlan: Variants = {
  hidden: { opacity: 0, x: 30, scale: 0.97 },
  visible: { opacity: 1, x: 0, scale: 1, transition: { duration: 0.65, ease: "easeOut", delay: 0.1 } }
}
