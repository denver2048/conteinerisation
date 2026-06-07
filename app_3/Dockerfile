# syntax=docker/dockerfile:1

FROM python:3.11-slim AS build

WORKDIR /app

COPY requirements.txt .
RUN --mount=type=cache,target=/root/.cache/pip \
    pip install --prefix=/install -r requirements.txt


    FROM gcr.io/distroless/python3-debian12:nonroot

WORKDIR /app

COPY --from=build /install /usr/local
COPY app.py .

ENV PYTHONPATH=/usr/local/lib/python3.11/site-packages \
    PYTHONUNBUFFERED=1 \
    PYTHONDONTWRITEBYTECODE=1

USER nonroot
EXPOSE 8000

CMD ["-m", "uvicorn", "app:app", "--host", "0.0.0.0", "--port", "8000"]
