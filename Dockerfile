ARG ARG_PYTHON_VERSION=3.14.5
FROM python:${ARG_PYTHON_VERSION}-alpine

ARG ARG_UVICORN_PORT=80

USER root

RUN addgroup -S appgroup && \
    adduser -S appuser -G appgroup

WORKDIR /app

COPY ./app/requirements.txt .
RUN pip install --no-cache-dir --upgrade pip && \
    pip install --no-cache-dir -r requirements.txt && \
    \
    python -m pip uninstall -y pip setuptools wheel && \
    rm -rf /root/.cache /usr/local/bin/pip*

COPY ./app/app.py .

USER appuser

ENV UVICORN_PORT=$ARG_UVICORN_PORT
ENTRYPOINT ["python", "-m", "uvicorn", "--host", "0.0.0.0", "app:app"]

HEALTHCHECK \
 #--interval=5s --timeout=5s --start-period=1s --retries=1 \
 CMD wget -qO- http://127.0.0.1:$UVICORN_PORT/health | grep -q \"status\":\"ok\" || exit 1

EXPOSE $ARG_UVICORN_PORT

