'use client'

import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { useContactFrequencyQuery } from '@/hooks/queries/use-contact-query'

export function TopContacts() {
  const { data } = useContactFrequencyQuery(5)

  return (
    <Card className="flex-1">
      <CardHeader>
        <CardTitle className="text-sm text-zinc-500 font-light">Top Contacts</CardTitle>
      </CardHeader>
      <CardContent className="flex flex-col gap-2">
        {data?.map((item, index) => (
          <div key={item.contactId} className="flex items-center justify-between">
            <div className="flex items-center gap-2">
              <span className="text-xs text-zinc-400">{index + 1}.</span>
              <span className="text-sm text-zinc-700">{item.contactName}</span>
            </div>
            <span className="text-xs text-zinc-400">{item.sharedCosts}×</span>
          </div>
        ))}
      </CardContent>
    </Card>
  )
}
