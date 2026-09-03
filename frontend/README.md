# React + TypeScript + Vite

This template provides a minimal setup to get React working in Vite with HMR and some ESLint
rules. It requires Node.js 22.13+ (20.19+ and 24+ also work).

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

The dev server proxies `/api/*` to the backend on http://localhost:8080 and strips the `/api`
prefix, so no CORS setup is needed. With `docker compose up` running in `/backend`:

```bash
curl http://localhost:5173/api/health
```

should return `{"status":"ok"}`.

## Other scripts

- `npm run lint` — ESLint
- `npm run build` — type-check and build for production
- `npm run preview` — serve the production build
