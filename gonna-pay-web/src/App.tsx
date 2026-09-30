import { BrowserRouter, Routes, Route } from "react-router"
import { Layout } from "@/components/layout"
import { Home } from "@/pages/home"
import { Pricing } from "@/pages/pricing"
import { Blog } from "@/pages/blog"

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route element={<Layout />}>
          <Route index element={<Home />} />
          <Route path="/pricing" element={<Pricing />} />
          <Route path="/blog" element={<Blog />} />
        </Route>
      </Routes>
    </BrowserRouter>
  )
}

export default App
