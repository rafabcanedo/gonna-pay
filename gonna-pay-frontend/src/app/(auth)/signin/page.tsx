"use client"

import { Card, CardHeader, CardContent, CardFooter } from "@/components/ui/card"
import { Logo } from "@/components/logo"
import { Button } from "@/components/ui/button"
import { useForm, FormProvider } from "react-hook-form"
import { yupResolver } from "@hookform/resolvers/yup"
import { HookFormTextInput } from "@/components/hook-form-text-input"
import Link from "next/link"
import { signInSchema } from "@/validations/schemas"
import { useAuthMutations } from "@/hooks/mutations/use-auth-mutations"
import { IForm } from "./interfaces"

export default function SignInForm() {

  const { loginMutation } = useAuthMutations()

  const methods = useForm<IForm>({
    resolver: yupResolver(signInSchema),
    mode: "onSubmit",
  });

  const handleOnSubmit = (data: IForm) => {
    loginMutation.mutate(data)
  }

  return (
    <div className="flex items-center justify-center min-h-screen">
      <FormProvider {...methods}>
        <form onSubmit={methods.handleSubmit(handleOnSubmit)}>
          <Card className="w-[350px]">
            <CardHeader className="flex items-center justify-center">
              <Logo size="md" />
            </CardHeader>
            <CardContent className="space-y-4">
              <HookFormTextInput title="Email" name="email" label="jhon@email.com" type="text" />
              <HookFormTextInput title="Password" name="password" label="Password" type="password" />
            </CardContent>
            <CardFooter className="flex flex-col justify-end">
              <div className="flex justify-end w-72 mb-2">
                <Link
                  href="/forgot-your-password"
                  className="text-sm text-blue-500 hover:underline focus:outline-none"
                >
                  Forgot Password?
                </Link>
              </div>

              <Button className="w-full bg-primary hover:bg-hover" type="submit" disabled={loginMutation.isPending}>
                {loginMutation.isPending ? "Loading" : "Login"}
              </Button>
            </CardFooter>
            <div className="flex flex-row justify-center text-center mt-6">
              <span className="text-sm">
                Don&apos;t have an account?
                <Link
                  href="/signup"
                  className="hover:text-secondary hover:underline focus:outline-none ml-1"
                >
                  Sign up
                </Link>
              </span>
            </div>
          </Card>
        </form>
      </FormProvider>
    </div>
  )
}