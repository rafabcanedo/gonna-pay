import { Suspense } from 'react'
import { NavBreadcrumb } from './components/nav-breadcrumb'
import { NotificationDropdown } from './components/notification-dropdown'
import { NewDropdown } from './components/new-dropdown'
import { UserDropdown } from './components/user-dropdown'

export function DashboardNavbar() {
  return (
    <header className="flex items-center justify-between h-14 px-8 border-b bg-white shrink-0">
      <Suspense>
        <NavBreadcrumb />
      </Suspense>
      <div className="flex items-center gap-2">
        <NotificationDropdown />
        <NewDropdown />
        <UserDropdown />
      </div>
    </header>
  )
}
