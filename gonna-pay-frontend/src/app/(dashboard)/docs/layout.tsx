import { ReactNode } from 'react'
import { TableOfContents } from './components/table-of-contents'

export default function DocsLayout({ children }: { children: ReactNode }) {
  return (
    <div className="flex gap-16 py-8 max-w-5xl">
      <article className="flex-1 min-w-0 prose prose-zinc max-w-none">
        {children}
      </article>
      <aside className="w-48 shrink-0">
        <TableOfContents />
      </aside>
    </div>
  )
}
