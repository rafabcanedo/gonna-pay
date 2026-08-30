'use client'

import { useState } from 'react'
import { useRouter } from 'next/navigation'
import { useGroupQuery } from '@/hooks/queries/use-group-query'
import { useDeleteGroup, useRemoveMember } from '@/hooks/mutations/use-group-mutations'
import { BadgeType } from '@/utils/badge-types'
import { AddMember } from '../add-member'
import { Button } from '@/components/ui/button'
import { ChevronLeft, Loader2, X } from 'lucide-react';
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
} from '@/components/ui/alert-dialog'

const getInitials = (name: string) => {
  const parts = name.trim().split(' ')
  if (parts.length === 1) return parts[0][0].toUpperCase()
  return (parts[0][0] + parts[parts.length - 1][0]).toUpperCase()
}

import { IGroupDetailProps } from './interfaces'

export const GroupDetails = ({ groupId }: IGroupDetailProps) => {
  const router = useRouter()
  const { data: group } = useGroupQuery(groupId)
  const { mutateAsync: deleteGroup, isPending } = useDeleteGroup()
  const { mutateAsync: removeMember } = useRemoveMember()
  const [hoveredMemberId, setHoveredMemberId] = useState<string | null>(null)
  const [removingMemberId, setRemovingMemberId] = useState<string | null>(null)

  if (!group) return null

  const handleBack = () => router.push('/groups')

  const handleDelete = async () => {
    await deleteGroup(groupId)
    handleBack()
  }

  const handleRemoveMember = async (contactId: string) => {
    setRemovingMemberId(contactId)
    await removeMember({ groupId, contactId })
    setRemovingMemberId(null)
  }

  return (
    <div className="flex flex-col gap-6 px-8 w-full">
      <div>
        <ChevronLeft className='w-8 h-8 text-stone-300 cursor-pointer' onClick={handleBack} />
      </div>
      <header className="flex flex-row items-center gap-3 h-12 mt-4">
        <h1 className="font-montserrat text-xl text-zinc-700 font-semibold">{group.name}</h1>
        <BadgeType type={group.category} />
      </header>

      <div className="flex flex-row items-center gap-4">
        <div className="flex flex-row flex-wrap items-center gap-2 border border-zinc-200 rounded-lg px-3 py-2 min-h-[52px]">
          {(group.members ?? []).map((member) => {
            const isHovered = hoveredMemberId === member.id
            const isRemoving = removingMemberId === member.id

            return (
              <button
                key={member.id}
                type="button"
                title={member.name}
                onClick={() => handleRemoveMember(member.id)}
                onMouseEnter={() => setHoveredMemberId(member.id)}
                onMouseLeave={() => setHoveredMemberId(null)}
                disabled={!!removingMemberId}
                className={`flex items-center justify-center w-9 h-9 rounded-full border font-poppins text-sm font-medium select-none transition-colors cursor-pointer
                  ${isHovered || isRemoving
                    ? 'bg-red-100 border-red-300 text-red-500'
                    : 'bg-zinc-100 border-zinc-300 text-zinc-600'
                  }`}
              >
                {isRemoving
                  ? <Loader2 className="w-4 h-4 animate-spin" />
                  : isHovered
                    ? <X className="w-4 h-4" />
                    : getInitials(member.name)
                }
              </button>
            )
          })}
        </div>
        <AddMember groupId={groupId} currentMembers={group.members} />
      </div>

      <AlertDialog>
        <AlertDialogTrigger asChild>
          <Button
            variant="outline"
            className="w-40 text-red-400 border-red-200 hover:bg-red-50 hover:text-red-500"
          >
            Delete Group
          </Button>
        </AlertDialogTrigger>
        <AlertDialogContent>
          <AlertDialogHeader>
            <AlertDialogTitle>Are you sure you want to delete this group?</AlertDialogTitle>
            <AlertDialogDescription>
              This action cannot be undone. All group data will be permanently removed.
            </AlertDialogDescription>
          </AlertDialogHeader>
          <AlertDialogFooter>
            <AlertDialogCancel>Cancel</AlertDialogCancel>
            <AlertDialogAction
              className="bg-red-600 hover:bg-red-700"
              disabled={isPending}
              onClick={handleDelete}
            >
              Delete
            </AlertDialogAction>
          </AlertDialogFooter>
        </AlertDialogContent>
      </AlertDialog>
    </div>
  )
}
