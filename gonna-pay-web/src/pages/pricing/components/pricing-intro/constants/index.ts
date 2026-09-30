import type { MockCardData } from "../types"

export const MOCK_CARDS: MockCardData[] = [
  {
    emoji: "🧳",
    title: "Weekend Trip",
    total: "$280.00",
    splits: [
      { name: "You", amount: "$93.00", paid: true },
      { name: "Carlos", amount: "$93.00", paid: true },
      { name: "Maria", amount: "$94.00", paid: false },
    ],
  },
  {
    emoji: "🍕",
    title: "Dinner with Friends",
    total: "$45.00",
    splits: [
      { name: "You", amount: "$20.00", paid: true },
      { name: "Ana", amount: "$10.00", paid: true },
      { name: "João", amount: "$15.00", paid: false },
    ],
  },
  {
    emoji: "🎬",
    title: "Movie Night",
    total: "$32.00",
    splits: [
      { name: "You", amount: "$16.00", paid: true },
      { name: "Isa", amount: "$16.00", paid: false },
    ],
  },
]
