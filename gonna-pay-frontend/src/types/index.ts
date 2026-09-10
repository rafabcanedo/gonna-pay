export enum ContactCategory {
  FAMILY = "Family",
  FRIEND = "Friend",
  WORK = "Work",
}

export enum TransactionCategory {
  DINNER = "Dinner",
  LUNCH = "Lunch",
  ENTERTAINMENT = "Entertainment",
  TRAVEL = "Travel",
  OTHERS = "Others",
}

export type Contact = {
  id: string
  name: string
  email: string
  phone: string
  category: ContactCategory
}

export type GetContactsResponse = Contact[]

export type ContactFrequency = {
  contactId: string
  contactName: string
  sharedCosts: number
}

export type GetContactFrequencyResponse = ContactFrequency[]

export type CreateContactInput = Omit<Contact, 'id'>

export type GroupMember = {
  id: string
  name: string
  email: string
}

export type Group = {
  id: string
  name: string
  category: TransactionCategory
  createdAt: string
  updatedAt: string
}

export type GroupDetail = Group & {
  members?: GroupMember[]
}

export type GetGroupsResponse = {
  groups: GroupDetail[]
  total: number
}

export type CreateGroupInput = {
  name: string
  category: TransactionCategory
  memberIds?: string[]
}

export type UpdateGroupInput = {
  name?: string
  category?: TransactionCategory
}

export type AddMemberInput = {
  contactId: string
}

export type CostSplit = {
  id: string
  contactId: string
  contactName: string
  value: number
  percentage: number
}

export type Cost = {
  id: string
  costName: string
  totalValue: number
  ownerPercentage: number
  ownerValue: number
  category: TransactionCategory
  groupId?: string
  groupName?: string
  splitCount: number
  createdAt: string
}

export type CostDetail = Cost & {
  updatedAt: string
  splits: CostSplit[]
}

export type GetCostsResponse = Cost[]

export type CreateCostInput = {
  groupId?: string
  costName: string
  category: TransactionCategory
  totalValue: number
  ownerPercentage?: number
}

export type UpdateCostInput = {
  costName: string
  category: TransactionCategory
  totalValue: number
  ownerPercentage?: number
}

export interface IRestError {
  message: string
  error: string
  code: number
  causes?: Array<{
    field: string
    message: string
  }>
}
