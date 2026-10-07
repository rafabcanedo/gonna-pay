import { BLOG_POSTS } from "./content"
import { BlogCard } from "./components/blog-card"
import { HowItWorksArt } from "./components/how-it-works-art"
import { FeatureComingSoonArt } from "./components/feature-coming-soon-art"
import type { ComponentType } from "react"

const artMap: Record<string, ComponentType> = {
  "how-it-works": HowItWorksArt,
  "feature-coming-soon": FeatureComingSoonArt,
}

export function Blog() {
  return (
    <div className="px-8 py-24 max-w-4xl mx-auto w-full">
      <div className="mb-12">
        <h1 className="text-4xl font-semibold tracking-tight text-foreground">Blog</h1>
        <p className="text-muted-foreground mt-3 text-base">Guides, updates and news from Gonna Pay.</p>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
        {BLOG_POSTS.map(post => (
          <BlogCard
            key={post.slug}
            {...post}
            artComponent={artMap[post.slug]}
          />
        ))}
      </div>
    </div>
  )
}
