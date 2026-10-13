export function SopTeamArt() {
  return (
    <svg
      viewBox="0 0 480 260"
      width="100%"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
      style={{ fontFamily: "'Geist Variable', sans-serif" }}
    >
      <rect x="60" y="20" width="360" height="220" rx="16" fill="#18181b" />

      <line x1="240" y1="55" x2="297" y2="97" stroke="#27272a" strokeWidth="1" strokeDasharray="4 4" />
      <line x1="297" y1="97" x2="275" y2="164" stroke="#27272a" strokeWidth="1" strokeDasharray="4 4" />
      <line x1="275" y1="164" x2="205" y2="164" stroke="#27272a" strokeWidth="1" strokeDasharray="4 4" />
      <line x1="205" y1="164" x2="183" y2="97" stroke="#27272a" strokeWidth="1" strokeDasharray="4 4" />
      <line x1="183" y1="97" x2="240" y2="55" stroke="#27272a" strokeWidth="1" strokeDasharray="4 4" />
      <line x1="240" y1="55" x2="275" y2="164" stroke="#27272a" strokeWidth="0.5" strokeDasharray="3 5" opacity="0.5" />
      <line x1="240" y1="55" x2="205" y2="164" stroke="#27272a" strokeWidth="0.5" strokeDasharray="3 5" opacity="0.5" />

      {/* Top center */}
      <circle cx="240" cy="55" r="15" stroke="#a1ed6f" strokeWidth="1.5" />
      <path d="M 226 73 Q 240 84 254 73" stroke="#a1ed6f" strokeWidth="1.5" />

      {/* Upper right */}
      <circle cx="297" cy="97" r="13" stroke="#a1ed6f" strokeWidth="1.5" />
      <path d="M 284 114 Q 297 124 310 114" stroke="#a1ed6f" strokeWidth="1.5" />

      {/* Lower right */}
      <circle cx="275" cy="164" r="13" stroke="#a1ed6f" strokeWidth="1.5" />
      <path d="M 262 181 Q 275 191 288 181" stroke="#a1ed6f" strokeWidth="1.5" />

      {/* Lower left */}
      <circle cx="205" cy="164" r="13" stroke="#a1ed6f" strokeWidth="1.5" />
      <path d="M 192 181 Q 205 191 218 181" stroke="#a1ed6f" strokeWidth="1.5" />

      {/* Upper left */}
      <circle cx="183" cy="97" r="13" stroke="#a1ed6f" strokeWidth="1.5" />
      <path d="M 170 114 Q 183 124 196 114" stroke="#a1ed6f" strokeWidth="1.5" />

      {/* Handshake center — Lucide paths scaled 1.8x, centered at (240, 115) */}
      <g transform="translate(218.4 93.4) scale(1.8)" stroke="#f59e0b" strokeWidth="1.4" strokeLinecap="round" strokeLinejoin="round">
        <path d="m11 17 2 2a1 1 0 1 0 3-3" />
        <path d="m14 14 2.5 2.5a1 1 0 1 0 3-3l-3.88-3.88a3 3 0 0 0-4.24 0l-.88.88a1 1 0 1 1-3-3l2.81-2.81a5.79 5.79 0 0 1 7.06-.87l.47.28a2 2 0 0 0 1.42.25L21 4" />
        <path d="m21 3 1 11h-2" />
        <path d="M3 3 2 14l6.5 6.5a1 1 0 1 0 3-3" />
        <path d="M3 4h8" />
      </g>

      <rect x="188" y="208" width="104" height="26" rx="8" fill="#09090b" stroke="#f59e0b" strokeWidth="1.5" />
      <text x="240" y="225" fill="#f59e0b" fontSize="10" fontWeight="600" textAnchor="middle">SOP Team</text>

      <circle cx="66" cy="26" r="2.5" fill="#a1ed6f" opacity="0.4" />
      <circle cx="414" cy="234" r="2" fill="#3f3f46" />
    </svg>
  )
}
