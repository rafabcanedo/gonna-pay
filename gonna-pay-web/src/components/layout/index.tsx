import { Outlet } from "react-router"
import { Navbar } from "@/components/navbar"
import { Footer } from "@/components/footer"

export function Layout() {
  return (
    <>
      <Navbar />
      <Outlet />
      <Footer />
    </>
  )
}
