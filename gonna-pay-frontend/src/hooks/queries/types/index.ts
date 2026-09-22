export type ContactQueryFilters = {
  category?: string
  search?: string
}

export type GroupQueryFilters = {
  category?: string
  search?: string
}

export type CostQueryFilters = {
  category?: string
  period?: string
  type?: string
  minValue?: number
  maxValue?: number
}
