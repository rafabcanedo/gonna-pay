'use client'

import { useState } from 'react'
import { useRouter } from 'next/navigation'
import { MoreHorizontal } from 'lucide-react'
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@/components/ui/table'
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu'
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
} from '@/components/ui/alert-dialog'
import { Button } from '@/components/ui/button'
import { BadgeType } from '@/utils/badge-types'
import { TableEmptyState } from '@/components/table-empty-state'
import { useCostsQuery } from '@/hooks/queries/use-cost-query'
import { useDeleteCost } from '@/hooks/mutations/use-cost-mutations'
import type { Cost } from '@/types'

export const CostsTable = () => {
  const router = useRouter()
  const { data } = useCostsQuery()
  const { mutate: deleteCost, isPending } = useDeleteCost()

  const [costToDelete, setCostToDelete] = useState<Cost | null>(null)

  const costs = data ?? []
  const total = data?.length ?? 0

  return (
    <>
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
                  <TableHead className="w-[50px]" />
                </TableRow>
              </TableHeader>
              <TableBody>
                {costs.length > 0 ? (
                  costs.map((cost) => (
                    <TableRow key={cost.id}>
                      <TableCell className="font-medium text-zinc-800">{cost.costName}</TableCell>
                      <TableCell>{cost.groupName || '-'}</TableCell>
                      <TableCell>
                        {cost.totalValue.toLocaleString('en-US', { style: 'currency', currency: 'USD' })}
                      </TableCell>
                      <TableCell>{cost.ownerPercentage}%</TableCell>
                      <TableCell>{cost.splitCount}</TableCell>
                      <TableCell className="text-right">
                        <BadgeType type={cost.category} />
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
                              onClick={() => router.push(`/costs/details/${cost.id}?name=${encodeURIComponent(cost.costName)}`)}
                            >
                              Edit
                            </DropdownMenuItem>
                            <DropdownMenuItem
                              className="text-red-600 focus:text-red-600"
                              onClick={() => setCostToDelete(cost)}
                            >
                              Delete
                            </DropdownMenuItem>
                          </DropdownMenuContent>
                        </DropdownMenu>
                      </TableCell>
                    </TableRow>
                  ))
                ) : (
                  <TableEmptyState colSpan={7} message="No costs found" />
                )}
              </TableBody>
            </Table>
          </div>
        </div>
      </div>

      <AlertDialog open={!!costToDelete} onOpenChange={(open) => { if (!open) setCostToDelete(null) }}>
        <AlertDialogContent>
          <AlertDialogHeader>
            <AlertDialogTitle>Delete cost</AlertDialogTitle>
            <AlertDialogDescription>
              Are you sure you want to delete <strong>{costToDelete?.costName}</strong>? This action cannot be undone.
            </AlertDialogDescription>
          </AlertDialogHeader>
          <AlertDialogFooter>
            <AlertDialogCancel disabled={isPending}>Cancel</AlertDialogCancel>
            <AlertDialogAction
              className="bg-red-600 hover:bg-red-700"
              disabled={isPending}
              onClick={() => {
                if (costToDelete) {
                  deleteCost(costToDelete.id, { onSuccess: () => setCostToDelete(null) })
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
