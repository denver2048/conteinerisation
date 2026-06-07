# BUILDER
FROM python:3.12-slim AS builder

WORKDIR /app

COPY requirements.txt .
RUN pip wheel --wheel-dir=/app/wheels -r requirements.txt

# RUNTIME
FROM python:3.12-slim

WORKDIR /app

RUN useradd -m pythonuser

COPY --from=builder /app/wheels /app/wheels
RUN pip install --no-cache-dir /app/wheels/*

COPY . .

RUN chown -R pythonuser:pythonuser /app \
    && chmod 755 /app 

EXPOSE 8090
USER pythonuser

CMD ["uvicorn", "app:app", "--host", "0.0.0.0", "--port", "8090"]