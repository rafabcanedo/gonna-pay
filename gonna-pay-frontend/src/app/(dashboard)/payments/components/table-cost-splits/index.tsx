"use client";

import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { BadgeType } from "@/utils/badge-types";
import { TableEmptyState } from "@/components/table-empty-state";
import { useContactsQuery } from "@/hooks/queries/use-contact-query";
import { useCostsWithSplitsQuery } from "@/hooks/queries/use-cost-query";

export const TableCostSplits = () => {
  const { data: contacts = [] } = useContactsQuery();
  const { data: costsWithSplits = [] } = useCostsWithSplitsQuery();

  const contactsMap = new Map(contacts.map((c) => [c.id, c]));

  const rows = costsWithSplits.flatMap((cost) =>
    cost.splits.map((split) => ({
      id: `${cost.id}-${split.id}`,
      name: split.contactName,
      group: cost.groupName ?? "—",
      value: split.value,
      category: contactsMap.get(split.contactId)?.category,
    }))
  );

  return (
    <div className="w-full flex justify-center mt-6 mb-6">
      <div className="w-full max-w-7xl rounded-xl border bg-white shadow-sm">
        <div className="flex items-center justify-between px-6 pt-4 pb-2">
          <h2 className="text-lg">Shared costs</h2>
          <span className="text-xs">{rows.length} splits</span>
        </div>

        <div className="px-2 pb-2">
          <Table className="w-full">
            <TableHeader>
              <TableRow>
                <TableHead>Name</TableHead>
                <TableHead>Group</TableHead>
                <TableHead>Value</TableHead>
                <TableHead className="text-right">Category</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {rows.length > 0 ? (
                rows.map((row) => (
                  <TableRow key={row.id}>
                    <TableCell className="font-medium text-zinc-800">{row.name}</TableCell>
                    <TableCell>{row.group}</TableCell>
                    <TableCell>
                      {row.value.toLocaleString("en-US", { style: "currency", currency: "USD" })}
                    </TableCell>
                    <TableCell className="text-right">
                      {row.category ? <BadgeType type={row.category} /> : "—"}
                    </TableCell>
                  </TableRow>
                ))
              ) : (
                <TableEmptyState colSpan={4} message="No shared costs found" />
              )}
            </TableBody>
          </Table>
        </div>
      </div>
    </div>
  );
};
