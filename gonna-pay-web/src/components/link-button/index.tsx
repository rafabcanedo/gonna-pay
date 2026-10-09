import { cn } from "@/lib/utils"
import { buttonVariants } from "@/components/ui/button"
import type { LinkButtonProps } from "./interfaces"

export function LinkButton({ href, children, className, variant, size, onClick }: LinkButtonProps) {
  return (
    <a
      href={href}
      onClick={onClick}
      className={cn(buttonVariants({ variant, size }), className)}
    >
      {children}
    </a>
  )
}
