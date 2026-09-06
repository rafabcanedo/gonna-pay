import Image from 'next/image'
import gonnaLogo from '@/assets/gonna-logo.svg'
import gLogo from '@/assets/g-logo.svg'
import type { LogoProps, LogoSize } from './types'

const widthMap: Record<LogoSize, number> = {
  sm: 140,
  md: 220,
}

export const Logo = ({ size = 'md', variant = 'default' }: LogoProps) => {
  if (variant === 'icon') {
    return (
      <Image
        src={gLogo}
        alt="Gonna Pay"
        width={32}
        height={35}
      />
    )
  }

  const width = widthMap[size]
  return (
    <Image
      src={gonnaLogo}
      alt="Gonna Pay"
      width={width}
      height={Math.round(width * (280 / 1060))}
    />
  )
}