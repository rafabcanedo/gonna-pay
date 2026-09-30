import type { ReactNode } from "react"

export type TitleVariant = 'default' | 'plain'
export type TitleSize = 'sm' | 'md' | 'lg'

export interface TitleProps {
  children: ReactNode
  variant?: TitleVariant
  size?: TitleSize
  className?: string
}
