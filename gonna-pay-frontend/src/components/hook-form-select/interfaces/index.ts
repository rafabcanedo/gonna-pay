import type { LucideIcon } from 'lucide-react'

export interface HookFormSelectProps {
  name: string;
  label: string;
  placeholder?: string;
  options: { label: string; value: string; icon?: LucideIcon }[];
  groupLabel?: string;
}
