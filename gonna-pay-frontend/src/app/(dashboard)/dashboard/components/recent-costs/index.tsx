import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table"
import Link from "next/link"
import { BadgeType } from "@/utils/badge-types"
import { CostsTableProps } from '../interfaces'
import { RECENT_COSTS_LIMIT } from './constants'

export const RecentCosts = ({ costs }: CostsTableProps) => {
  const limitedCosts = costs.slice(0, RECENT_COSTS_LIMIT)

  return (
    <div className="mt-6 mb-6">
      <div className="w-full max-w-5xl rounded-xl border border-zinc-200 bg-white shadow-sm">
        <div className="flex items-center justify-between px-6 pt-4 pb-2">
          <h2 className="font-mono text-lg font-semibold text-zinc-600">
            Recent Costs
          </h2>
        </div>

        <div className="px-2 pb-2">
          <Table className="w-full">
            <TableHeader>
              <TableRow>
                <TableHead className="w-[140px]">Name</TableHead>
                <TableHead className="text-right">Value</TableHead>
                <TableHead className="text-right">Percent</TableHead>
                <TableHead className="text-right">Category</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {limitedCosts.map((cost) => (
                <TableRow key={cost.id}>
                  <TableCell className="font-medium text-zinc-800">
                    {cost.costName}
                  </TableCell>

                  <TableCell className="text-right font-mono text-sm text-zinc-700">
                    {Number(cost.totalValue).toLocaleString("en-US", {
                      style: "currency",
                      currency: "USD",
                    })}
                  </TableCell>

                  <TableCell className="text-right font-mono text-sm text-zinc-700">
                    {cost.ownerPercentage}%
                  </TableCell>

                  <TableCell className="text-right text-sm text-zinc-500">
                    <BadgeType type={cost.category} />
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </div>

        <div className="flex justify-end px-6 py-3 border-t border-zinc-100">
          <Link href="/costs">
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