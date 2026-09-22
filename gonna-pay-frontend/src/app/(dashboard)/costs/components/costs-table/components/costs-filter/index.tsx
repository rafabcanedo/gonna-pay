import { SlidersHorizontal } from 'lucide-react'
import { Popover, PopoverContent, PopoverTrigger } from '@/components/ui/popover'
import { Slider } from '@/components/ui/slider'
import { FilterSelect } from '@/components/filter-select'
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
        <FilterSelect value={category} onChange={onCategoryChange} options={COST_CATEGORY_OPTIONS} placeholder="Category" />
        <FilterSelect value={period} onChange={onPeriodChange} options={PERIOD_OPTIONS} placeholder="Period" />
        <FilterSelect value={type} onChange={onTypeChange} options={TYPE_OPTIONS} placeholder="Type" />

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
