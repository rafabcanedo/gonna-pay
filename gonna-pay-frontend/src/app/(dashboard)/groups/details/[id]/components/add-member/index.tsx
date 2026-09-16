'use client'

import { useState } from 'react'
import { Loader2 } from 'lucide-react'
import { Button } from '@/components/ui/button'
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@/components/ui/dialog'
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectLabel,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select'
import { useContactsQuery } from '@/hooks/queries/use-contact-query'
import { useAddMember } from '@/hooks/mutations/use-group-mutations'
import { IAddMemberProps } from './interfaces'

export const AddMember = ({ groupId, currentMembers = [] }: IAddMemberProps) => {
  const [open, setOpen] = useState(false)
  const [selectedContactId, setSelectedContactId] = useState('')

  const { data: contacts } = useContactsQuery()
  const { mutateAsync: addMember, isPending } = useAddMember()

  const currentMemberIds = new Set(currentMembers.map((m) => m.id))
  const availableContacts = (contacts?.data ?? []).filter((c) => !currentMemberIds.has(c.id))

  const handleSubmit = async () => {
    if (!selectedContactId) return
    await addMember({ groupId, data: { contactId: selectedContactId } })
    setSelectedContactId('')
    setOpen(false)
  }

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild>
        <Button variant="outline">Add Member</Button>
      </DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Add Member</DialogTitle>
        </DialogHeader>
        <div className="flex flex-col gap-4 py-4">
          <Select value={selectedContactId} onValueChange={setSelectedContactId}>
            <SelectTrigger>
              <SelectValue placeholder="Select a contact" />
            </SelectTrigger>
            <SelectContent>
              <SelectGroup>
                <SelectLabel>Contacts</SelectLabel>
                {availableContacts.map((contact) => (
                  <SelectItem key={contact.id} value={contact.id}>
                    {contact.name}
                  </SelectItem>
                ))}
              </SelectGroup>
            </SelectContent>
          </Select>
          <div className="flex flex-row gap-2">
            <Button
              variant="outline"
              className="w-1/2"
              onClick={() => setOpen(false)}
              disabled={isPending}
            >
              Cancel
            </Button>
            <Button
              className="w-1/2"
              onClick={handleSubmit}
              disabled={!selectedContactId || isPending}
            >
              {isPending ? <Loader2 className="w-4 h-4 animate-spin" /> : 'Add'}
            </Button>
          </div>
        </div>
      </DialogContent>
    </Dialog>
  )
}
