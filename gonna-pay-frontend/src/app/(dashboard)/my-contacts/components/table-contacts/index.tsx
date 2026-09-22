"use client"

import { useState, useEffect } from "react"
import { useRouter } from "next/navigation"
import { MoreHorizontal } from "lucide-react"
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table"
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
} from "@/components/ui/alert-dialog"
import { Button } from "@/components/ui/button"
import { BadgeType } from "@/utils/badge-types"
import { TableEmptyState } from "@/components/table-empty-state"
import { useContactsQuery } from "@/hooks/queries/use-contact-query"
import type { ContactQueryFilters } from "@/hooks/queries/types"
import { useDeleteContact } from "@/hooks/mutations/use-contact-mutations"
import { TablePagination } from "@/components/table-pagination"
import { useDebounce } from "@/hooks/use-debounce"
import { ContactsFilter } from "./components/contacts-filter"
import type { Contact } from "@/types"

export const TableContact = () => {
  const router = useRouter()
  const [page, setPage] = useState(1)
  const [category, setCategory] = useState('')
  const [search, setSearch] = useState('')
  const debouncedSearch = useDebounce(search)

  const filters: ContactQueryFilters = {
    ...(category ? { category } : {}),
    ...(debouncedSearch ? { search: debouncedSearch } : {}),
  }

  useEffect(() => { setPage(1) }, [category, debouncedSearch])

  const { data } = useContactsQuery(page, 20, filters)
  const { mutate: deleteContact, isPending } = useDeleteContact()

  const [contactToDelete, setContactToDelete] = useState<Contact | null>(null)

  const contacts = data?.data ?? []
  const total = data?.total ?? 0
  const totalPages = data?.totalPages ?? 1

  return (
    <>
      <div className="w-full flex justify-center mt-6 mb-6">
        <div className="w-full max-w-7xl rounded-xl border bg-white shadow-sm">
          <div className="flex items-center justify-between px-6 pt-4 pb-2">
            <h2 className="text-lg">Contacts</h2>
            <div className="flex items-center gap-2">
              <ContactsFilter
                category={category}
                search={search}
                onCategoryChange={setCategory}
                onSearchChange={setSearch}
              />
              <span className="text-xs">{total} contacts</span>
            </div>
          </div>

          <div className="px-2 pb-2">
            <Table className="w-full">
              <TableHeader>
                <TableRow>
                  <TableHead className="w-[100px]">Name</TableHead>
                  <TableHead>Email</TableHead>
                  <TableHead>Phone</TableHead>
                  <TableHead>Category</TableHead>
                  <TableHead className="w-[50px]" />
                </TableRow>
              </TableHeader>
              <TableBody>
                {contacts.length > 0 ? (
                  contacts.map((contact) => (
                    <TableRow key={contact.id}>
                      <TableCell className="font-medium">{contact.name}</TableCell>
                      <TableCell>{contact.email}</TableCell>
                      <TableCell>{contact.phone}</TableCell>
                      <TableCell>
                        <BadgeType type={contact.category} />
                      </TableCell>
                      <TableCell>
                        <DropdownMenu>
                          <DropdownMenuTrigger asChild>
                            <Button variant="ghost" className="h-8 w-8 p-0">
                              <span className="sr-only">Open menu</span>
                              <MoreHorizontal className="h-4 w-4" />
                            </Button>
                          </DropdownMenuTrigger>
                          <DropdownMenuContent align="end">
                            <DropdownMenuItem
                              onClick={() => router.push(`/my-contacts/details/${contact.id}?name=${encodeURIComponent(contact.name)}`)}
                            >
                              Edit
                            </DropdownMenuItem>
                            <DropdownMenuItem
                              className="text-red-600 focus:text-red-600"
                              onClick={() => setContactToDelete(contact)}
                            >
                              Delete
                            </DropdownMenuItem>
                          </DropdownMenuContent>
                        </DropdownMenu>
                      </TableCell>
                    </TableRow>
                  ))
                ) : (
                  <TableEmptyState colSpan={5} message="No contacts found" />
                )}
              </TableBody>
            </Table>
          </div>

          {totalPages > 1 && (
            <div className="px-6 pb-4">
              <TablePagination page={page} totalPages={totalPages} onPageChange={setPage} />
            </div>
          )}
        </div>
      </div>

      <AlertDialog open={!!contactToDelete} onOpenChange={(open) => { if (!open) setContactToDelete(null) }}>
        <AlertDialogContent>
          <AlertDialogHeader>
            <AlertDialogTitle>Delete contact</AlertDialogTitle>
            <AlertDialogDescription>
              Are you sure you want to delete <strong>{contactToDelete?.name}</strong>? This action cannot be undone.
            </AlertDialogDescription>
          </AlertDialogHeader>
          <AlertDialogFooter>
            <AlertDialogCancel disabled={isPending}>Cancel</AlertDialogCancel>
            <AlertDialogAction
              className="bg-red-600 hover:bg-red-700"
              disabled={isPending}
              onClick={() => {
                if (contactToDelete) {
                  deleteContact(contactToDelete.id, { onSuccess: () => setContactToDelete(null) })
                }
              }}
            >
              Delete
            </AlertDialogAction>
          </AlertDialogFooter>
        </AlertDialogContent>
      </AlertDialog>
    </>
  )
}
