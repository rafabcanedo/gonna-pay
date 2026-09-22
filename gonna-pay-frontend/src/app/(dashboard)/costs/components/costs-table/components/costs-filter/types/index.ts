export type IPropsCostsFilter = {
  category: string
  period: string
  type: string
  valueRange: [number, number]
  onCategoryChange: (value: string) => void
  onPeriodChange: (value: string) => void
  onTypeChange: (value: string) => void
  onValueRangeChange: (value: [number, number]) => void
}
