"use client"
import { ReactNode } from "react"
import { ReactQueryProvider } from "@/lib/react-query"
import { UserProvider } from "@/providers/contexts/user-context"

export function AppProviders({ children }: { children: ReactNode }) {
    return (
        <ReactQueryProvider>
            <UserProvider>
                {children}
            </UserProvider>
        </ReactQueryProvider>
    )
}
