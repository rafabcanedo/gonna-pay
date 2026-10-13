'use client'

import { useEffect } from 'react'
import { useRouter } from 'next/navigation'
import { FormProvider, useForm, type SubmitHandler } from 'react-hook-form'
import { yupResolver } from '@hookform/resolvers/yup'
import { ArrowLeft } from 'lucide-react'
import { Card, CardContent, CardFooter, CardHeader, CardTitle } from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { HookFormTextInput } from '@/components/hook-form-text-input'
import { HookFormSelect } from '@/components/hook-form-select'
import { CONTACT_CATEGORIES } from '@/app/(dashboard)/my-contacts/constants'
import { editContactSchema } from '@/validations/schemas'
import { useQueryClient } from '@tanstack/react-query'
import { useContactQuery } from '@/hooks/queries/use-contact-query'
import { useUpdateContact } from '@/hooks/mutations/use-contact-mutations'
import { HookFormPhoneInput, formatPhone } from '@/components/hook-form-phone-input'
import type { Contact, ContactCategory } from '@/types'
import { IPropsContactDetails } from './interfaces'
import type { EditContactForm } from './types'

export const ContactDetails = ({ contactId }: IPropsContactDetails) => {
  const router = useRouter()
  const queryClient = useQueryClient()
  const { data: contact } = useContactQuery(contactId)
  const { mutateAsync: updateContact, isPending } = useUpdateContact()

  const getDefaultValues = () => {
    const cached = queryClient.getQueryData<Contact>(['contacts', contactId])
    return {
      name: cached?.name ?? '',
      email: cached?.email ?? '',
      phone: formatPhone(cached?.phone ?? ''),
      category: (cached?.category ?? '') as ContactCategory,
    }
  }

  const methods = useForm<EditContactForm>({
    resolver: yupResolver<EditContactForm, object, EditContactForm>(editContactSchema),
    defaultValues: getDefaultValues(),
    mode: 'onTouched',
  })

  const { handleSubmit, formState, reset } = methods
  const { isSubmitting, isDirty } = formState

  useEffect(() => {
    if (contact) {
      reset({
        name: contact.name,
        email: contact.email,
        phone: formatPhone(contact.phone ?? ''),
        category: contact.category as ContactCategory,
      })
    }
  }, [contact, reset])

  const handleSubmitContact: SubmitHandler<EditContactForm> = async (data) => {
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
              <HookFormPhoneInput title="Phone" name="phone" label="(11) 99711-7911" type="tel" />
              <HookFormSelect
                name="category"
                label="Category"
                placeholder="Select a category"
                options={CONTACT_CATEGORIES}
                groupLabel="Category"
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
