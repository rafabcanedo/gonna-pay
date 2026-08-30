# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
npm run dev      # Start dev server (Next.js + Turbopack)
npm run build    # Production build (Turbopack)
npm run start    # Start production server
npm run lint     # Run ESLint with auto-fix on src/**/*.{ts,tsx}
```

No test runner is configured.

## Tech Stack

- **Next.js 15** (App Router) + **React 19** + **TypeScript**
- **Tailwind CSS v4** + **Shadcn/ui** (New York style, RSC enabled)
- **TanStack React Query v5** — server state, queries, mutations
- **React Hook Form v7** + **Yup** — form state and validation
- **Sonner** — toast notifications
- **Lucide React** — icons

## Architecture

### Route Groups

```
src/app/
├── (auth)/          # Public routes: /signin, /signup, /forgot-your-password
├── (dashboard)/     # Protected routes: /dashboard, /my-wallet, /payments, /costs, /my-contacts, /profile
└── api/             # Next.js Route Handlers (proxy layer to Go backend)
```

The middleware at [src/middleware.ts](src/middleware.ts) guards all routes by checking the `access_token` cookie — it does not validate JWTs (delegated to backend).

### Data Flow

```
Page/Component
  → custom hook (src/hooks/queries/ or src/hooks/mutations/)
    → service (src/services/)
      → API client (src/lib/api-client/)
        → Next.js API route (src/app/api/)
          → Go backend
```

- **Queries**: `src/hooks/queries/` — TanStack `useQuery` wrappers per domain
- **Mutations**: `src/hooks/mutations/` — TanStack `useMutation` wrappers per domain
- **Services**: `src/services/` — pure async functions, no React, handle CRUD per entity
- **API Client**: `src/lib/api-client/index.ts` — single `apiCall<T>()` with typed error handling (`ValidationError`, `UnauthorizedError`, etc.)

### Key Conventions

- **Server Components by default**; use `'use client'` only when needed (interactivity, hooks, browser APIs)
- **Absolute imports** via `@/*` alias (maps to `src/*`)
- **Shadcn components** live in `src/components/ui/`; domain components in `src/components/(auth)/` or `src/components/(dashboard)/`
- **Validation schemas** in `src/validations/schemas.ts` using Yup; reusable field validators in `src/validations/fields.ts`
- **Types** centralized in `src/types/index.ts`; includes domain enums like `ContactCategory` and `CostCategory`

### React Query Config

Configured in [src/lib/react-query/index.tsx](src/lib/react-query/index.tsx):
- `staleTime: 60000` (1 minute)
- `refetchOnWindowFocus: false`
- Devtools included (initially closed)

## Project Standards

### Colocation

Colocation means keeping related files close to where they are used. If a file (types, constants, helpers, subcomponents) only exists because of a specific component or route, it should live next to it — not in a global folder.

Only move to a shared folder (`src/types/`, `src/hooks/`, etc.) when the file is genuinely reused across multiple places.

**How it's applied in this project:**

Components with enough logic to split responsibilities adopt the following structure:

```
src/components/[component]/
├── index.tsx          # Component — rendering only
├── constants/
│   └── index.ts       # Static data scoped to this component
├── interfaces/
│   └── index.ts       # Interface declarations only
└── types/
    └── index.ts       # Type declarations only — or mixed (used as generic)
```

**When to use each subfolder:**
- `constants/` — static data arrays, config objects, maps
- `interfaces/` — when the file has **only** `interface` declarations
- `types/` — when the file has **only** `type` declarations, or a mix of both `type` and `interface`

The same pattern applies to routes inside `src/app/(dashboard)/`:

```
src/app/(dashboard)/[feature]/
├── page.tsx
├── constants/
│   └── index.ts
├── interfaces/
│   └── index.ts
├── types/
│   └── index.ts
├── components/
│   └── [component]/
│       ├── index.tsx
│       ├── constants/
│       │   └── index.ts
│       ├── interfaces/
│       │   └── index.ts
│       └── types/
│           └── index.ts
└── [sub-route]/
    ├── page.tsx
    └── components/
        └── [component]/
            ├── index.tsx
            ├── interfaces/
            │   └── index.ts
            └── types/
                └── index.ts
```
