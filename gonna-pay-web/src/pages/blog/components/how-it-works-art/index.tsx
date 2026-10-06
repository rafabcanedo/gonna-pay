export function HowItWorksArt() {
  return (
    <svg
      viewBox="0 0 480 260"
      width="100%"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
      style={{ fontFamily: "'Geist Variable', sans-serif" }}
    >
      {/* Card background */}
      <rect x="60" y="20" width="360" height="220" rx="16" fill="#18181b" />

      {/* Header */}
      <rect x="60" y="20" width="360" height="48" rx="16" fill="#111111" />
      <rect x="60" y="52" width="360" height="16" fill="#111111" />
      <circle cx="84" cy="44" r="8" fill="#a1ed6f" />
      <text x="98" y="49" fill="white" fontSize="13" fontWeight="600">Weekend Trip</text>
      <text x="396" y="49" fill="#52525b" fontSize="11" textAnchor="end">Travel</text>

      {/* Total */}
      <text x="80" y="100" fill="#52525b" fontSize="10">Total cost</text>
      <text x="80" y="124" fill="white" fontSize="28" fontWeight="700">$360.00</text>

      {/* Divider */}
      <line x1="80" y1="140" x2="400" y2="140" stroke="#27272a" strokeWidth="1" />

      {/* Split rows */}
      <circle cx="90" cy="162" r="5" fill="#a1ed6f" />
      <text x="102" y="167" fill="white" fontSize="12">You</text>
      <text x="400" y="167" fill="#a1ed6f" fontSize="12" fontWeight="600" textAnchor="end">$120.00</text>
      <text x="340" y="167" fill="#52525b" fontSize="10" textAnchor="end">33%</text>

      <circle cx="90" cy="188" r="5" fill="#a1ed6f" />
      <text x="102" y="193" fill="white" fontSize="12">Carlos</text>
      <text x="400" y="193" fill="#a1ed6f" fontSize="12" fontWeight="600" textAnchor="end">$120.00</text>
      <text x="340" y="193" fill="#52525b" fontSize="10" textAnchor="end">33%</text>

      <circle cx="90" cy="214" r="5" fill="none" stroke="#3f3f46" strokeWidth="1.5" />
      <text x="102" y="219" fill="#71717a" fontSize="12">Maria</text>
      <text x="400" y="219" fill="#71717a" fontSize="12" textAnchor="end">$120.00</text>
      <text x="340" y="219" fill="#52525b" fontSize="10" textAnchor="end">34%</text>

      {/* Floating badge */}
      <rect x="282" y="82" width="118" height="36" rx="10" fill="#09090b" stroke="#27272a" strokeWidth="1" />
      <circle cx="298" cy="100" r="5" fill="#a1ed6f" />
      <text x="308" y="104" fill="white" fontSize="10" fontWeight="500">Split equally</text>

      {/* Decorative dots */}
      <circle cx="66" cy="26" r="2.5" fill="#a1ed6f" opacity="0.4" />
      <circle cx="414" cy="234" r="2" fill="#3f3f46" />
    </svg>
  )
}
