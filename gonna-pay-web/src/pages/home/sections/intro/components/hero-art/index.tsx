import { motion } from "framer-motion"

export function HeroArt() {
  return (
    <motion.div
      animate={{ y: [-6, 6, -6] }}
      transition={{ duration: 3, ease: "easeInOut", repeat: Infinity }}
    >
    <svg
      viewBox="0 0 500 560"
      width="100%"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
      style={{ fontFamily: "'Geist Variable', sans-serif" }}
    >
      <defs>
        <clipPath id="phone-screen">
          <rect x="128" y="24" width="244" height="492" rx="26" />
        </clipPath>
      </defs>

      {/* Phone bezel */}
      <rect x="120" y="10" width="260" height="520" rx="36" fill="#09090b" />

      {/* Screen background */}
      <rect x="128" y="24" width="244" height="492" rx="26" fill="#09090b" />

      {/* Notch */}
      <rect x="215" y="19" width="70" height="10" rx="5" fill="#09090b" />

      <g clipPath="url(#phone-screen)">
        {/* App header */}
        <rect x="128" y="24" width="244" height="52" fill="#111111" />
        <circle cx="148" cy="50" r="8" fill="#a1ed6f" />
        <text x="162" y="55" fill="white" fontSize="13" fontWeight="600">Gonna Pay</text>
        <circle cx="356" cy="50" r="11" fill="#27272a" />
        <circle cx="356" cy="46" r="4" fill="#3f3f46" />
        <path d="M347 60 Q356 56 365 60" stroke="#3f3f46" strokeWidth="1.5" strokeLinecap="round" />
        <line x1="128" y1="76" x2="372" y2="76" stroke="#1c1c1c" strokeWidth="1" />

        {/* Card 1: Weekend Trip */}
        <rect x="136" y="84" width="228" height="162" rx="12" fill="#18181b" />
        <text x="150" y="104" fill="#52525b" fontSize="10">Last expense</text>
        <rect x="150" y="114" width="18" height="18" rx="5" fill="#292524" />
        <text x="174" y="127" fill="white" fontSize="12" fontWeight="500">Weekend Trip</text>
        <text x="150" y="163" fill="white" fontSize="22" fontWeight="700">$280.00</text>
        <line x1="150" y1="173" x2="350" y2="173" stroke="#27272a" strokeWidth="1" />
        <circle cx="160" cy="188" r="4" fill="#a1ed6f" />
        <text x="170" y="192" fill="white" fontSize="11">You</text>
        <text x="350" y="192" fill="white" fontSize="11" textAnchor="end">$93.00</text>
        <circle cx="160" cy="207" r="4" fill="#a1ed6f" />
        <text x="170" y="211" fill="white" fontSize="11">Carlos</text>
        <text x="350" y="211" fill="white" fontSize="11" textAnchor="end">$93.00</text>
        <circle cx="160" cy="226" r="4" fill="none" stroke="#3f3f46" strokeWidth="1.5" />
        <text x="170" y="230" fill="#52525b" fontSize="11">Maria</text>
        <text x="350" y="230" fill="#52525b" fontSize="11" textAnchor="end">$94.00</text>

        {/* Card 2: Dinner with Friends */}
        <rect x="136" y="254" width="228" height="162" rx="12" fill="#18181b" />
        <text x="150" y="274" fill="#52525b" fontSize="10">Last expense</text>
        <rect x="150" y="284" width="18" height="18" rx="5" fill="#431407" />
        <text x="174" y="297" fill="white" fontSize="12" fontWeight="500">Dinner with Friends</text>
        <text x="150" y="333" fill="white" fontSize="22" fontWeight="700">$45.00</text>
        <line x1="150" y1="343" x2="350" y2="343" stroke="#27272a" strokeWidth="1" />
        <circle cx="160" cy="358" r="4" fill="#a1ed6f" />
        <text x="170" y="362" fill="white" fontSize="11">You</text>
        <text x="350" y="362" fill="white" fontSize="11" textAnchor="end">$20.00</text>
        <circle cx="160" cy="377" r="4" fill="#a1ed6f" />
        <text x="170" y="381" fill="white" fontSize="11">Ana</text>
        <text x="350" y="381" fill="white" fontSize="11" textAnchor="end">$10.00</text>
        <circle cx="160" cy="396" r="4" fill="none" stroke="#3f3f46" strokeWidth="1.5" />
        <text x="170" y="400" fill="#52525b" fontSize="11">João</text>
        <text x="350" y="400" fill="#52525b" fontSize="11" textAnchor="end">$15.00</text>

        {/* Card 3: Movie Night (partially visible) */}
        <rect x="136" y="424" width="228" height="162" rx="12" fill="#18181b" />
        <text x="150" y="444" fill="#52525b" fontSize="10">Last expense</text>
        <rect x="150" y="454" width="18" height="18" rx="5" fill="#1e1b4b" />
        <text x="174" y="467" fill="white" fontSize="12" fontWeight="500">Movie Night</text>
        <text x="150" y="503" fill="white" fontSize="22" fontWeight="700">$32.00</text>
      </g>

      {/* Floating left chip */}
      <rect x="5" y="220" width="118" height="58" rx="12" fill="white" stroke="#e4e4e7" strokeWidth="1.5" />
      <text x="18" y="242" fill="#71717a" fontSize="10">You&apos;re owed</text>
      <text x="18" y="264" fill="#18181b" fontSize="16" fontWeight="700">$14.00</text>
      <circle cx="108" cy="249" r="8" fill="#a1ed6f" />

      {/* Floating right chip */}
      <rect x="377" y="310" width="118" height="58" rx="12" fill="white" stroke="#e4e4e7" strokeWidth="1.5" />
      <text x="390" y="332" fill="#71717a" fontSize="10">Group members</text>
      <circle cx="396" cy="352" r="7" fill="#e4e4e7" />
      <circle cx="409" cy="352" r="7" fill="#d4d4d8" />
      <circle cx="422" cy="352" r="7" fill="#a1ed6f" />
      <text x="434" y="356" fill="#71717a" fontSize="10">+2</text>

      {/* Decorative dots */}
      <circle cx="80" cy="100" r="4" fill="#a1ed6f" opacity="0.5" />
      <circle cx="430" cy="160" r="3" fill="#a1ed6f" opacity="0.4" />
      <circle cx="420" cy="480" r="2.5" fill="#d4d4d8" />
      <circle cx="70" cy="440" r="2" fill="#d4d4d8" />
    </svg>
    </motion.div>
  )
}
