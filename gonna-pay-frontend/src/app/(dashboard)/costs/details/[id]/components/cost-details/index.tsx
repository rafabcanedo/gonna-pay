'use client'

import { useEffect } from 'react'
import { useRouter } from 'next/navigation'
import { useQueryClient } from '@tanstack/react-query'
import { FormProvider, useForm, type SubmitHandler } from 'react-hook-form'
import { yupResolver } from '@hookform/resolvers/yup'
import { ArrowLeft } from 'lucide-react'
import { Card, CardContent, CardFooter, CardHeader, CardTitle } from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { HookFormTextInput } from '@/components/hook-form-text-input'
import { HookFormSelect } from '@/components/hook-form-select'
import { editCostSchema } from '@/validations/schemas'
import { useCostQuery } from '@/hooks/queries/use-cost-query'
import { useUpdateCost } from '@/hooks/mutations/use-cost-mutations'
import { COST_CATEGORY_FORM_OPTIONS } from '@/app/(dashboard)/costs/constants'
import { TransactionCategory, type CostDetail } from '@/types'
import type { IPropsCostDetails } from './interfaces'
import type { CostDetailsForm } from './types'

export const CostDetails = ({ costId }: IPropsCostDetails) => {
  const router = useRouter()
  const queryClient = useQueryClient()
  const { data: cost } = useCostQuery(costId)
  const { mutateAsync: updateCost, isPending } = useUpdateCost()

  const getDefaultValues = () => {
    const cached = queryClient.getQueryData<CostDetail>(['costs', costId])
    return {
      costName: cached?.costName ?? '',
      totalValue: cached ? String(cached.totalValue) : '',
      category: (cached?.category ?? '') as TransactionCategory,
      ownerPercentage: cached ? String(cached.ownerPercentage) : '',
    }
  }

  const methods = useForm<CostDetailsForm>({
    resolver: yupResolver<CostDetailsForm, object, CostDetailsForm>(editCostSchema),
    defaultValues: getDefaultValues(),
    mode: 'onTouched',
  })

  const { handleSubmit, formState, reset } = methods
  const { isSubmitting, isDirty } = formState

  useEffect(() => {
    if (cost) {
      reset({
        costName: cost.costName,
        totalValue: String(cost.totalValue),
        category: cost.category,
        ownerPercentage: String(cost.ownerPercentage),
      })
    }
  }, [cost, reset])

  const handleSubmitCost: SubmitHandler<CostDetailsForm> = async (data) => {
    await updateCost({
      id: costId,
      data: {
        ...(data.costName ? { costName: data.costName } : {}),
        ...(data.totalValue ? { totalValue: Number(data.totalValue) } : {}),
        ...(data.category ? { category: data.category } : {}),
        ...(data.ownerPercentage ? { ownerPercentage: Number(data.ownerPercentage) } : {}),
      },
    })
    reset(data)
  }

  return (
    <div className="flex flex-col gap-6 px-8 w-full mt-8">
      <header className="flex flex-row items-center gap-3">
        <Button variant="ghost" size="icon" onClick={() => router.back()}>
          <ArrowLeft className="h-5 w-5" />
        </Button>
        <h1 className="font-montserrat text-xl text-zinc-600">
          {cost?.costName ?? '...'}
        </h1>
      </header>

      <FormProvider {...methods}>
        <form onSubmit={handleSubmit(handleSubmitCost)}>
          <Card className="w-full max-w-xl">
            <CardHeader>
              <CardTitle>Cost details</CardTitle>
            </CardHeader>
            <CardContent className="flex flex-col gap-4">
              <HookFormTextInput title="Name" name="costName" label="Dinner at restaurant" type="text" />
              <HookFormTextInput title="Total Value" name="totalValue" label="100.50" type="text" />
              <HookFormSelect
                name="category"
                label="Category"
                placeholder="Select a category"
                groupLabel="Categories"
                options={COST_CATEGORY_FORM_OPTIONS}
              />
              <HookFormTextInput title="Your percentage (optional)" name="ownerPercentage" label="50" type="number" />
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
