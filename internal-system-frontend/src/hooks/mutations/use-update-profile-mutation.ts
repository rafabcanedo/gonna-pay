import { useMutation, useQueryClient } from "@tanstack/react-query"
import { authService, IUpdateProfileRequest } from "@/services/auth"
import { IRestError } from "./use-auth-mutations"
import { toast } from "sonner"

export function useUpdateProfileMutation() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (data: IUpdateProfileRequest) => authService.updateProfile(data),
    onSuccess: (response) => {
      queryClient.setQueryData(["user"], response.user)
      toast.success(response.message || "Profile updated successfully.")
    },
    onError: (error: IRestError) => {
      toast.error(error.message || "Failed to update profile. Try again.")
    },
  })
}
