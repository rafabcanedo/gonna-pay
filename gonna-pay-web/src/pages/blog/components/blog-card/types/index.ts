import type { ComponentType } from "react"

import type { BlogPostMeta } from "@/pages/blog/types"

export interface BlogCardProps extends BlogPostMeta {
  artComponent: ComponentType
}
