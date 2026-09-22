import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import type { IPropsCategoryCards } from './interfaces'
import { CONTACT_CATEGORIES } from '@/app/(dashboard)/my-contacts/constants'

export function CategoryCards({ stats }: IPropsCategoryCards) {
  return (
    <>
      {CONTACT_CATEGORIES.map(({ label, value }) => {
        const count = stats.byCategory[value] ?? 0
        return (
          <Card key={value} className="w-40">
            <CardHeader>
              <CardTitle className="text-sm text-zinc-500 font-light">{label}</CardTitle>
            </CardHeader>
            <CardContent>
              <span className="text-2xl font-semibold text-zinc-700">{count}</span>
            </CardContent>
          </Card>
        )
      })}
    </>
  )
}
