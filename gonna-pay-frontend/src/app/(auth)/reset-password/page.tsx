"use client"

import { Suspense } from "react"
import { useSearchParams } from "next/navigation"
import { useForm, FormProvider } from "react-hook-form"
import { yupResolver } from "@hookform/resolvers/yup"
import { Card, CardContent, CardFooter, CardHeader } from "@/components/ui/card"
import { Button } from "@/components/ui/button"
import { Logo } from "@/components/logo"
import { HookFormTextInput } from "@/components/hook-form-text-input"
import { resetPasswordSchema } from "@/validations/schemas"
import { useAuthMutations } from "@/hooks/mutations/use-auth-mutations"
import type { IResetPasswordForm } from "./types"

function ResetPasswordContent() {
  const searchParams = useSearchParams()
  const token = searchParams.get("token") ?? ""
  const { resetPasswordMutation } = useAuthMutations()

  const methods = useForm<IResetPasswordForm>({
    resolver: yupResolver(resetPasswordSchema),
    mode: "onSubmit",
  })

  const handleSubmit = (data: IResetPasswordForm) => {
    resetPasswordMutation.mutate({ token, password: data.password })
  }

  if (!token) {
    return (
      <div className="min-h-screen flex flex-col justify-center items-center">
        <div className="w-full max-w-2xl bg-white rounded-lg shadow-md px-12 py-16 mx-auto">
          <div className="flex flex-col items-center justify-center gap-1">
            <span className="text-zinc-900 font-poppins text-lg">
              Invalid link
            </span>
            <span className="text-zinc-500 font-poppins text-base text-center">
              This password reset link is invalid or has expired.
            </span>
          </div>
        </div>
      </div>
    )
  }

  return (
    <div className="flex items-center justify-center min-h-screen">
      <FormProvider {...methods}>
        <form method="post" onSubmit={methods.handleSubmit(handleSubmit)}>
          <Card className="w-[400px]">
            <CardHeader className="flex items-center justify-center">
              <Logo size="md" />
            </CardHeader>

            <CardContent className="space-y-4">
              <HookFormTextInput
                name="password"
                title="New Password"
                label="New password"
                type="password"
              />
              <HookFormTextInput
                name="confirmPassword"
                title="Confirm Password"
                label="Confirm your password"
                type="password"
              />
            </CardContent>

            <CardFooter className="flex justify-end">
              <Button
                className="w-full bg-primary hover:bg-hover"
                type="submit"
                disabled={resetPasswordMutation.isPending}
              >
                {resetPasswordMutation.isPending ? "Loading" : "Reset Password"}
              </Button>
            </CardFooter>
          </Card>
        </form>
      </FormProvider>
    </div>
  )
}

export default function ResetPasswordPage() {
  return (
    <Suspense>
      <ResetPasswordContent />
    </Suspense>
  )
}
