import { apiCall } from '@/lib/api-client'
import type {
  GetContactsResponse,
  Contact,
  CreateContactInput,
  GetContactFrequencyResponse,
} from '@/types'

export const ContactService = {
  getAll: async (params?: URLSearchParams) => {
    const query = params ? `?${params.toString()}` : ''
    return apiCall<GetContactsResponse>(`/contacts${query}`)
  },

  getById: async (id: string) => {
    return apiCall<Contact>(`/contacts/${id}`)
  },

  create: async (data: CreateContactInput) => {
    return apiCall<Contact>('/contacts', {
      method: 'POST',
      body: JSON.stringify(data),
    })
  },

  update: async (id: string, data: Partial<CreateContactInput>) => {
    return apiCall<Contact>(`/contacts/${id}`, {
      method: 'PATCH',
      body: JSON.stringify(data),
    })
  },

  delete: async (id: string) => {
    return apiCall<void>(`/contacts/${id}`, {
      method: 'DELETE',
    })
  },

  getFrequency: async (limit = 5) => {
    return apiCall<GetContactFrequencyResponse>(`/contacts/frequency?limit=${limit}`)
  },
}
