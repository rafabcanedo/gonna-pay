"use client";

import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { CreateReminders } from "../create-reminders";
import { useEffect, useState } from "react";
import { IReminder, IReminderSerialized } from '../interfaces'

export const Reminders = () => {
  const [reminders, setReminders] = useState<IReminder[]>([]);

  useEffect(() => {
    const savedData = sessionStorage.getItem("@app:reminders");
    if (savedData) {
      try {
        const parsed: IReminderSerialized[] = JSON.parse(savedData);

        const formatted: IReminder[] = parsed.map((item) => ({
          ...item,
          date: new Date(item.date),
        }));

        setReminders(formatted);
      } catch (error) {
        console.error("Failed to load session reminders:", error);
      }
    }
  }, []);

  useEffect(() => {
    sessionStorage.setItem("@app:reminders", JSON.stringify(reminders));
  }, [reminders]);

  const handleAddReminder = (reminder: Omit<IReminder, "id">) => {
    const newReminder: IReminder = {
      ...reminder,
      id: crypto.randomUUID(),
    };
    setReminders((prev) => [...prev, newReminder]);
  };

  return (
    <div className="mt-6 mb-6">
      <div className="w-full max-w-5xl rounded-xl border border-zinc-200 bg-white shadow-sm">
        <header className="flex flex-row items-center justify-between px-6 pt-4 pb-2">
          <h2 className="font-mono text-lg font-semibold text-zinc-600">
            Reminders
          </h2>
          <CreateReminders onAddReminder={handleAddReminder} />
        </header>

        <div className="px-2 pb-2">
          <Table className="w-full">
            <TableHeader>
              <TableRow>
                <TableHead className="w-[140px]">Reminders Name</TableHead>
                <TableHead className="text-right">Value</TableHead>
                <TableHead className="text-right">Date</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {reminders.length === 0 ? (
                <TableRow>
                  <TableCell colSpan={3} className="text-center text-zinc-400">
                    No reminders yet
                  </TableCell>
                </TableRow>
              ) : (
                reminders.map((reminder) => (
                  <TableRow key={reminder.id}>
                    <TableCell className="font-medium">
                      {reminder.name}
                    </TableCell>
                    <TableCell className="text-right">
                      {reminder.value}
                    </TableCell>
                    <TableCell className="text-right">
                      {reminder.date.toLocaleDateString()}
                    </TableCell>
                  </TableRow>
                ))
              )}
            </TableBody>
          </Table>
        </div>
      </div>
    </div>
  );
};
