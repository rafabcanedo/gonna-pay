import { GonnaLogo } from "@/assets/gonna-logo"
import { widthMap } from "./constants"
import type { LogoProps } from "./types"

export function Logo({ size = 'md', className, variant = 'light', color }: LogoProps) {
  const width = widthMap[size]
  const height = Math.round(width * (280 / 1060))

  return (
    <div style={{ width, height }} className={className}>
      <GonnaLogo variant={variant} color={color} width={width} height={height} />
    </div>
  )
}
