# Backend - Running the application:

To run the application simply run docker compose in the terminal of your choice:

```bash
docker compose up
```

This should start the database on port 5432 and the backend server on port 8080. To verify that
the database has started correctly read the logs in your terminal (and try to connect with any
database tool of your choice). To verify the server go to http://localhost:8080 in your browser
and make sure you see the "Server running OK" message.

## Endpoints

- `GET /` — plain-text liveness message, "Server running OK"
- `GET /health` — pings the database and returns `{"status":"OK"}`, or a 503 with
  `{"status":"Service Unavailable"}` when the database cannot be reached

The server sends CORS headers for `CORS_ORIGIN` (the frontend dev server by default), so the
frontend can call these endpoints directly from the browser. `CORS_ORIGIN` is a single origin,
written exactly as the browser sends it — scheme, host and port, with no trailing slash.

## Configuration

The server reads its database connection from the environment. Defaults in brackets apply when
running outside Docker:

| Variable | Default |
|---|---|
| `DB_HOST` | `localhost` |
| `DB_PORT` | `5432` |
| `DB_USER` | `postgres` |
| `DB_PASSWORD` | `postgres` |
| `DB_NAME` | `main_db` |
| `CORS_ORIGIN` | `http://localhost:5173` |

The server exits on startup if the database is unreachable. Under Compose it waits for the
database's healthcheck before it starts, so a cold start is not a race. Started on its own against
a database that is not up yet, it exits and stays down until you start it again.

## Database data

The Postgres data lives in the `postgres_data_18` volume. Postgres 18 keeps its cluster in
`/var/lib/postgresql/18/docker`, so the volume is named per major version: a cluster created by an
older image is not readable by this one, and reusing the old volume would silently start an empty
database. To start over, remove the volume:

```bash
docker compose down -v
```
