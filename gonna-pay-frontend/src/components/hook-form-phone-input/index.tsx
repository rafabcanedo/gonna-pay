'use client'

import { FC } from 'react'
import { useFormContext, useController } from 'react-hook-form'
import { Input } from '@/components/ui/input'
import type { IPhoneInput } from './types'

export const formatPhone = (value: string): string => {
  const digits = (value ?? '').replace(/\D/g, '').slice(0, 11)
  const len = digits.length
  if (len <= 2) return digits
  if (len <= 6) return `(${digits.slice(0, 2)}) ${digits.slice(2)}`
  if (len <= 10) return `(${digits.slice(0, 2)}) ${digits.slice(2, 6)}-${digits.slice(6)}`
  return `(${digits.slice(0, 2)}) ${digits.slice(2, 7)}-${digits.slice(7)}`
}

export const HookFormPhoneInput: FC<IPhoneInput> = ({ name, title, label, type = 'tel', disabled }) => {
  const { control, formState: { errors } } = useFormContext()
  const { field } = useController({ control, name })

  return (
    <div>
      <label className='text-xs text-zinc-600'>{title}</label>
      <Input
        {...field}
        type={type}
        placeholder={label}
        disabled={disabled}
        onChange={(e) => field.onChange(formatPhone(e.target.value))}
        className='disabled:border-zinc-400 disabled:cursor-not-allowed'
      />
      <p className="text-xs text-red-400">
        {errors[name]?.message as string}
      </p>
    </div>
  )
}
