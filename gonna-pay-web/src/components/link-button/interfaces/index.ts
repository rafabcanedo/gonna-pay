import type { VariantProps } from "class-variance-authority"
import type { buttonVariants } from "@/components/ui/button"

export interface LinkButtonProps extends VariantProps<typeof buttonVariants> {
  href: string
  children: React.ReactNode
  className?: string
  onClick?: React.MouseEventHandler<HTMLAnchorElement>
}
