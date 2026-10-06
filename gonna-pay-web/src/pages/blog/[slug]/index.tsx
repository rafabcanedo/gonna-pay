import { Suspense } from "react"
import { useParams, Link, Navigate } from "react-router"
import type { ComponentType } from "react"

import { BLOG_POSTS } from "../content"
import { HowItWorksArt } from "../components/how-it-works-art"
import { FeatureComingSoonArt } from "../components/feature-coming-soon-art"

const artMap: Record<string, ComponentType> = {
  "how-it-works": HowItWorksArt,
  "feature-coming-soon": FeatureComingSoonArt,
}

export function BlogPost() {
  const { slug } = useParams<{ slug: string }>()
  const post = BLOG_POSTS.find(p => p.slug === slug)

  if (!post) return <Navigate to="/blog" replace />

  const Art = artMap[post.slug]
  const Content = post.component

  return (
    <div className="px-8 py-12 max-w-2xl mx-auto w-full">
      <Link
        to="/blog"
        className="inline-flex items-center gap-1.5 text-sm text-muted-foreground hover:text-foreground transition-colors mb-10"
      >
        ← Back to blog
      </Link>

      <div className="bg-zinc-950 w-full rounded-2xl overflow-hidden aspect-video flex items-center justify-center p-6 mb-8">
        <Art />
      </div>

      <h1 className="text-3xl font-semibold tracking-tight text-foreground">{post.title}</h1>

      <div className="flex items-center gap-2 mt-3 text-sm text-muted-foreground">
        <span>by <strong className="text-foreground font-medium">{post.author}</strong></span>
        <span>·</span>
        <span>{post.date}</span>
      </div>

      <div className="flex gap-2 mt-4 flex-wrap">
        {post.tags.map(tag => (
          <span key={tag} className="text-xs px-2.5 py-1 rounded-md bg-zinc-100 text-zinc-600 font-medium">
            {tag}
          </span>
        ))}
      </div>

      <div className="mt-10">
        <Suspense fallback={<div className="text-muted-foreground text-sm">Loading...</div>}>
          <Content />
        </Suspense>
      </div>
    </div>
  )
}
