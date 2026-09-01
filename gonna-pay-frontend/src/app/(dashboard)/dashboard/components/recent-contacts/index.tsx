import {
    Table,
    TableBody,
    TableCell,
    TableHead,
    TableHeader,
    TableRow,
} from "@/components/ui/table"
import { BadgeType } from "@/utils/badge-types"
import Link from "next/link"
import { RecentContactsProps } from '../interfaces'

const LIMIT = 6

export const RecentContacts = ({ contacts }: RecentContactsProps) => {

    const limitedContacts = contacts.slice(0, LIMIT)

    return (
        <div className="mt-6 mb-6">
            <div className="w-full max-w-5xl rounded-xl border border-zinc-200 bg-white shadow-sm">
                <div className="flex items-center justify-between px-6 pt-4 pb-2">
                    <h2 className="font-mono text-lg font-semibold text-zinc-600">
                        Recent Contacts
                    </h2>
                </div>

                <div className="px-2 pb-2">
                    <Table className="w-full">
                        <TableHeader>
                            <TableRow>
                                <TableHead className="w-[160px]">Name</TableHead>
                                <TableHead>Email</TableHead>
                                <TableHead className="w-[140px]">Phone</TableHead>
                                <TableHead className="text-right w-[140px]">Category</TableHead>
                            </TableRow>
                        </TableHeader>
                        <TableBody>
                            {limitedContacts.map((contact) => (
                                <TableRow key={contact.id}>
                                    <TableCell className="font-medium text-zinc-800">
                                        {contact.name}
                                    </TableCell>

                                    <TableCell className="text-sm text-zinc-700">
                                        {contact.email}
                                    </TableCell>

                                    <TableCell className="text-sm text-zinc-700">
                                        {contact.phone}
                                    </TableCell>

                                    <TableCell className="text-right text-sm text-zinc-500">
                                        <BadgeType type={contact.category} />
                                    </TableCell>
                                </TableRow>
                            ))}
                        </TableBody>
                    </Table>
                </div>

                <div className="flex justify-end px-6 py-3 border-t border-zinc-100">
                    <Link href="/my-contacts">
                        <button
                            type="button"
                            className="text-sm font-medium text-zinc-600 hover:text-zinc-500 hover:underline cursor-pointer"
                        >
                            See more
                        </button>
                    </Link>
                </div>
            </div>
        </div>
    )
}