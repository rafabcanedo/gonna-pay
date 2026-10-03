import type { Variants } from "framer-motion"

export const fadeUpContainer: Variants = {
  hidden: {},
  visible: { transition: { staggerChildren: 0.12 } }
}

export const fadeUpItem: Variants = {
  hidden: { opacity: 0, y: 20 },
  visible: { opacity: 1, y: 0, transition: { duration: 0.5, ease: "easeOut" } }
}
