import { useState } from "react"
import { Link } from "react-router"
import { TextAlignStart, X } from "lucide-react"
import { Logo } from "@/components/logo"
import { Button } from "@/components/ui/button"
import { navLinks } from "./constants"

export function Navbar() {
  const [isOpen, setIsOpen] = useState(false)

  return (
    <nav className="sticky top-0 z-50 bg-white/80 backdrop-blur-sm">
      <div className="flex items-center justify-between h-16 px-8">
        <Link to="/">
          <Logo size="sm" />
        </Link>

        <div className="hidden md:flex items-center gap-6">
          {navLinks.map(link =>
            link.href.startsWith('/') && !link.href.startsWith('/#')
              ? <Link key={link.label} to={link.href} className="text-sm text-muted-foreground hover:text-foreground transition-colors">{link.label}</Link>
              : <a key={link.label} href={link.href} className="text-sm text-muted-foreground hover:text-foreground transition-colors">{link.label}</a>
          )}
        </div>

        <Button variant="default" size="default" className="hidden md:flex w-40 h-12 border-2 border-primary bg-transparent hover:bg-primary hover:border-slate-400">
          Get Started
        </Button>

        <button className="md:hidden" onClick={() => setIsOpen(!isOpen)}>
          {isOpen ? <X size={20} /> : <TextAlignStart size={20} />}
        </button>
      </div>

      {isOpen && (
        <div className="md:hidden absolute top-full left-0 right-0 bg-white/90 backdrop-blur-md border-t border-border px-8 py-4 flex flex-col gap-4">
          {navLinks.map(link =>
            link.href.startsWith('/') && !link.href.startsWith('/#')
              ? <Link key={link.label} to={link.href} onClick={() => setIsOpen(false)} className="text-sm text-muted-foreground hover:text-foreground transition-colors py-2">{link.label}</Link>
              : <a key={link.label} href={link.href} onClick={() => setIsOpen(false)} className="text-sm text-muted-foreground hover:text-foreground transition-colors py-2">{link.label}</a>
          )}
          <Button variant="default" size="default" className="h-10" onClick={() => setIsOpen(false)}>
            Get Started
          </Button>
        </div>
      )}
    </nav>
  )
}
