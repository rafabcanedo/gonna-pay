'use client'

import { useEffect } from 'react'
import { FormProvider, useForm, type SubmitHandler } from 'react-hook-form'
import { yupResolver } from '@hookform/resolvers/yup'
import { Dialog, DialogContent, DialogHeader, DialogTitle } from '@/components/ui/dialog'
import { Button } from '@/components/ui/button'
import { HookFormTextInput } from '@/components/hook-form-text-input'
import { HookFormSelect } from '@/components/hook-form-select'
import { GROUP_CATEGORY_FORM_OPTIONS } from '@/app/(dashboard)/groups/constants'
import { editGroupSchema } from '@/validations/schemas'
import { useUpdateGroup } from '@/hooks/mutations/use-group-mutations'
import type { TransactionCategory } from '@/types'
import type { IEditGroupProps, EditGroupForm } from './types'

export const EditGroup = ({ group, open, onOpenChange }: IEditGroupProps) => {
  const { mutateAsync: updateGroup, isPending } = useUpdateGroup()

  const methods = useForm<EditGroupForm>({
    resolver: yupResolver<EditGroupForm, object, EditGroupForm>(editGroupSchema),
    defaultValues: {
      name: group.name,
      category: group.category as TransactionCategory,
    },
    mode: 'onTouched',
  })

  const { handleSubmit, formState, reset } = methods
  const { isSubmitting, isDirty } = formState

  useEffect(() => {
    if (open) {
      reset({
        name: group.name,
        category: group.category as TransactionCategory,
      })
    }
  }, [open, group, reset])

  const handleSubmitGroup: SubmitHandler<EditGroupForm> = async (data) => {
    await updateGroup({ id: group.id, data })
    onOpenChange(false)
  }

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Edit Group</DialogTitle>
        </DialogHeader>
        <FormProvider {...methods}>
          <form onSubmit={handleSubmit(handleSubmitGroup)} className="flex flex-col gap-4">
            <HookFormTextInput title="Name" name="name" label="Trip to Paris" type="text" />
            <HookFormSelect
              name="category"
              label="Category"
              placeholder="Select a category"
              options={GROUP_CATEGORY_FORM_OPTIONS}
              groupLabel="Category"
            />
            <Button
              type="submit"
              disabled={isSubmitting || !isDirty || isPending}
            >
              {isSubmitting || isPending ? 'Saving...' : 'Save changes'}
            </Button>
          </form>
        </FormProvider>
      </DialogContent>
    </Dialog>
  )
}
