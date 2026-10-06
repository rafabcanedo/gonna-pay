import { CATEGORY_ICONS } from '@/utils/category-icons'
import { CATEGORY_STYLES } from './constants'
import type { ITypeBadge } from './interfaces'

export function BadgeType({ type }: ITypeBadge) {
  const style = CATEGORY_STYLES[type] ?? "bg-blue-100 text-blue-800";
  const Icon = CATEGORY_ICONS[type];

  return (
    <span
      className={`inline-flex items-center gap-1 px-2.5 py-0.5 rounded-full text-xs font-medium ${style}`}
    >
      {Icon != null && <Icon className="w-3 h-3" />}
      {type}
    </span>
  );
}
