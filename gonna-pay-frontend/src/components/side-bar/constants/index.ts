import { House, Wallet, Receipt, Users, Banknote, BookUser, CircleUser, BookText } from "lucide-react";
import { NavGroup, NavItem } from "../types";

export const navMain: NavGroup[] = [
  {
    title: "Initial Steps",
    items: [
      { title: "Home",        url: "/dashboard",   icon: House },
      { title: "My wallet",   url: "/my-wallet",   icon: Wallet },
      { title: "Costs",       url: "/costs",       icon: Receipt },
      { title: "Groups",      url: "/groups",      icon: Users },
      { title: "Payments",    url: "/payments",    icon: Banknote },
      { title: "My contacts", url: "/my-contacts", icon: BookUser },
    ],
  },
  {
    title: "Getting Started",
    items: [
      { title: "How can I start?", url: "#", icon: BookText },
      { title: "How it works",     url: "#", icon: BookText },
      { title: "Contact Us",       url: "#", icon: BookText },
    ],
  },
]

export const navAccount: NavItem[] = [
  { title: "My Account", url: "/profile", icon: CircleUser },
]
