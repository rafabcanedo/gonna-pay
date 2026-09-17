export interface ITablePagination {
  page: number
  totalPages: number
  onPageChange: (page: number) => void
}
