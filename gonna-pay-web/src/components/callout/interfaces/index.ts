import type { ReactNode } from "react"

export interface IPropsCallout {
  type: "tip" | "important"
  children: ReactNode
}
