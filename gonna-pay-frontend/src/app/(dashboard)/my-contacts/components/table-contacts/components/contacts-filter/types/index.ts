export type IPropsContactsFilter = {
  category: string
  search: string
  onCategoryChange: (value: string) => void
  onSearchChange: (value: string) => void
}
