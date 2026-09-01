import { Contact, Cost } from "@/types";

export interface IPropsDashboardCards {
  title: string;
  value: number;
}

export interface RecentContactsProps {
  contacts: Contact[];
  total: number;
}

export interface CostsTableProps {
  costs: Cost[];
  total: number;
}

export interface IReminder {
  id: string;
  name: string;
  value: string;
  date: Date;
}

export interface IReminderSerialized {
  id: string;
  name: string;
  value: string;
  date: string;
}

export interface ICreateReminder {
  name: string;
  value: string;
  date: Date;
}

export interface CreateRemindersProps {
  onAddReminder: (reminder: ICreateReminder) => void;
}
