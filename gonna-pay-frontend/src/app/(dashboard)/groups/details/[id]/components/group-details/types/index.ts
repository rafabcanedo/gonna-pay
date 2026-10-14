import type * as yup from 'yup'
import type { editGroupSchema } from '@/validations/schemas'

export type EditGroupForm = yup.InferType<typeof editGroupSchema>

export interface IGroupDetailProps {
  groupId: string
}
