import { Hotel, UtensilsCrossed, Car } from "lucide-react"

import type { TripCost, TripSplit, Step } from "../types"

export const TRIP_NAME = "Madrid Trip"
export const TRIP_OWNER_PERCENTAGE = 50

export const TRIP_COSTS: TripCost[] = [
  { icon: Hotel,           name: "Hotel Santo Mauro",  total: 450 },
  { icon: UtensilsCrossed, name: "Casa Carmen dinner", total: 90  },
  { icon: Car,             name: "Uber to airport",    total: 36  },
]

export const TRIP_SPLITS: TripSplit[] = [
  { name: "You",          owes: 288, isOwner: true  },
  { name: "Ana Lima",     owes: 144, isOwner: false },
  { name: "Marcos Souza", owes: 144, isOwner: false },
]

export const STEPS: Step[] = [
  {
    number: "01",
    title: "Add your contacts",
    description: "Register who you share expenses with — friends, family, or coworkers.",
  },
  {
    number: "02",
    title: "Create a group",
    description: "Group your contacts by context: a trip, a dinner, a work project.",
  },
  {
    number: "03",
    title: "Log a cost",
    description: "Register the expense, link it to a group and set your percentage to cover.",
  },
  {
    number: "04",
    title: "Splits appear automatically",
    description: "Gonna Pay calculates and assigns each person's share instantly.",
  },
]
