import type * as yup from 'yup'
import type { editCostSchema } from '@/validations/schemas'

export type CostDetailsForm = yup.InferType<typeof editCostSchema>
