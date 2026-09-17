'use client'

import {
  Pagination,
  PaginationContent,
  PaginationItem,
  PaginationNext,
  PaginationPrevious,
} from '@/components/ui/pagination'
import type { ITablePagination } from './interfaces'

export const TablePagination = ({ page, totalPages, onPageChange }: ITablePagination) => (
  <Pagination>
    <PaginationContent>
      <PaginationItem>
        <PaginationPrevious
          onClick={() => onPageChange(page - 1)}
          aria-disabled={page <= 1}
          className={page <= 1 ? 'pointer-events-none opacity-50' : ''}
        />
      </PaginationItem>
      <PaginationItem>
        <span className="text-sm text-muted-foreground px-4">{page} / {totalPages}</span>
      </PaginationItem>
      <PaginationItem>
        <PaginationNext
          onClick={() => onPageChange(page + 1)}
          aria-disabled={page >= totalPages}
          className={page >= totalPages ? 'pointer-events-none opacity-50' : ''}
        />
      </PaginationItem>
    </PaginationContent>
  </Pagination>
)
