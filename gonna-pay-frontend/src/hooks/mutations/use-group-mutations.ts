'use client'

import { useMutation, useQueryClient } from '@tanstack/react-query'
import { GroupService } from '@/services/group.service'
import { toast } from 'sonner'
import type { CreateGroupInput, AddMemberInput, UpdateGroupInput, Group } from '@/types'
import { ApiError } from '@/lib/errors/api.error'

export function useCreateGroup() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (data: CreateGroupInput) => GroupService.create(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['groups'] })
      toast.success('Group created successfully!')
    },
    onError: (error: ApiError) => {
      toast.error(error.message)
    },
  })
}

export function useUpdateGroup() {
  const queryClient = useQueryClient()

  return useMutation<Group, ApiError, { id: string; data: UpdateGroupInput}>({
    mutationFn: ({ id, data }) => GroupService.update(id, data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['groups']})
      toast.success('Group updated successfully!')
    },
    onError: (error: ApiError) => {
      toast.error(error.message)
    },
  })
}

export function useDeleteGroup() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (id: string) => GroupService.delete(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['groups'] })
      toast.success('Group deleted successfully!')
    },
    onError: (error: ApiError) => {
      toast.error(error.message)
    },
  })
}

export function useAddMember() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: ({ groupId, data }: { groupId: string; data: AddMemberInput }) =>
      GroupService.addMember(groupId, data),
    onSuccess: (_, { groupId }) => {
      queryClient.invalidateQueries({ queryKey: ['groups', groupId] })
      toast.success('Member added successfully!')
    },
    onError: (error: ApiError) => {
      toast.error(error.message)
    },
  })
}

export function useRemoveMember() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: ({ groupId, contactId }: { groupId: string; contactId: string }) =>
      GroupService.removeMember(groupId, contactId),
    onSuccess: (_, { groupId }) => {
      queryClient.invalidateQueries({ queryKey: ['groups', groupId] })
      toast.success('Member removed successfully!')
    },
    onError: (error: ApiError) => {
      toast.error(error.message)
    },
  })
}
