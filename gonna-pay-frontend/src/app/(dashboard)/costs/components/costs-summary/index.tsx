import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import type { IPropsCostsSummary } from './interfaces'
import { SUMMARY_CARDS } from './constants'

export function CostsSummary({ thisMonth, inSplits, solo }: IPropsCostsSummary) {
  const values = { thisMonth, inSplits, solo }

  return (
    <>
      {SUMMARY_CARDS.map(({ label, key }) => (
        <Card key={label} className="flex-1">
          <CardHeader>
            <CardTitle className="text-sm text-zinc-500 font-light">{label}</CardTitle>
          </CardHeader>
          <CardContent>
            <span className="text-2xl font-semibold text-zinc-700">
              $ {values[key].toFixed(2)}
            </span>
          </CardContent>
        </Card>
      ))}
    </>
  )
}
