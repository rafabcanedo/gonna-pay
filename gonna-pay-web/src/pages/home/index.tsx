import { Intro } from "./sections/intro"
import { Services } from "./sections/services"
import { UseCases } from "./sections/use-cases"
import { Feedbacks } from "./sections/feedbacks"
import { HomePricing } from "./sections/pricing"
import { ContactUs } from "./sections/contact-us"

export function Home() {
  return (
    <div className="flex flex-col">
      <Intro />
      <Services />
      <UseCases />
      <Feedbacks />
      <HomePricing />
      <ContactUs />
    </div>
  )
}
