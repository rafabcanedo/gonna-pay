import { TableCell, TableRow } from '@/components/ui/table'
import type { ITableEmptyState } from './interfaces'

export const TableEmptyState = ({ colSpan, message }: ITableEmptyState) => (
  <TableRow>
    <TableCell colSpan={colSpan} className="text-center text-gray-500 py-10">
      {message}
    </TableCell>
  </TableRow>
)
