import type * as yup from 'yup'
import type { editGroupSchema } from '@/validations/schemas'
import type { GroupDetail } from '@/types'

export type EditGroupForm = yup.InferType<typeof editGroupSchema>

export interface IEditGroupProps {
  group: GroupDetail
  open: boolean
  onOpenChange: (open: boolean) => void
}
