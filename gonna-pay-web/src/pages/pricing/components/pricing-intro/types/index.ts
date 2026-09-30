export interface MockSplit {
  name: string
  amount: string
  paid: boolean
}

export interface MockCardData {
  emoji: string
  title: string
  total: string
  splits: MockSplit[]
}
