import type * as yup from 'yup'
import type { editContactSchema } from '@/validations/schemas'

export type EditContactForm = yup.InferType<typeof editContactSchema>
