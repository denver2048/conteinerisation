# Demo App

## Run locally (optional)
pip install -r requirements.txt
uvicorn app:app --host 0.0.0.0 --port 8000

## Endpoints
- GET /
- GET /health

## Docker

Each app ships a multi-stage `Dockerfile` (distroless, non-root) and a `docker-compose.yaml`.

- `app/` — standalone FastAPI service on port 8000.
- `app_3/` — FastAPI + PostgreSQL. Two networks (`frontend` public, `backend` internal),
  persistent DB volume, healthchecks, restart policies and hardened containers.

### Run

```bash
cd app          # or: cd app_3
cp .env.example .env   # app_3 only — set DB_PASSWORD
docker compose up --build
```