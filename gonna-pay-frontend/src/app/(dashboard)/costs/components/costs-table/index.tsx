"use client";

import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { useCostsQuery } from "@/hooks/queries/use-cost-query";
import { BadgeType } from "@/utils/badge-types";

export const CostsTable = () => {
  const { data } = useCostsQuery();

  const costs = data ?? [];
  const total = data?.length ?? 0;

  return (
    <div className="w-full flex justify-center mt-6 mb-6">
      <div className="w-full max-w-7xl rounded-xl border bg-white shadow-sm">
        <div className="flex items-center justify-between px-6 pt-4 pb-2">
          <h2 className="text-lg">Costs</h2>
          <span className="text-xs">{total} costs</span>
        </div>

        <div className="px-2 pb-2">
          <Table className="w-full">
            <TableHeader>
              <TableRow>
                <TableHead>Name</TableHead>
                <TableHead>Group</TableHead>
                <TableHead>Total Value</TableHead>
                <TableHead>Owner %</TableHead>
                <TableHead>Splits</TableHead>
                <TableHead className="text-right">Category</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {costs.length > 0 ? (
                costs.map((cost) => (
                  <TableRow key={cost.id}>
                    <TableCell className="font-medium text-zinc-800">
                      {cost.costName}
                    </TableCell>
                    <TableCell>{cost.groupName}</TableCell>
                    <TableCell>
                      {cost.totalValue.toLocaleString("en-US", {
                        style: "currency",
                        currency: "USD",
                      })}
                    </TableCell>
                    <TableCell>{cost.ownerPercentage}%</TableCell>
                    <TableCell>{cost.splitCount}</TableCell>
                    <TableCell className="text-right">
                      <BadgeType type={cost.category} />
                    </TableCell>
                  </TableRow>
                ))
              ) : (
                <TableRow>
                  <TableCell
                    colSpan={6}
                    className="text-center text-gray-500 py-10"
                  >
                    No costs found
                  </TableCell>
                </TableRow>
              )}
            </TableBody>
          </Table>
        </div>
      </div>
    </div>
  );
};
