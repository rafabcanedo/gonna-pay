import { SlidersHorizontal } from 'lucide-react'
import { Popover, PopoverContent, PopoverTrigger } from '@/components/ui/popover'
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from '@/components/ui/select'
import { Slider } from '@/components/ui/slider'
import { COST_CATEGORY_OPTIONS, PERIOD_OPTIONS, TYPE_OPTIONS, MAX_VALUE } from '@/app/(dashboard)/costs/constants'
import type { IPropsCostsFilter } from './types'

export function CostsFilter({ category, period, type, valueRange, onCategoryChange, onPeriodChange, onTypeChange, onValueRangeChange }: IPropsCostsFilter) {
  const hasActiveFilters = category !== '' || period !== '' || type !== '' || valueRange[0] !== 0 || valueRange[1] !== MAX_VALUE

  return (
    <Popover>
      <PopoverTrigger asChild>
        <button
          type="button"
          className={`p-1.5 rounded-md hover:bg-zinc-100 transition-colors relative ${hasActiveFilters ? 'text-zinc-800' : 'text-zinc-400 hover:text-zinc-600'}`}
        >
          <SlidersHorizontal className="w-4 h-4" />
          {hasActiveFilters && (
            <span className="absolute -top-0.5 -right-0.5 w-2 h-2 bg-zinc-800 rounded-full" />
          )}
        </button>
      </PopoverTrigger>

      <PopoverContent align="end" className="w-72 flex flex-col gap-4 p-4">
        <Select value={category} onValueChange={onCategoryChange}>
          <SelectTrigger className="h-8 text-xs">
            <SelectValue placeholder="Category" />
          </SelectTrigger>
          <SelectContent>
            {COST_CATEGORY_OPTIONS.map((opt) => (
              <SelectItem key={opt.value} value={opt.value}>{opt.label}</SelectItem>
            ))}
          </SelectContent>
        </Select>

        <Select value={period} onValueChange={onPeriodChange}>
          <SelectTrigger className="h-8 text-xs">
            <SelectValue placeholder="Period" />
          </SelectTrigger>
          <SelectContent>
            {PERIOD_OPTIONS.map((opt) => (
              <SelectItem key={opt.value} value={opt.value}>{opt.label}</SelectItem>
            ))}
          </SelectContent>
        </Select>

        <Select value={type} onValueChange={onTypeChange}>
          <SelectTrigger className="h-8 text-xs">
            <SelectValue placeholder="Type" />
          </SelectTrigger>
          <SelectContent>
            {TYPE_OPTIONS.map((opt) => (
              <SelectItem key={opt.value} value={opt.value}>{opt.label}</SelectItem>
            ))}
          </SelectContent>
        </Select>

        <div className="flex flex-col gap-2">
          <span className="text-xs text-zinc-500">
            Value: ${valueRange[0].toLocaleString()} — ${valueRange[1].toLocaleString()}
          </span>
          <Slider
            min={0}
            max={MAX_VALUE}
            step={50}
            value={valueRange}
            onValueChange={(v) => onValueRangeChange(v as [number, number])}
          />
        </div>
      </PopoverContent>
    </Popover>
  )
}
