"use client"

import { Card, CardHeader, CardTitle, CardContent, CardFooter } from "@/components/ui/card"
import { Button } from "@/components/ui/button"
import { useForm, FormProvider } from "react-hook-form"
import { yupResolver } from "@hookform/resolvers/yup"
import { HookFormTextInput } from "@/components/hook-form-text-input"
import Link from "next/link"
import { signUpSchema } from "@/validations/schemas"
import { useAuthMutations } from "@/hooks/mutations/use-auth-mutations"
import { IRegisterAccount } from "./interfaces"

export default function SignUpForm() {

  const { registerMutation } = useAuthMutations()

  const methods = useForm<IRegisterAccount>({
    resolver: yupResolver(signUpSchema),
    mode: "onSubmit",
  })

  const handleRegisterAccount = (data: IRegisterAccount) => {
    registerMutation.mutate(data)
  }

  return (
    <div className="flex items-center justify-center min-h-screen">
      <FormProvider {...methods}>
        <form onSubmit={methods.handleSubmit(handleRegisterAccount)}>
          <Card className="w-[400px]">
            <CardHeader>
              <CardTitle className="text-center">Sign Up</CardTitle>
            </CardHeader>
            <CardContent className="space-y-4">
              <HookFormTextInput title="Name" name="name" label="Jhon Jason" type="text" />
              <HookFormTextInput title="Email" name="email" label="jhon@email.com" type="text" />
              <HookFormTextInput title="Password" name="password" label="Password" type="password" />
              <HookFormTextInput title="Phone" name="phone" label="+55 11 997117911" type="text" />
            </CardContent>
            <CardFooter className="flex justify-end">
              <Button className="w-full bg-primary hover:bg-hover" type="submit" disabled={registerMutation.isPending}>
                {registerMutation.isPending ? "Loading" : "Register"}
              </Button>
            </CardFooter>

            <div className="flex flex-row justify-center text-center mt-6">
              <span className="text-zinc-500 font-poppins text-sm">
                Already have an account?
                <Link href="/signin">
                  <span className="text-blue-500 hover:underline focus:outline-none ml-1">
                    Sign in
                  </span>
                </Link>
              </span>
            </div>
          </Card>
        </form>
      </FormProvider>
    </div>
  )
}