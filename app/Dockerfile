# syntax=docker/dockerfile:1

# -----------------------------------------------------------------------------
# Stage 1: resolve and install dependencies (cached independently of app code)
# -----------------------------------------------------------------------------
FROM python:3.11-slim-bookworm AS builder

WORKDIR /build

COPY requirements.txt .
RUN --mount=type=cache,target=/root/.cache/pip \
    pip install --upgrade pip && \
    pip install --no-cache-dir --target=/deps -r requirements.txt

# -----------------------------------------------------------------------------
# Stage 2: minimal runtime (distroless)
# -----------------------------------------------------------------------------
FROM gcr.io/distroless/python3-debian12

WORKDIR /app

ENV PYTHONPATH=/deps \
    PYTHONUNBUFFERED=1 \
    PYTHONDONTWRITEBYTECODE=1

COPY --from=builder --chown=65532:65532 /deps /deps
COPY --chown=65532:65532 app.py /app/app.py

USER 65532:65532

EXPOSE 8000

ENTRYPOINT ["/usr/bin/python3", "-m", "uvicorn", "app:app", "--host", "0.0.0.0", "--port", "8000"]
