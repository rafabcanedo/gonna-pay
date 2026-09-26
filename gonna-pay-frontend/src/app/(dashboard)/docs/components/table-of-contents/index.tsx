'use client'

import { useEffect, useState } from 'react'
import type { Heading } from './types'

export function TableOfContents() {
  const [headings, setHeadings] = useState<Heading[]>([])

  useEffect(() => {
    const elements = document.querySelectorAll('article h2, article h3')
    const list = Array.from(elements).map((el) => ({
      id: el.id,
      text: el.textContent ?? '',
      level: Number(el.tagName[1]),
    }))
    setHeadings(list)
  }, [])

  if (headings.length === 0) return null

  return (
    <nav className="sticky top-8">
      <p className="text-sm font-medium text-zinc-700 mb-3">On this page</p>
      <ul className="flex flex-col gap-1.5">
        {headings.map((heading) => (
          <li
            key={heading.id}
            style={{ paddingLeft: heading.level === 3 ? '0.75rem' : '0' }}
          >
            <a
              href={`#${heading.id}`}
              className="text-sm text-zinc-400 hover:text-zinc-900 transition-colors"
            >
              {heading.text}
            </a>
          </li>
        ))}
      </ul>
    </nav>
  )
}
