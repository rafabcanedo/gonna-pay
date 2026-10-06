import Image from 'next/image'
import gonnaLogo from '@/assets/gonna-logo.svg'
import gLogo from '@/assets/g-logo.svg'
import type { LogoProps } from './types'
import { WIDTH_MAP } from './constants'

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

  const width = WIDTH_MAP[size]
  return (
    <Image
      src={gonnaLogo}
      alt="Gonna Pay"
      width={width}
      height={Math.round(width * (280 / 1060))}
    />
  )
}