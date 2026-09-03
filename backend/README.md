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
- `GET /health` — pings the database and returns `{"status":"ok"}`, or a 503 with
  `{"status":"unavailable"}` when the database cannot be reached

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

The server exits on startup if the database is unreachable, so check the compose logs first if
the container keeps restarting.
