"use client"
import { createContext, useContext, ReactNode } from "react"
import { useQuery, useQueryClient } from "@tanstack/react-query"
import { authService, IUserData } from "@/services/auth"
import { UserContextType } from "./types"

const UserContext = createContext<UserContextType | null>(null)

export function UserProvider({ children }: { children: ReactNode }) {
    const queryClient = useQueryClient()

    const { data: user = null, isLoading } = useQuery<IUserData | null>({
        queryKey: ["user"],
        queryFn: async () => {
            try {
                return await authService.getProfile()
            } catch {
                return null
            }
        },
        retry: false,
        staleTime: Infinity,
    })

    function clearUser() {
        queryClient.removeQueries({ queryKey: ["user"] })
    }

    return (
        <UserContext.Provider value={{ user, isLoading, clearUser }}>
            {children}
        </UserContext.Provider>
    )
}

export function useUser() {
    const context = useContext(UserContext)
    if (!context) throw new Error("useUser must be used within UserProvider")
    return context
}
