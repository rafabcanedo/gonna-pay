import { ContactCategory, TransactionCategory } from "@/types";

interface ITypeBadge {
  type: ContactCategory | TransactionCategory
}

const categoryStyles: Record<string, string> = {
  [ContactCategory.FAMILY]: "bg-blue-500/60 text-blue-800",
  [ContactCategory.FRIEND]: "bg-yellow-300/60 text-yellow-800",
  [ContactCategory.WORK]: "bg-green-500/60 text-green-800",

  [TransactionCategory.DINNER]: "bg-blue-500/60 text-blue-800",
  [TransactionCategory.LUNCH]: "bg-green-500/60 text-green-800",
  [TransactionCategory.ENTERTAINMENT]: "bg-emerald-500/60 text-emerald-800",
  [TransactionCategory.TRAVEL]: "bg-orange-500/60 text-orange-800",
  [TransactionCategory.OTHERS]: "bg-sky-500/60 text-sky-800",
};

export function BadgeType({ type }: ITypeBadge) {
  const style = categoryStyles[type] ?? "bg-blue-100 text-blue-800";

  return (
    <span
      className={`inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium ${style}`}
    >
      {type}
    </span>
  );
}
