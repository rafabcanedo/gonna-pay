import gonnaLogo from "@/assets/gonna-logo.svg"
import { widthMap } from "./constants"
import type { LogoProps } from "./types"

export function Logo({ size = 'md', className }: LogoProps) {
  const width = widthMap[size]
  const height = Math.round(width * (280 / 1060))

  return (
    <div style={{ width, height }} className={className}>
      <img src={gonnaLogo} alt="Gonna Pay" width={width} height={height} />
    </div>
  )
}
