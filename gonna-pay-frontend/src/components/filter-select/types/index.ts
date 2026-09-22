export type IPropsFilterSelect = {
  value: string
  onChange: (value: string) => void
  options: { label: string; value: string }[]
  placeholder: string
  className?: string
}
