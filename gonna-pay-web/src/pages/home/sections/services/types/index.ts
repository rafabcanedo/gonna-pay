export type ServiceVariant = "default" | "primary" | "dark"
export type ServiceArtKey = "tracking" | "group" | "math" | "contacts"

export interface ServiceItem {
  artKey: ServiceArtKey
  title: string
  description: string
  variant: ServiceVariant
}
