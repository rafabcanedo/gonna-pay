'use client'

import { useState } from 'react'
import { Bell, BellRing } from 'lucide-react'
import { Button } from '@/components/ui/button'
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu'

export function NotificationDropdown() {
  const [isOpen, setIsOpen] = useState(false)

  return (
    <DropdownMenu onOpenChange={setIsOpen}>
      <DropdownMenuTrigger asChild>
        <Button variant="ghost" size="icon">
          {isOpen ? <BellRing className="h-4 w-4" /> : <Bell className="h-4 w-4" />}
        </Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent align="end" className="w-64">
        <div className="flex flex-col items-center justify-center py-8 gap-2 text-sm text-muted-foreground">
          <Bell className="h-8 w-8 opacity-40" />
          No notifications yet
        </div>
      </DropdownMenuContent>
    </DropdownMenu>
  )
}
