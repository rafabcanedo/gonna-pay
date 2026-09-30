import { useState } from "react"
import { Button } from "@/components/ui/button"

export const SubscribeForm = () => {
    const [email, setEmail] = useState("")

    return (
        <div className="flex items-start gap-2 ml-100">
            <input
                type="email"
                placeholder="Email"
                value={email}
                onChange={e => setEmail(e.target.value)}
                className="h-10 rounded-lg border border-zinc-600 bg-zinc-700 px-3 text-sm text-white placeholder:text-zinc-400 outline-none focus-visible:border-zinc-400 transition-all"
            />
            <Button variant="default" size="default" className="h-10">Subscribe to news</Button>
        </div>
    )
}
