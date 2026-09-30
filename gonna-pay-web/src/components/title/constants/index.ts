import type { TitleSize, TitleVariant } from "../types"

export const sizeMap: Record<TitleSize, string> = {
  sm: "text-sm",
  md: "text-base",
  lg: "text-lg",
}

export const variantMap: Record<TitleVariant, string> = {
  default: "bg-primary text-primary-foreground px-3 py-1 rounded-md",
  plain:   "text-foreground",
}
