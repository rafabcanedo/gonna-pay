import { apiCall } from '@/lib/api-client'
import type { GetGroupsResponse, Group, GroupDetail, CreateGroupInput, AddMemberInput, UpdateGroupInput } from '@/types'

export const GroupService = {
  getAll: async (params?: URLSearchParams) => {
    const query = params ? `?${params.toString()}` : ''
    return apiCall<GetGroupsResponse>(`/groups${query}`)
  },

  getById: async (id: string) => {
    return apiCall<GroupDetail>(`/groups/${id}`)
  },

  create: async (data: CreateGroupInput) => {
    return apiCall<Group>('/groups', {
      method: 'POST',
      body: JSON.stringify(data),
    })
  },

  update: async (id: string, data: UpdateGroupInput) => {
    return apiCall<Group>(`/groups/${id}`, {
      method: 'PATCH',
      body: JSON.stringify(data),
    })
  },

  delete: async (id: string) => {
    return apiCall<void>(`/groups/${id}`, {
      method: 'DELETE',
    })
  },

  addMember: async (groupId: string, data: AddMemberInput) => {
    return apiCall<void>(`/groups/${groupId}/member/${data.contactId}`, {
      method: 'POST',
    })
  },

  removeMember: async (groupId: string, contactId: string) => {
    return apiCall<void>(`/groups/${groupId}/member/${contactId}`, {
      method: 'DELETE',
    })
  },
}
