ARG BUILDER_IMAGE_VERSION="3.13"
ARG RUNTIME_IMAGE_VERSION="3.13-slim"
ARG NONROOT="user42"

FROM python:${BUILDER_IMAGE_VERSION} AS builder

ENV PYTHONDONTWRITEBYTECODE=1 \
    PYTHONUNBUFFERED=1 \
    PATH="/opt/venv/bin:$PATH"

WORKDIR /app

RUN apt-get update && apt-get install -y --no-install-recommends \
    build-essential \
    && rm -rf /var/lib/apt/lists/* \
    && python -m venv /opt/venv


COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt


FROM python:${RUNTIME_IMAGE_VERSION} AS runtime

ARG NONROOT

ENV PYTHONDONTWRITEBYTECODE=1 \
    PYTHONUNBUFFERED=1 \
    PATH="/opt/venv/bin:$PATH"

WORKDIR /app

COPY --from=builder /opt/venv /opt/venv

COPY . .

RUN useradd --create-home ${NONROOT} && chown -R ${NONROOT}:${NONROOT} /app
USER ${NONROOT}

EXPOSE 8000

CMD ["uvicorn", "app:app", "--host", "0.0.0.0", "--port", "8000"]

