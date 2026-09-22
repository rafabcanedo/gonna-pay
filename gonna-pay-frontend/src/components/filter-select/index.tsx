import { cn } from '@/lib/utils'
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select'
import type { IPropsFilterSelect } from './types'

export function FilterSelect({ value, onChange, options, placeholder, className }: IPropsFilterSelect) {
  return (
    <Select value={value === '' ? 'all' : value} onValueChange={(val) => onChange(val === 'all' ? '' : val)}>
      <SelectTrigger className={cn('h-8 text-xs', className)}>
        <SelectValue placeholder={placeholder} />
      </SelectTrigger>
      <SelectContent>
        {options.map((opt) => (
          <SelectItem key={opt.value} value={opt.value}>
            {opt.label}
          </SelectItem>
        ))}
      </SelectContent>
    </Select>
  )
}
