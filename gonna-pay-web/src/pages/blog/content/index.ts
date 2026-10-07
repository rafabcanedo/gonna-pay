import { lazy } from "react"

import type { BlogPostMeta } from "../types"

type BlogPost = BlogPostMeta & {
  component: React.LazyExoticComponent<React.ComponentType>
}

export const BLOG_POSTS: BlogPost[] = [
  {
    slug: "how-it-works",
    title: "How Gonna Pay Works",
    description: "A quick guide to splitting costs, managing groups, and settling up with the people you share expenses with.",
    date: "2026-10-06",
    author: "Gonna Pay",
    tags: ["Guide", "Product"],
    component: lazy(() => import("./how-it-works.mdx")),
  },
  {
    slug: "feature-coming-soon",
    title: "What's Coming to Gonna Pay",
    description: "A look at the features we're building next. From a smarter wallet to advanced reports and real-time notifications.",
    date: "2026-10-06",
    author: "Gonna Pay",
    tags: ["Updates", "Product"],
    component: lazy(() => import("./feature-coming-soon.mdx")),
  },
]
