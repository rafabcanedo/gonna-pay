<div align="center">
  <img src="src/assets/gonna-logo.svg" alt="Gonna Pay" width="200" />
</div>

<br />

<div align="center">
  <p>Dashboard App — Gonna Pay Frontend</p>
</div>

---

## About

The main Gonna Pay application. A protected dashboard where users manage contacts, groups, costs and track expense splits.

## Tech Stack

- **Next.js 15** — framework (App Router)
- **React 19** — UI
- **TypeScript** — language
- **Tailwind CSS v4** — styling
- **Shadcn/ui** — component library
- **TanStack React Query v5** — server state
- **React Hook Form + Yup** — forms and validation
- **Sonner** — toast notifications

## Getting Started

### Prerequisites

- [Node.js 18+](https://nodejs.org/)
- [gonna-pay-backend](../gonna-pay-backend) running on `http://localhost:3333`

### Environment Variables

```bash
cp .env.local.example .env.local
```

Fill in the required values in `.env.local`.

### Running Locally

```bash
npm install
npm run dev
```

The app will be available at `http://localhost:3000`.

## Commands

```bash
npm run dev    # Start dev server (Turbopack)
npm run build  # Production build
npm run start  # Start production server
npm run lint   # Run ESLint with auto-fix
```
