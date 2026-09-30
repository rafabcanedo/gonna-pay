export type LogoSize = 'xs' | 'sm' | 'md'

export interface LogoProps {
  size?: LogoSize
  className?: string
  variant?: 'light' | 'dark'
  color?: string
}
