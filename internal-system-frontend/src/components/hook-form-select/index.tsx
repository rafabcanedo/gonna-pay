"use client"

import { useFormContext } from "react-hook-form"
import {
    Select,
    SelectContent,
    SelectGroup,
    SelectItem,
    SelectLabel,
    SelectTrigger,
    SelectValue,
} from "@/components/ui/select"
import { FormField, FormItem, FormLabel, FormMessage } from "@/components/ui/form"
import { HookFormSelectProps } from './interfaces'

export function HookFormSelect({
    name,
    label,
    placeholder,
    options,
    groupLabel,
}: HookFormSelectProps) {
    const { control } = useFormContext()

    return (
        <FormField
            control={control}
            name={name}
            render={({ field }) => (
                <FormItem>
                    <FormLabel className="font-normal">{label}</FormLabel>
                    <Select
                        onValueChange={field.onChange}
                        value={field.value}
                    >
                        <SelectTrigger>
                            <SelectValue placeholder={placeholder} />
                        </SelectTrigger>
                        <SelectContent>
                            {groupLabel && (
                                <SelectGroup>
                                    <SelectLabel>{groupLabel}</SelectLabel>
                                    {options.map((opt) => (
                                        <SelectItem key={opt.value} value={opt.value}>
                                            {opt.label}
                                        </SelectItem>
                                    ))}
                                </SelectGroup>
                            )}
                            {!groupLabel &&
                                options.map((opt) => (
                                    <SelectItem key={opt.value} value={opt.value}>
                                        {opt.label}
                                    </SelectItem>
                                ))
                            }
                        </SelectContent>
                    </Select>
                    <FormMessage />
                </FormItem>
            )}
        />
    )
}