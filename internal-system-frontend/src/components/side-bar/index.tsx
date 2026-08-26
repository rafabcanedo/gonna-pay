'use client'

import * as React from "react";
import Link from "next/link";
import { usePathname } from "next/navigation";
import { House, type LucideIcon } from "lucide-react";

import { SearchForm } from "@/components/search-form";
import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarGroup,
  SidebarGroupContent,
  SidebarGroupLabel,
  SidebarHeader,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
  SidebarRail,
} from "@/components/ui/sidebar";
import { Logo } from "../logo";

type NavItem = {
  title: string
  url: string
  icon: LucideIcon
  isActive?: boolean
}

type NavGroup = {
  title: string
  items: NavItem[]
}

const navMain: NavGroup[] = [
  {
    title: "Initial Steps",
    items: [
      { title: "Home",        url: "/dashboard",   icon: House },
      { title: "My wallet",   url: "/my-wallet",   icon: House },
      { title: "Costs",       url: "/costs",       icon: House },
      { title: "Groups",      url: "/groups",      icon: House },
      { title: "Payments",    url: "/payments",    icon: House },
      { title: "My contacts", url: "/my-contacts", icon: House },
    ],
  },
  {
    title: "Getting Started",
    items: [
      { title: "How can I start?", url: "#", icon: House },
      { title: "How it works",     url: "#", icon: House },
      { title: "Contact Us",       url: "#", icon: House },
    ],
  },
]

const navAccount: NavItem[] = [
  { title: "My Account", url: "/profile", icon: House },
]

export function SideBar({ ...props }: React.ComponentProps<typeof Sidebar>) {
  const pathname = usePathname()

  return (
    <Sidebar collapsible="icon" {...props}>
      <div className="flex items-center justify-center mt-4 mb-4 group-data-[collapsible=icon]:hidden">
        <Logo />
      </div>
      <SidebarHeader>
        <div className="group-data-[collapsible=icon]:hidden">
          <SearchForm />
        </div>
      </SidebarHeader>
      <SidebarContent>
        {navMain.map((group) => (
          <SidebarGroup key={group.title}>
            <SidebarGroupLabel>{group.title}</SidebarGroupLabel>
            <SidebarGroupContent>
              <SidebarMenu>
                {group.items.map((item) => (
                  <SidebarMenuItem key={item.title}>
                    <SidebarMenuButton
                      asChild
                      isActive={item.url !== '#' && pathname.startsWith(item.url)}
                      tooltip={item.title}
                    >
                      <Link href={item.url}>
                        <item.icon />
                        {item.title}
                      </Link>
                    </SidebarMenuButton>
                  </SidebarMenuItem>
                ))}
              </SidebarMenu>
            </SidebarGroupContent>
          </SidebarGroup>
        ))}
      </SidebarContent>
      <SidebarFooter>
        <SidebarMenu>
          {navAccount.map((item) => (
            <SidebarMenuItem key={item.title}>
              <SidebarMenuButton
                asChild
                isActive={pathname === item.url}
                tooltip={item.title}
              >
                <Link href={item.url}>
                  <item.icon />
                  {item.title}
                </Link>
              </SidebarMenuButton>
            </SidebarMenuItem>
          ))}
        </SidebarMenu>
      </SidebarFooter>
      <SidebarRail />
    </Sidebar>
  )
}
