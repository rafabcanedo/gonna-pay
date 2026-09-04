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
import { useCostsQuery } from "@/hooks/queries/use-cost-query";

export const TablePayments = () => {
  const { data: costs = [] } = useCostsQuery();

  const soloPayments = costs.filter((cost) => cost.splitCount === 0);

  return (
    <div className="w-full flex justify-center mt-6 mb-6">
      <div className="w-full max-w-7xl rounded-xl border bg-white shadow-sm">
        <div className="flex items-center justify-between px-6 pt-4 pb-2">
          <h2 className="text-lg">Payments</h2>
          <span className="text-xs">{soloPayments.length} payments</span>
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
              {soloPayments.length > 0 ? (
                soloPayments.map((cost) => (
                  <TableRow key={cost.id}>
                    <TableCell className="font-medium text-zinc-800">{cost.costName}</TableCell>
                    <TableCell>{cost.groupName ?? "—"}</TableCell>
                    <TableCell>
                      {cost.totalValue.toLocaleString("en-US", { style: "currency", currency: "USD" })}
                    </TableCell>
                    <TableCell className="text-right">
                      <BadgeType type={cost.category} />
                    </TableCell>
                  </TableRow>
                ))
              ) : (
                <TableEmptyState colSpan={4} message="No payments found" />
              )}
            </TableBody>
          </Table>
        </div>
      </div>
    </div>
  );
};
