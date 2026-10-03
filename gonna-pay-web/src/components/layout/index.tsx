import { useEffect } from "react"
import { Outlet, useLocation } from "react-router"

import { Navbar } from "@/components/navbar"
import { Footer } from "@/components/footer"

export function Layout() {
  const { pathname, hash } = useLocation()

  useEffect(() => {
    if (hash) {
      const element = document.querySelector(hash)
      if (element) {
        element.scrollIntoView({ behavior: "smooth" })
      }
    } else {
      window.scrollTo(0, 0)
    }
  }, [pathname, hash])

  return (
    <>
      <Navbar />
      <Outlet />
      <Footer />
    </>
  )
}
