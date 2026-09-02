"use client";

import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";

import { useContactsQuery } from "@/hooks/queries/use-contact-query";
import { useCostsWithSplitsQuery } from "@/hooks/queries/use-cost-query";

export const TablePayments = () => {
  const { data: contacts = [] } = useContactsQuery()
  const { data: costsWithSplits = [] } = useCostsWithSplitsQuery()

  const contactsMap = new Map(contacts.map((c) => [c.id, c]))

  const rows = costsWithSplits.flatMap((cost) =>
    cost.splits.map((split) => ({
      name: split.contactName,
      group: cost.groupName ?? '—',
      value: split.value,
      category: contactsMap.get(split.contactId)?.category ?? '—',
    }))
  )

  return (
    <div className="w-full flex justify-center mt-6 mb-6">
      <div className="w-full max-w-7xl rounded-xl border bg-white shadow-sm">
        <div className="flex items-center justify-between px-6 pt-4 pb-2">
          <h2 className="text-lg">Contacts</h2>
          <span className="text-xs">{rows.length} contacts</span>
        </div>

        <div className="px-2 pb-2">
          <Table className="w-full">
            <TableHeader>
              <TableRow>
                <TableHead className="w-[100px]">Name</TableHead>
                <TableHead>Group</TableHead>
                <TableHead>Value</TableHead>
                <TableHead className="text-right">Category</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {rows.map((row, index) => (
                <TableRow key={index}>
                  <TableCell className="font-medium">{row.name}</TableCell>
                  <TableCell>{row.group}</TableCell>
                  <TableCell>{row.value.toLocaleString("en-US", { style: "currency", currency: "USD" })}</TableCell>
                  <TableCell>{row.category}</TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </div>
      </div>
    </div>
  );
};
