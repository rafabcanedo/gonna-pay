import type { IPropsCategoryBadges } from './interfaces'

export function CategoryBadges({ items }: IPropsCategoryBadges) {
  return (
    <div className="not-prose flex flex-wrap gap-2 my-3">
      {items.map((item) => (
        <span
          key={item}
          className="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-zinc-100 text-zinc-700 border border-zinc-200"
        >
          {item}
        </span>
      ))}
    </div>
  )
}
