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
          <SelectItem value={ContactCategory.WORK}>Work</SelectItem>
          <SelectItem value={ContactCategory.FAMILY}>Family</SelectItem>
          <SelectItem value={ContactCategory.FRIEND}>Friend</SelectItem>
        </SelectGroup>
      </SelectContent>
    </Select>
  )
}