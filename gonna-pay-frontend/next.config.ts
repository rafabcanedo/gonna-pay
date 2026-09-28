import type { NextConfig } from 'next'
import createMDX from '@next/mdx'

const withMDX = createMDX({
  options: {
    rehypePlugins: ['rehype-slug'],
  },
})

const nextConfig: NextConfig = {
  pageExtensions: ['js', 'jsx', 'md', 'mdx', 'ts', 'tsx'],
}

export default withMDX(nextConfig)
