"use client";

import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";

import { TablePaymentsProps } from './interfaces'

export const TablePayments = ({ total }: TablePaymentsProps) => {
  return (
    <div className="w-full flex justify-center mt-6 mb-6">
      <div className="w-full max-w-7xl rounded-xl border bg-white shadow-sm">
        <div className="flex items-center justify-between px-6 pt-4 pb-2">
          <h2 className="text-lg">Contacts</h2>
          <span className="text-xs">{total} contacts</span>
        </div>

        <div className="px-2 pb-2">
          <Table className="w-full">
            <TableHeader>
              <TableRow>
                <TableHead className="w-[100px]">Name</TableHead>
                <TableHead>Who</TableHead>
                <TableHead>Value</TableHead>
                <TableHead className="text-right">Category</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              <TableRow>
                <TableCell className="font-medium">Rafael</TableCell>
                <TableCell>Canedo</TableCell>
                <TableCell>$ 200</TableCell>
              </TableRow>
            </TableBody>
          </Table>
        </div>
      </div>
    </div>
  );
};
