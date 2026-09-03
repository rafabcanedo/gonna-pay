"use client"

import { useEffect } from "react"
import { Card, CardHeader, CardTitle, CardContent, CardFooter } from "@/components/ui/card"
import { Button } from "@/components/ui/button"
import { FormProvider, useForm, type SubmitHandler } from "react-hook-form"
import { yupResolver } from "@hookform/resolvers/yup"
import { HookFormTextInput } from "@/components/hook-form-text-input"
import { HookFormPhoneInput, formatPhone } from "@/components/hook-form-phone-input"
import { profileSchema } from "@/validations/schemas"
import { useUser } from "@/providers/contexts/user-context"
import { useUpdateProfileMutation } from "@/hooks/mutations/use-update-profile-mutation"
import { ProfileFormValues } from './types'

export const ProfileForm = () => {
  const { user } = useUser()
  const updateProfile = useUpdateProfileMutation()

  const methods = useForm<ProfileFormValues>({
    resolver: yupResolver(profileSchema),
    defaultValues: { name: "", email: "", phone: "", street: "", neighborhood: "", zip: "" },
    mode: "onChange",
  })

  const { handleSubmit, formState, reset } = methods
  const { isSubmitting, isDirty } = formState

  useEffect(() => {
    if (user) {
      reset({ name: user.name, email: user.email, phone: formatPhone(user.phone ?? '') })
    }
  }, [user, reset])

  const handleSubmitProfile: SubmitHandler<ProfileFormValues> = async (data) => {
    await updateProfile.mutateAsync(data)
    reset(data)
  }

  return (
    <div>
      <FormProvider {...methods}>
        <form onSubmit={handleSubmit(handleSubmitProfile)}>
          <Card className="w-[800px]">
            <CardHeader>
              <CardTitle className="text-center">Your Profile</CardTitle>
            </CardHeader>
            <CardContent className="grid grid-cols-2 gap-4 text-zinc-600">
              <HookFormTextInput title="Name" name="name" label="John Jason" type="text" />
              <HookFormTextInput title="Email" name="email" label="john@example.com" type="text" />
              <HookFormPhoneInput title="Phone" name="phone" label="(11) 99711-7911" type="tel" />

              <HookFormTextInput title="Street" name="street" label="Street" type="text" disabled />
              <HookFormTextInput title="Neighborhood" name="neighborhood" label="Neighborhood" type="text" disabled />
              <HookFormTextInput title="Zip" name="zip" label="Zip" type="text" disabled />
            </CardContent>
            <CardFooter className="flex flex-col justify-end">
              <Button
                className="w-full"
                type="submit"
                disabled={isSubmitting || !isDirty || updateProfile.isPending}
              >
                {isSubmitting || updateProfile.isPending ? "Saving..." : "Save"}
              </Button>
            </CardFooter>
          </Card>
        </form>
      </FormProvider>
    </div>
  )
}
