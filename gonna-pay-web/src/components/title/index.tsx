import { cn } from "@/lib/utils"
import type { TitleProps } from "./types"

export function Title({ children, className }: TitleProps) {
  return (
    <h2 className={cn("w-fit bg-primary text-primary-foreground font-semibold px-3 py-1 rounded-md", className)}>
      {children}
    </h2>
  )
}
