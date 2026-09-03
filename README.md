# Scaffolding for coding interview

The repository contains a starting point for the coding excercise given to you by mail. It contains:

- A backend repository running a golang application (with postgres and air for hot reloading) in `/backend`
- A frontend repository running React with Typescript (with Vite as the development server) in `/frontend`

## Prerequisites

- Docker with Compose — runs the backend and the database
- Node.js 22.12+ (20.19+ also works) — runs the frontend
- Go 1.27 — only needed if you want to run the backend outside Docker

## Getting started

Start the database and the API:

```bash
cd backend
docker compose up
```

Then, in a second terminal, start the frontend:

```bash
cd frontend
npm install
npm run dev
```

The API listens on http://localhost:8080 and the app on http://localhost:5173. The backend allows
the dev server's origin via CORS, so the frontend calls the API directly. Use this as an
end-to-end smoke test once both sides are running:

```bash
curl http://localhost:8080/health
# {"status":"ok"}
```

See `backend/README.md` and `frontend/README.md` for details on each side.
