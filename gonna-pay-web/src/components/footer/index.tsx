import { Link } from "react-router"
import { GithubIcon } from "@/assets/github-icon"
import { Logo } from "@/components/logo"
import { Title } from "@/components/title"
import { footerNavLinks, socialLinks } from "./constants"
import { SubscribeForm } from "./components/subscribe-form"

export function Footer() {
  return (
    <footer className="bg-zinc-800 mx-8 mb-8 rounded-2xl px-12 py-10 flex flex-col gap-8">

      <div className="flex items-center justify-between">
        <Logo size="xs" variant="dark" />

        <div className="flex items-center gap-6">
          {footerNavLinks.map(link =>
            link.href.startsWith('/') && !link.href.startsWith('/#')
              ? <Link key={link.label} to={link.href} className="text-sm text-zinc-400 hover:text-white transition-colors underline">{link.label}</Link>
              : <a key={link.label} href={link.href} className="text-sm text-zinc-400 hover:text-white transition-colors underline">{link.label}</a>
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
              className="text-zinc-400 hover:text-white transition-colors"
            >
              <GithubIcon />
            </a>
          ))}
        </div>
      </div>

      <div className="flex gap-12">
        <div className="flex flex-col gap-2">
          <Title variant="plain" size="sm" className="text-zinc-400">Contact us:</Title>
          <p className="text-sm text-zinc-400">contact@gonnapay.app</p>
          <p className="text-xs text-zinc-500 max-w-xs">
            This is a study project. The email above is fictional and not monitored.
          </p>
        </div>

        <SubscribeForm />
      </div>

      <div className="border-t border-zinc-700 pt-6 flex items-center gap-8">
        <p className="text-xs text-zinc-500">© 2026 Gonna Pay. All rights reserved.</p>
        <a href="#about" className="text-xs text-zinc-500 hover:text-white transition-colors underline">
          About this project
        </a>
      </div>

    </footer>
  )
}
