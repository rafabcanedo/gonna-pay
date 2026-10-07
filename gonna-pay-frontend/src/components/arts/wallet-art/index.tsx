export function WalletArt() {
  return (
    <svg
      width="220"
      height="200"
      viewBox="0 0 220 200"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
    >
      {/* Coin - left */}
      <circle cx="62" cy="68" r="16" fill="#fbbf24" />
      <circle cx="62" cy="68" r="12" fill="#fcd34d" />
      <text x="62" y="73" textAnchor="middle" fontSize="11" fontWeight="bold" fill="#92400e">$</text>

      {/* Coin - top center */}
      <circle cx="105" cy="48" r="13" fill="#fbbf24" />
      <circle cx="105" cy="48" r="9.5" fill="#fcd34d" />
      <text x="105" y="52.5" textAnchor="middle" fontSize="9" fontWeight="bold" fill="#92400e">$</text>

      {/* Coin - right */}
      <circle cx="152" cy="62" r="11" fill="#fbbf24" />
      <circle cx="152" cy="62" r="8" fill="#fcd34d" />
      <text x="152" y="66" textAnchor="middle" fontSize="8" fontWeight="bold" fill="#92400e">$</text>

      {/* Bill flying out */}
      <g transform="rotate(-20, 145, 85)">
        <rect x="125" y="76" width="42" height="26" rx="4" fill="#a1ed6f" />
        <rect x="131" y="81" width="30" height="4" rx="1" fill="#76b84a" />
        <rect x="131" y="89" width="22" height="4" rx="1" fill="#76b84a" />
      </g>

      {/* Wallet body */}
      <g transform="rotate(-6, 86, 138)">
        <rect x="42" y="112" width="96" height="60" rx="10" fill="#a07848" />
        <rect x="38" y="108" width="96" height="60" rx="10" fill="#c89b6e" />
        <rect x="38" y="108" width="96" height="28" rx="10" fill="#ddb07e" />
        <line x1="38" y1="136" x2="134" y2="136" stroke="#a07848" strokeWidth="1.5" />
        <rect x="50" y="142" width="32" height="10" rx="3" fill="#ddb07e" opacity="0.7" />
        <rect x="88" y="142" width="28" height="10" rx="3" fill="#ddb07e" opacity="0.5" />
      </g>

      {/* Face */}
      <g transform="rotate(-6, 86, 138)">
        <circle cx="72" cy="120" r="5" fill="white" />
        <circle cx="72" cy="121" r="3" fill="#1e1e2e" />
        <circle cx="100" cy="120" r="5" fill="white" />
        <circle cx="100" cy="121" r="3" fill="#1e1e2e" />
        <ellipse cx="86" cy="130" rx="6" ry="5" fill="#1e1e2e" />
        <ellipse cx="86" cy="131" rx="4" ry="3" fill="#c87878" />
      </g>

      {/* Speed lines */}
      <line x1="30" y1="118" x2="10" y2="118" stroke="#d1d5db" strokeWidth="2" strokeLinecap="round" />
      <line x1="28" y1="130" x2="6" y2="130" stroke="#d1d5db" strokeWidth="2" strokeLinecap="round" />
      <line x1="30" y1="142" x2="12" y2="142" stroke="#d1d5db" strokeWidth="2" strokeLinecap="round" />
    </svg>
  )
}
