import { motion } from "framer-motion"

export function ContactArt() {
  return (
    <div className="flex items-center justify-center w-full md:flex-1">
      <svg viewBox="0 0 360 320" width="100%" height="auto" fill="none" xmlns="http://www.w3.org/2000/svg">

        <motion.line
          x1="180" y1="160" x2="235" y2="65"
          stroke="#d4d4d8" strokeWidth="1" opacity="0.6"
          strokeDasharray="110"
          initial={{ strokeDashoffset: 110 }}
          whileInView={{ strokeDashoffset: 0 }}
          viewport={{ once: true }}
          transition={{ duration: 1, ease: "easeInOut", delay: 0 }}
        />
        <motion.line
          x1="180" y1="160" x2="199" y2="268"
          stroke="#d4d4d8" strokeWidth="1" opacity="0.6"
          strokeDasharray="110"
          initial={{ strokeDashoffset: 110 }}
          whileInView={{ strokeDashoffset: 0 }}
          viewport={{ once: true }}
          transition={{ duration: 1, ease: "easeInOut", delay: 0.2 }}
        />
        <motion.line
          x1="180" y1="160" x2="80" y2="207"
          stroke="#d4d4d8" strokeWidth="1" opacity="0.6"
          strokeDasharray="110"
          initial={{ strokeDashoffset: 110 }}
          whileInView={{ strokeDashoffset: 0 }}
          viewport={{ once: true }}
          transition={{ duration: 1, ease: "easeInOut", delay: 0.4 }}
        />

        <circle cx="180" cy="160" r="110" stroke="#e4e4e7" strokeWidth="1.5" />

        <motion.g
          style={{ transformOrigin: "180px 160px" }}
          animate={{ rotate: 360 }}
          transition={{ duration: 8, ease: "linear", repeat: Infinity }}
        >
          <polygon
            points="180,95 193,147 245,160 193,173 180,225 167,173 115,160 167,147"
            fill="#18181b"
          />
        </motion.g>

        <motion.circle
          cx="235" cy="65" r="9" fill="#a1ed6f"
          style={{ transformOrigin: "235px 65px" }}
          animate={{ scale: [1, 1.3, 1] }}
          transition={{ duration: 2, ease: "easeInOut", repeat: Infinity, delay: 0 }}
        />
        <motion.circle
          cx="199" cy="268" r="6" fill="#a1ed6f"
          style={{ transformOrigin: "199px 268px" }}
          animate={{ scale: [1, 1.3, 1] }}
          transition={{ duration: 2, ease: "easeInOut", repeat: Infinity, delay: 0.6 }}
        />
        <motion.circle
          cx="80" cy="207" r="5" fill="#a1ed6f"
          style={{ transformOrigin: "80px 207px" }}
          animate={{ scale: [1, 1.3, 1] }}
          transition={{ duration: 2, ease: "easeInOut", repeat: Infinity, delay: 1.2 }}
        />

        <polygon
          points="295,232 299,245 311,248 299,252 295,264 291,252 279,248 291,244"
          fill="#a1ed6f"
        />

      </svg>
    </div>
  )
}
