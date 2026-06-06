# Домашнє завдання №1

## Endpoints

- GET /
- GET /health

## Dockerfile

Обрано найменший базовий docker image.

```dockerfile
ARG ARG_PYTHON_VERSION=3.14.5
FROM python:${ARG_PYTHON_VERSION}-alpine
```

Оскільки номер порта використовується в трьох місцях Dockerfile, а саме
у інструкціях ENTRYPOINT, HEALTHCHECK, та EXPOSE, я вирішив створити build argument:

```dockerfile
ARG ARG_UVICORN_PORT=80
```

Всі наступні команди будуть виконуватися з правами root. З технічної точки зору цей рядок коду не потрібен, бо він прописаний в базовому docker image. Але я додав його по двом причинам:

- якщо зміниться user в базовому docker image, то це ніяк не вплине на цей Dockerfile
- для документування

```dockerfile
USER root
```

Створюється user.

```dockerfile
RUN addgroup -S appgroup && \
    adduser -S appuser -G appgroup
```

Встановлюється поточна директорія.

```dockerfile
WORKDIR /app
```

Інсталюються залежності програми.

1. Копіюється requirements.txt зі списком залежностей.
2. Оновлюється пакетний менеджер pip та встановлюються залежності.
3. Видаляється пакетний менеджер pip разом з усіма кешами.

```dockerfile
COPY ./app/requirements.txt .
RUN pip install --no-cache-dir --upgrade pip && \
    pip install --no-cache-dir -r requirements.txt && \
    \
    python -m pip uninstall -y pip setuptools wheel && \
    rm -rf /root/.cache /usr/local/bin/pip*
```

Копіювання файлів програми. Права на читання файлів та відкриття папки матимуть усі user-и, тому я не пишу код, що міняє права доступа і власника файлів. Порядок, в якому створюються шари - спочатку шар залежностей (міняється рідше), потім шар коду (міняється частіше), обрано з метою оптимізації процесу побудови docker image.

```dockerfile
COPY ./app/app.py .
```

Встановлюється поточний user контейнера.

```dockerfile
USER appuser
```

Задається команда, яка запуститься при старті docker container з правами користувача appuser. Довелося використати змінну оточення $UVICORN_PORT, бо exec form запису команди в ENTRYPOINT на працює з підстановкою $ARG_UVICORN_PORT.

```dockerfile
ENV UVICORN_PORT=$ARG_UVICORN_PORT
ENTRYPOINT ["python", "-m", "uvicorn", "--host", "0.0.0.0", "app:app"]
```

Health check. Другий рядок закоментований, використовується при тестуванні. Команда запускається з правами користувача appuser.

```dockerfile
HEALTHCHECK \
 #--interval=5s --timeout=5s --start-period=1s --retries=1 \
 CMD wget -qO- http://127.0.0.1:$UVICORN_PORT/health | grep -q \"status\":\"ok\" || exit 1
```

Документується факт того, що всередині контейнера unicorn буде використовувати $ARG_UVICORN_PORT.

```dockerfile
EXPOSE $ARG_UVICORN_PORT
```
