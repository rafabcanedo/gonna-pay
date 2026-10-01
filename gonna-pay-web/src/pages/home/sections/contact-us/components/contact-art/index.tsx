export function ContactArt() {
  return (
    <div className="flex flex-1 items-center justify-center">
      <svg viewBox="0 0 360 320" width="360" height="320" fill="none" xmlns="http://www.w3.org/2000/svg">

        <line x1="180" y1="160" x2="235" y2="65" stroke="#d4d4d8" strokeWidth="1" opacity="0.6" />
        <line x1="180" y1="160" x2="199" y2="268" stroke="#d4d4d8" strokeWidth="1" opacity="0.6" />
        <line x1="180" y1="160" x2="80"  y2="207" stroke="#d4d4d8" strokeWidth="1" opacity="0.6" />

        <circle cx="180" cy="160" r="110" stroke="#e4e4e7" strokeWidth="1.5" />

        <polygon
          points="180,95 193,147 245,160 193,173 180,225 167,173 115,160 167,147"
          fill="#18181b"
        />

        <circle cx="235" cy="65"  r="9" fill="#a1ed6f" />
        <circle cx="199" cy="268" r="6" fill="#a1ed6f" />
        <circle cx="80"  cy="207" r="5" fill="#a1ed6f" />

        <polygon
          points="295,232 299,245 311,248 299,252 295,264 291,252 279,248 291,244"
          fill="#a1ed6f"
        />

      </svg>
    </div>
  )
}
