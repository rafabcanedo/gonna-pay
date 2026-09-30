import { Outlet } from "react-router"
import { Navbar } from "@/components/navbar"

export function Layout() {
  return (
    <>
      <Navbar />
      <Outlet />
    </>
  )
}
