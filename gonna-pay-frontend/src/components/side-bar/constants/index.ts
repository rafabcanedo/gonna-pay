import { House, Wallet, Receipt, Users, Banknote, BookUser, BookText } from "lucide-react";
import { NavGroup } from "../types";

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
      { title: "How can I start?", url: "/docs/how-can-i-start", icon: BookText },
      { title: "How it works",     url: "/docs/how-it-works",    icon: BookText },
      { title: "Contact Us",       url: "/docs/contact-us",      icon: BookText },
    ],
  },
]
