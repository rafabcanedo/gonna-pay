import { Link } from "react-router"

import type { BlogCardProps } from "./types"

export function BlogCard({ slug, title, description, date, tags, artComponent: Art }: BlogCardProps) {
  return (
    <Link to={`/blog/${slug}`} className="group flex flex-col rounded-2xl border border-zinc-200 overflow-hidden hover:border-zinc-300 transition-colors">
      <div className="bg-zinc-950 w-full aspect-video flex items-center justify-center p-4">
        <Art />
      </div>

      <div className="flex flex-col gap-3 p-5">
        <div className="flex flex-col gap-1">
          <h2 className="text-base font-semibold text-foreground group-hover:text-primary transition-colors">{title}</h2>
          <p className="text-sm text-muted-foreground leading-relaxed">{description}</p>
        </div>

        <div className="flex items-center justify-between mt-1">
          <div className="flex gap-2 flex-wrap">
            {tags.map(tag => (
              <span key={tag} className="text-xs px-2.5 py-1 rounded-md bg-zinc-100 text-zinc-600 font-medium">
                {tag}
              </span>
            ))}
          </div>
          <span className="text-xs text-muted-foreground shrink-0">{date}</span>
        </div>
      </div>
    </Link>
  )
}
