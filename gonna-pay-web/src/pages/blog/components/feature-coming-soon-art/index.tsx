export function FeatureComingSoonArt() {
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

      {/* Wallet icon — top left */}
      <g opacity="0.35">
        <rect x="96" y="56" width="52" height="38" rx="7" stroke="#a1ed6f" strokeWidth="1.5" />
        <rect x="108" y="50" width="28" height="10" rx="4" stroke="#a1ed6f" strokeWidth="1.5" />
        <circle cx="134" cy="75" r="5" fill="#a1ed6f" opacity="0.5" />
      </g>

      {/* Chart / report icon — top right */}
      <g opacity="0.35">
        <rect x="332" y="52" width="52" height="46" rx="7" stroke="#a1ed6f" strokeWidth="1.5" />
        <rect x="342" y="78" width="8" height="12" rx="2" fill="#a1ed6f" opacity="0.6" />
        <rect x="356" y="68" width="8" height="22" rx="2" fill="#a1ed6f" opacity="0.6" />
        <rect x="370" y="60" width="8" height="30" rx="2" fill="#a1ed6f" opacity="0.6" />
      </g>

      {/* Bell / notification icon — bottom left */}
      <g opacity="0.35">
        <path d="M122 176 C122 166 130 158 140 158 C150 158 158 166 158 176 L162 188 H118 L122 176Z" stroke="#a1ed6f" strokeWidth="1.5" />
        <rect x="136" y="188" width="8" height="6" rx="3" stroke="#a1ed6f" strokeWidth="1.5" />
        <circle cx="152" cy="158" r="5" fill="#a1ed6f" opacity="0.8" />
      </g>

      {/* Gear / settings icon — bottom right */}
      <g opacity="0.35">
        <circle cx="340" cy="178" r="10" stroke="#a1ed6f" strokeWidth="1.5" />
        <circle cx="340" cy="178" r="4" fill="#a1ed6f" opacity="0.4" />
        <rect x="338" y="163" width="4" height="6" rx="2" fill="#a1ed6f" opacity="0.6" />
        <rect x="338" y="187" width="4" height="6" rx="2" fill="#a1ed6f" opacity="0.6" />
        <rect x="349" y="176" width="6" height="4" rx="2" fill="#a1ed6f" opacity="0.6" />
        <rect x="325" y="176" width="6" height="4" rx="2" fill="#a1ed6f" opacity="0.6" />
        <rect x="347" y="168" width="4" height="6" rx="2" fill="#a1ed6f" opacity="0.6" transform="rotate(45 349 170)" />
        <rect x="326" y="181" width="4" height="6" rx="2" fill="#a1ed6f" opacity="0.6" transform="rotate(45 328 184)" />
        <rect x="347" y="181" width="4" height="6" rx="2" fill="#a1ed6f" opacity="0.6" transform="rotate(-45 349 184)" />
        <rect x="326" y="168" width="4" height="6" rx="2" fill="#a1ed6f" opacity="0.6" transform="rotate(-45 328 170)" />
      </g>

      {/* Central "Soon" badge */}
      <rect x="176" y="106" width="128" height="48" rx="14" fill="#09090b" stroke="#a1ed6f" strokeWidth="1.5" />
      <text x="240" y="136" fill="#a1ed6f" fontSize="20" fontWeight="700" textAnchor="middle">Soon</text>

      {/* Subtle connecting lines */}
      <line x1="148" y1="80" x2="190" y2="118" stroke="#27272a" strokeWidth="1" strokeDasharray="4 4" />
      <line x1="340" y1="90" x2="296" y2="118" stroke="#27272a" strokeWidth="1" strokeDasharray="4 4" />
      <line x1="152" y1="174" x2="190" y2="148" stroke="#27272a" strokeWidth="1" strokeDasharray="4 4" />
      <line x1="326" y1="174" x2="296" y2="148" stroke="#27272a" strokeWidth="1" strokeDasharray="4 4" />

      {/* Decorative dots */}
      <circle cx="240" cy="42" r="2.5" fill="#a1ed6f" opacity="0.3" />
      <circle cx="240" cy="218" r="2.5" fill="#a1ed6f" opacity="0.3" />
    </svg>
  )
}
