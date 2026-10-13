import path from "path"
import react from "@vitejs/plugin-react"
import tailwindcss from "@tailwindcss/vite"
import { defineConfig } from "vite"
import mdx from "@mdx-js/rollup"
import remarkFrontmatter from "remark-frontmatter"
import rehypeSlug from "rehype-slug"

export default defineConfig({
  plugins: [
    { enforce: "pre", ...mdx({ remarkPlugins: [remarkFrontmatter], rehypePlugins: [rehypeSlug] }) },
    react(),
    tailwindcss(),
  ],
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./src"),
    },
  },
})
