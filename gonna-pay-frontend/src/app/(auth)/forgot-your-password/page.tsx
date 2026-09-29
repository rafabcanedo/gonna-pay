"use client";

import React from "react";
import { useForm, FormProvider } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import { useRouter } from "next/navigation";
import { Button } from "@/components/ui/button";
import { forgotPasswordSchema } from "@/validations/schemas";
import { HookFormTextInput } from "@/components/hook-form-text-input";
import { useAuthMutations } from "@/hooks/mutations/use-auth-mutations";
import { IForgotPassword } from "./types"

export default function ForgotYourPassword() {

  const { forgotPasswordMutation } = useAuthMutations();

  const methods = useForm<IForgotPassword>({
    resolver: yupResolver(forgotPasswordSchema),
    mode: "onSubmit",
    defaultValues: { email: "" },
  });

  const handleSubmit = (data: IForgotPassword) => {
    forgotPasswordMutation.mutate({ email: data.email });
  };

  const router = useRouter();

  return (
    <div className="min-h-screen flex flex-col justify-center items-center">
      <div className="w-full max-w-2xl bg-white rounded-lg shadow-md px-12 py-16 mx-auto">
        <div className="flex flex-col items-center justify-center gap-1 mb-4">
          <span className="text-zinc-900 font-poppins text-lg">
            Forgot your Password?
          </span>
          <span className="text-zinc-500 font-poppins text-base">
            Please enter your email for the code
          </span>
        </div>

        <FormProvider {...methods}>
          <form
            method="post"
            onSubmit={methods.handleSubmit(handleSubmit)}
            className="flex flex-col items-center justify-center space-y-1"
          >
            <div className="w-70">
              <HookFormTextInput
                name="email"
                title="Email"
                label="Enter your email"
                type="email"
              />
            </div>

            <div className="mt-4">
              <Button className="bg-primary hover:bg-hover" size="sm" type="submit" disabled={forgotPasswordMutation.isPending}>
                {forgotPasswordMutation.isPending ? "Loading" : "Reset my password"}
              </Button>
            </div>
          </form>
        </FormProvider>

        <div className="flex justify-end mt-4">
          <button
            onClick={() => router.back()}
            className="bg-transparent font-poppins text-xs text-secondary hover:text-hover h-8 px-4 hover:bg-tambo-primary-hover/60 rounded-sm cursor-pointer"
          >
            Go Back
          </button>
        </div>
      </div>
    </div>
  );
}