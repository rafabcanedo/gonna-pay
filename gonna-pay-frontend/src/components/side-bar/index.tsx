'use client'

import * as React from "react";
import Link from "next/link";
import { usePathname } from "next/navigation";

import { SearchForm } from "@/components/search-form";
import {
  Sidebar,
  SidebarContent,
  SidebarGroup,
  SidebarGroupContent,
  SidebarGroupLabel,
  SidebarHeader,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
  SidebarRail,
  SidebarTrigger,
} from "@/components/ui/sidebar";
import { Logo } from "../logo";
import { navMain } from "./constants";

export function SideBar({ ...props }: React.ComponentProps<typeof Sidebar>) {
  const pathname = usePathname()

  return (
    <Sidebar collapsible="icon" {...props}>
      <SidebarHeader>
        <div className="flex items-center justify-between group-data-[collapsible=icon]:justify-center px-2 py-2">
          <div className="group-data-[collapsible=icon]:hidden">
            <Logo size="sm" />
          </div>
          <SidebarTrigger />
        </div>
        <div className="hidden group-data-[collapsible=icon]:flex justify-center pb-2">
          <Logo variant="icon" />
        </div>
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
      <SidebarRail />
    </Sidebar>
  )
}
