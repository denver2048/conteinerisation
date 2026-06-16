# Demo App — Containerization homework

Two apps live in this repo:

- **`app/`** — minimal FastAPI service (`GET /`, `GET /health`). Lecture 2 (package your app).
- **`app_3/`** — FastAPI + PostgreSQL (`GET /`, `GET /db`). Lecture 3 (build a system).

## Run locally (optional)
```bash
pip install -r requirements.txt
uvicorn app:app --host 0.0.0.0 --port 8000
```

## Lecture 2 — `app/Dockerfile`
```bash
cd app
DOCKER_BUILDKIT=1 docker build -t demo-app .
docker run --rm -p 8000:8000 demo-app
curl localhost:8000/         # {"message":"Hello from Docker course 🚀"}
curl localhost:8000/health   # {"status":"ok"}
```
- **L1 — works:** builds and serves on `:8000`.
- **L2 — better:** `requirements.txt` copied before the code (install layer stays cached on code-only changes); BuildKit `--mount=type=cache` keeps pip cache out of the layer; `python:3.12-slim` base; `.dockerignore` trims context.
- **L3 — production:** multi-stage (deps in `/opt/venv`, copied into a clean runtime stage); slim base; non-root `appuser`; `HEALTHCHECK` on `/health`.

## Lecture 3 — `app_3/` (`docker-compose.yaml`)
```bash
cd app_3
cp .env.example .env        # set DB_PASSWORD
docker compose up --build
curl localhost:8000/        # {"message":"Hello from Docker Compose 🚀"}
curl localhost:8000/db      # {"db_response":[1]}
```
- **L1 — connect services:** `app` + `db` (`postgres:16-alpine`); app reaches Postgres by service name (`DB_HOST=db`).
- **L2 — stable:** named volume `pgdata`; `restart: unless-stopped`; `.env` config (`DB_PASSWORD` feeds app `DB_PASSWORD` and Postgres `POSTGRES_PASSWORD`).
- **L3 — production:** healthchecks (app `/`, Postgres `pg_isready`); `depends_on: condition: service_healthy`; two networks — `frontend` (app published) and `backend` `internal: true` (Postgres unreachable from host); `${DB_PASSWORD:?...}` fails fast if unset; `no-new-privileges:true` on both.

## Endpoints
- `app/`: `GET /`, `GET /health`
- `app_3/`: `GET /`, `GET /db`
