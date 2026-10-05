import * as React from "react"
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectLabel,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select"
import { SelectCategoryProps } from './types'
import { ContactCategory } from "@/types"
import { CONTACT_CATEGORIES } from "@/app/(dashboard)/my-contacts/constants"

export const SelectCategory: React.FC<SelectCategoryProps> = ({ value, onValueChange }) => {
  return (
    <Select
      value={value}
      onValueChange={(val) => onValueChange(val as ContactCategory)}
    >
      <SelectTrigger className="w-[180px]">
        <SelectValue placeholder="Select a category" />
      </SelectTrigger>
      <SelectContent>
        <SelectGroup>
          <SelectLabel>Category</SelectLabel>
          {CONTACT_CATEGORIES.map((opt) => (
            <SelectItem key={opt.value} value={opt.value}>
              <span className="flex items-center gap-2">
                {opt.icon && <opt.icon className="w-4 h-4" />}
                {opt.label}
              </span>
            </SelectItem>
          ))}
        </SelectGroup>
      </SelectContent>
    </Select>
  )
}