'use client'

import { useEffect } from 'react'
import { useRouter } from 'next/navigation'
import { FormProvider, useForm, type SubmitHandler } from 'react-hook-form'
import { yupResolver } from '@hookform/resolvers/yup'
import { ArrowLeft } from 'lucide-react'
import { Card, CardContent, CardFooter, CardHeader, CardTitle } from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { HookFormTextInput } from '@/components/hook-form-text-input'
import { SelectCategory } from '@/app/(dashboard)/my-contacts/components/select-category'
import { addContactSchema } from '@/validations/schemas'
import { useContactQuery } from '@/hooks/queries/use-contact-query'
import { useUpdateContact } from '@/hooks/mutations/use-contact-mutations'
import type { ContactCategory, CreateContactInput } from '@/types'

interface IPropsContactDetails {
  contactId: string
}

export const ContactDetails = ({ contactId }: IPropsContactDetails) => {
  const router = useRouter()
  const { data: contact } = useContactQuery(contactId)
  const { mutateAsync: updateContact, isPending } = useUpdateContact()

  const methods = useForm<CreateContactInput>({
    resolver: yupResolver(addContactSchema),
    defaultValues: { name: '', email: '', phone: '', category: undefined },
    mode: 'onChange',
  })

  const { handleSubmit, formState, reset, watch, setValue } = methods
  const { isSubmitting, isDirty } = formState

  useEffect(() => {
    if (contact) {
      reset({
        name: contact.name,
        email: contact.email,
        phone: contact.phone,
        category: contact.category as ContactCategory,
      })
    }
  }, [contact, reset])

  const handleSubmitContact: SubmitHandler<CreateContactInput> = async (data) => {
    await updateContact({ id: contactId, data })
    reset(data)
  }

  return (
    <div className="flex flex-col gap-6 px-8 w-full mt-8">
      <header className="flex flex-row items-center gap-3">
        <Button variant="ghost" size="icon" onClick={() => router.back()}>
          <ArrowLeft className="h-5 w-5" />
        </Button>
        <h1 className="font-montserrat text-xl text-zinc-600">
          {contact?.name ?? '...'}
        </h1>
      </header>

      <FormProvider {...methods}>
        <form onSubmit={handleSubmit(handleSubmitContact)}>
          <Card className="w-full max-w-xl">
            <CardHeader>
              <CardTitle>Contact details</CardTitle>
            </CardHeader>
            <CardContent className="flex flex-col gap-4">
              <HookFormTextInput title="Name" name="name" label="John Jason" type="text" />
              <HookFormTextInput title="Email" name="email" label="john@example.com" type="email" />
              <HookFormTextInput title="Phone" name="phone" label="+55 11 997117911" type="text" />
              <SelectCategory
                value={watch('category')}
                onValueChange={(value: ContactCategory) =>
                  setValue('category', value, { shouldValidate: true, shouldDirty: true })
                }
              />
            </CardContent>
            <CardFooter>
              <Button
                className="w-full"
                type="submit"
                disabled={isSubmitting || !isDirty || isPending}
              >
                {isSubmitting || isPending ? 'Saving...' : 'Save changes'}
              </Button>
            </CardFooter>
          </Card>
        </form>
      </FormProvider>
    </div>
  )
}
