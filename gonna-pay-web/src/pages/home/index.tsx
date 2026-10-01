import { Intro } from "./sections/intro"
import { Services } from "./sections/services"
import { UseCases } from "./sections/use-cases"
import { Other } from "./sections/other"
import { HomePricing } from "./sections/pricing"
import { ContactUs } from "./sections/contact-us"

export function Home() {
  return (
    <div className="flex flex-col">
      <Intro />
      <Services />
      <UseCases />
      <Other />
      <HomePricing />
      <ContactUs />
    </div>
  )
}
