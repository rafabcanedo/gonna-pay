"use client"
import { useMutation, useQueryClient } from "@tanstack/react-query"
import { authService } from "@/services/auth"
import { useUser } from "@/providers/contexts/user-context"
import { useRouter } from "next/navigation"
import { toast } from "sonner"

export function useLogoutMutation() {
    const { clearUser } = useUser()
    const queryClient = useQueryClient()
    const router = useRouter()

    return useMutation({
        mutationFn: authService.logout,
        onSuccess: () => {
            clearUser()
            queryClient.clear()
            router.push("/signin")
        },
        onError: () => {
            toast.error("Failed to logout. Try again.")
        },
    })
}
