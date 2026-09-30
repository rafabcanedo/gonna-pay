import { cn } from "@/lib/utils"
import { sizeMap, variantMap } from "./constants"
import type { TitleProps } from "./types"

export function Title({ children, variant = 'default', size = 'md', className }: TitleProps) {
  return (
    <h2 className={cn("w-fit font-semibold", sizeMap[size], variantMap[variant], className)}>
      {children}
    </h2>
  )
}
