import { Contact, ContactCategory } from "@/types";

const contactsMock: Contact[] = [
  {
    id: "1",
    name: "Wania",
    email: "wania@gmail.com",
    phone: "99999987",
    category: ContactCategory.FAMILY,
  },
  {
    id: "2",
    name: "John Doe",
    email: "john.doe@example.com",
    phone: "555-0101",
    category: ContactCategory.FRIEND,
  },
  {
    id: "3",
    name: "Jane Smith",
    email: "jane.smith@example.com",
    phone: "555-0102",
    category: ContactCategory.WORK,
  },
  {
    id: "4",
    name: "Carlos Lima",
    email: "carlos.lima@example.com",
    phone: "555-0103",
    category: ContactCategory.FAMILY,
  },
  {
    id: "5",
    name: "Maria Garcia",
    email: "maria.garcia@example.com",
    phone: "555-0104",
    category: ContactCategory.FRIEND,
  },
  {
    id: "6",
    name: "Alex Johnson",
    email: "alex.johnson@example.com",
    phone: "555-0105",
    category: ContactCategory.WORK,
  },
  {
    id: "7",
    name: "Sofia Rossi",
    email: "sofia.rossi@example.com",
    phone: "555-0106",
    category: ContactCategory.FAMILY,
  },
  {
    id: "8",
    name: "Liam Brown",
    email: "liam.brown@example.com",
    phone: "555-0107",
    category: ContactCategory.FRIEND,
  },
]

export { contactsMock }
