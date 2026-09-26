import type { NextConfig } from 'next'
import createMDX from '@next/mdx'
import rehypeSlug from 'rehype-slug'

const withMDX = createMDX({
  options: {
    rehypePlugins: [rehypeSlug],
  },
})

const nextConfig: NextConfig = {
  pageExtensions: ['js', 'jsx', 'md', 'mdx', 'ts', 'tsx'],
}

export default withMDX(nextConfig)
