# STAGE 1: builder
FROM python:3.11-slim AS builder

# Створення робочої директорії білдера
WORKDIR /build

# Копіюємо залежності
COPY requirements.txt .

# Встановлюємо залежності у локальний шлях
RUN pip install --no-cache-dir --prefix=/install -r requirements.txt

# STAGE 2: runtime
FROM python:3.11-slim

# Створення non-root користувача
RUN useradd -m appuser

# Створення робочої директорії додатку
WORKDIR /app

# Копіюємо встановлені залежності з builder stage
COPY --from=builder /install /usr/local

# Копіюємо код застосунку
COPY . /app/

# Перехід на non-root user
USER appuser

# Порт контейнера
EXPOSE 8080

# Запуск FastAPI через uvicorn
CMD ["uvicorn", "app:app", "--host", "0.0.0.0", "--port", "8080"]