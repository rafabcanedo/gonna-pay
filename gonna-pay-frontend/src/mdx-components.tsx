import type { MDXComponents } from 'mdx/types'
import { Callout } from '@/app/(dashboard)/docs/components/callout'
import { CategoryBadges } from '@/app/(dashboard)/docs/components/category-badges'

export function useMDXComponents(components: MDXComponents): MDXComponents {
  return {
    ...components,
    Callout,
    CategoryBadges,
  }
}
