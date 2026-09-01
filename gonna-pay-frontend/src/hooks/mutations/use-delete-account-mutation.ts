"use client"
import { useMutation, useQueryClient } from "@tanstack/react-query"
import { authService } from "@/services/auth"
import { useUser } from "@/providers/contexts/user-context"
import { useRouter } from "next/navigation"
import { toast } from "sonner"

export function useDeleteAccountMutation() {
    const { clearUser } = useUser()
    const queryClient = useQueryClient()
    const router = useRouter()

    return useMutation({
        mutationFn: authService.deleteAccount,
        onSuccess: () => {
            clearUser()
            queryClient.clear()
            router.push("/signin")
        },
        onError: () => {
            toast.error("Failed to delete account. Try again.")
        },
    })
}
