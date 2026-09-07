# React + TypeScript + Vite

This template provides a minimal setup to get React working in Vite with HMR and some lint rules.
It requires Node.js 22.13+ (20.19+ and 24+ also work).

1. Install dependencies

```bash
npm install
```

2. Start the dev-server

```bash
npm run dev
```

The app is served on http://localhost:5173.

## Talking to the backend

The API runs on http://localhost:8080 and allows this origin via CORS, so you can call it
directly — no dev-server proxy involved:

```ts
const res = await fetch('http://localhost:8080/health')
```

With `docker compose up` running in `/backend`, `curl http://localhost:8080/health` should return
`{"status":"OK"}`.

## Tooling notes

- **oxlint** does the linting (`npm run lint`). Rules live in `.oxlintrc.json`. The React plugin
  is off in oxlint by default, so the config enables it explicitly along with the hooks rules.
- **React Compiler** is enabled in `vite.config.ts` via `react({ compiler: true })`, so components
  are memoised for you.
- `tsconfig.app.json` type-checks `src` (browser libs, JSX). `tsconfig.node.json` type-checks
  `vite.config.ts` itself (Node libs). `tsconfig.json` just references both so `tsc -b` builds them.

## Other scripts

- `npm run lint` — oxlint
- `npm run build` — type-check and build for production
- `npm run preview` — serve the production build
