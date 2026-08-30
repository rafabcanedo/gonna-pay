import { useMutation, useQueryClient } from "@tanstack/react-query";
import { authService, ISignInRequest, ISignUpRequest, IAuthResponse } from "@/services/auth";
import { toast } from "sonner";
import { useRouter } from "next/navigation";

export interface IRestError {
  message: string;
  error: string;
  code: number;
  causes?: Array<{
    field: string;
    message: string;
  }>;
}

export function useAuthMutations() {
  const router = useRouter();
  const queryClient = useQueryClient();

  const loginMutation = useMutation<IAuthResponse, IRestError, ISignInRequest>({
    mutationFn: authService.signIn,
    onSuccess: (data) => {
      toast.success(data.message || "Login successful!");
      
      queryClient.setQueryData(["user"], data.user ?? null)
      
      router.push("/dashboard");
      
      router.refresh();
    },
    onError: (error) => {
      toast.error(error.message || "Invalid credentials or server error");
    },
  });

  const registerMutation = useMutation<{ message: string }, IRestError, ISignUpRequest>({
    mutationFn: authService.signUp,
    onSuccess: (data) => {
      toast.success(data.message || "Account created! You can now sign in.");
      router.push("/signin");
    },
    onError: (error) => {
      toast.error(error.message || "Failed to create account. Try a different email.");
    },
  });

  return {
    loginMutation,
    registerMutation,
  };
}
