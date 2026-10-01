import { Button } from "@/components/ui/button"

import { HeroArt } from "./components/hero-art"

export function Intro() {
  return (
    <section id="intro" className="px-8 py-24 flex flex-col items-center text-center">
      <div className="flex flex-col">
        <h1>You enjoy the moment</h1>
        <h1>We track the bill</h1>
      </div>

      <p className="text-muted-foreground text-base mt-4 max-w-md">
        Create a group, add your expenses and set each person's share.
        Gonna Pay figures out the rest.
      </p>

      <Button className="mt-8">Get Started</Button>

      <div className="mt-16 w-full max-w-2xl mx-auto">
        <HeroArt />
      </div>
    </section>
  )
}
