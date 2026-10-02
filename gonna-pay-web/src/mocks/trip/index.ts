import { Hotel, UtensilsCrossed, Car } from "lucide-react"

import type { TripCost, TripSplit } from "./types"

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
