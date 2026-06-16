FROM python:3.11-slim AS builder

WORKDIR /app

RUN groupadd -r appuser && useradd -r -g appuser -u 10001 appuser

COPY requirements.txt .
RUN pip install --no-cache-dir --target=/app/deps -r requirements.txt

COPY app.py .

RUN chown -R appuser:appuser /app/deps /app/app.py


FROM gcr.io/distroless/python3-debian12

WORKDIR /app

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

COPY --from=builder /app/deps /app/deps
COPY --from=builder /app/app.py .

ENV PYTHONPATH=/app/deps

USER appuser

EXPOSE 8000

ENTRYPOINT ["python3", "-m", "uvicorn", "app:app", "--host", "0.0.0.0", "--port", "8000"]
