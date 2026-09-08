'use client'

import { useEffect, useState } from 'react'
import { CalendarClock } from 'lucide-react'
import { CreateReminders } from '../create-reminders'
import type { IReminder, IReminderSerialized, ICreateReminder } from '../interfaces'

export function RemindersWidget() {
  const [reminders, setReminders] = useState<IReminder[]>([])

  useEffect(() => {
    const saved = sessionStorage.getItem('@app:reminders')
    if (!saved) return
    try {
      const parsed: IReminderSerialized[] = JSON.parse(saved)
      setReminders(parsed.map((r) => ({ ...r, date: new Date(r.date) })))
    } catch {}
  }, [])

  useEffect(() => {
    sessionStorage.setItem('@app:reminders', JSON.stringify(reminders))
  }, [reminders])

  const handleAddReminder = (reminder: ICreateReminder) => {
    setReminders((prev) => [...prev, { ...reminder, id: crypto.randomUUID() }])
  }

  return (
    <div className="rounded-xl border border-zinc-200 bg-white shadow-sm">
      <header className="flex items-center justify-between px-6 pt-4 pb-3 border-b border-zinc-100">
        <h2 className="font-mono text-lg font-semibold text-zinc-600">Reminders</h2>
        <CreateReminders onAddReminder={handleAddReminder} />
      </header>

      <div className="flex flex-col divide-y divide-zinc-100">
        {reminders.length === 0 ? (
          <p className="px-6 py-8 text-center text-sm text-zinc-400">No reminders yet</p>
        ) : (
          reminders.map((r) => (
            <div key={r.id} className="flex items-center justify-between px-6 py-3">
              <div className="flex items-center gap-3">
                <CalendarClock className="h-4 w-4 text-zinc-400 shrink-0" />
                <span className="text-sm font-medium text-zinc-700">{r.name}</span>
              </div>
              <div className="flex items-center gap-4 text-sm text-zinc-500">
                <span className="font-mono">${r.value}</span>
                <span>
                  {r.date.toLocaleDateString('en-US', { day: '2-digit', month: 'short' })}
                </span>
              </div>
            </div>
          ))
        )}
      </div>
    </div>
  )
}
