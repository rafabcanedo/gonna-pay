import { useMutation, useQueryClient } from "@tanstack/react-query";
import { authService, ISignInRequest, ISignUpRequest, IAuthResponse, IMessageResponse, IForgotPasswordRequest, IResetPasswordRequest, IVerifyEmailRequest } from "@/services/auth";
import { IRestError } from "@/types";
import { toast } from "sonner";
import { useRouter } from "next/navigation";

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

  const registerMutation = useMutation<IMessageResponse, IRestError, ISignUpRequest>({
    mutationFn: authService.signUp,
    onSuccess: (data) => {
      toast.success(data.message || "Account created! You can now sign in.");
      router.push("/signin");
    },
    onError: (error) => {
      toast.error(error.message || "Failed to create account. Try a different email.");
    },
  });

  const forgotPasswordMutation = useMutation<IMessageResponse, IRestError, IForgotPasswordRequest>({
    mutationFn: authService.forgotPassword,
    onSuccess: (data) => {
      toast.success(data.message || "Reset link sent to your email.");
    },
    onError: (error) => {
      toast.error(error.message || "Failed to send reset link.");
    },
  });

  const resetPasswordMutation = useMutation<IMessageResponse, IRestError, IResetPasswordRequest>({
    mutationFn: authService.resetPassword,
    onSuccess: (data) => {
      toast.success(data.message || "Password updated successfully!");
      router.push("/signin");
    },
    onError: (error) => {
      toast.error(error.message || "Failed to reset password.");
    },
  });

  const verifyEmailMutation = useMutation<IMessageResponse, IRestError, IVerifyEmailRequest>({
    mutationFn: authService.verifyEmail,
  });

  return {
    loginMutation,
    registerMutation,
    forgotPasswordMutation,
    resetPasswordMutation,
    verifyEmailMutation,
  };
}
