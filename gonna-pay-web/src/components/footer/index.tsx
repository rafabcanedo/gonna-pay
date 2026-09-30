import { useState } from "react"
import { Link } from "react-router"
import { Github } from "lucide-react"
import { Logo } from "@/components/logo"
import { Title } from "@/components/title"
import { Button } from "@/components/ui/button"
import { footerNavLinks, socialLinks } from "./constants"

export function Footer() {
  const [email, setEmail] = useState("")

  return (
    <footer className="border-t border-border bg-background">
      <div className="px-8 py-10 flex flex-col gap-8">

        <div className="flex items-center justify-between">
          <Logo size="xs" />

          <div className="flex items-center gap-6">
            {footerNavLinks.map(link =>
              link.href.startsWith('/') && !link.href.startsWith('/#')
                ? <Link key={link.label} to={link.href} className="text-sm text-muted-foreground hover:text-foreground transition-colors">{link.label}</Link>
                : <a key={link.label} href={link.href} className="text-sm text-muted-foreground hover:text-foreground transition-colors">{link.label}</a>
            )}
          </div>

          <div className="flex items-center gap-3">
            {socialLinks.map(link => (
              <a
                key={link.label}
                href={link.href}
                target="_blank"
                rel="noopener noreferrer"
                aria-label={link.label}
                className="text-muted-foreground hover:text-foreground transition-colors"
              >
                <Github size={20} />
              </a>
            ))}
          </div>
        </div>

        <div className="flex gap-12">
          <div className="flex flex-col gap-2">
            <Title variant="plain" size="sm">Contact us:</Title>
            <p className="text-sm text-muted-foreground">contact@gonnapay.app</p>
            <p className="text-xs text-muted-foreground max-w-xs">
              This is a study project. The email above is fictional and not monitored.
            </p>
          </div>

          <div className="flex items-start gap-2 ml-auto">
            <input
              type="email"
              placeholder="Email"
              value={email}
              onChange={e => setEmail(e.target.value)}
              className="h-8 rounded-lg border border-border bg-background px-3 text-sm text-foreground placeholder:text-muted-foreground outline-none focus-visible:border-ring focus-visible:ring-3 focus-visible:ring-ring/50 transition-all"
            />
            <Button variant="default" size="default">Subscribe to news</Button>
          </div>
        </div>

        <div className="border-t border-border pt-6 flex items-center justify-between">
          <p className="text-xs text-muted-foreground">
            © 2026 Gonna Pay. All rights reserved.
          </p>
          <a href="#about" className="text-xs text-muted-foreground hover:text-foreground transition-colors">
            About this project
          </a>
        </div>

      </div>
    </footer>
  )
}
