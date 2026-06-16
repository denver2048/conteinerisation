# Homework for lecture 2


## Level 1

Dockerfile

    FROM python:3.13
    WORKDIR /app
    COPY . .
    RUN pip install -r ./requirements.txt
    EXPOSE 8000
    CMD ["uvicorn", "app:app", "--host=0.0.0.0", "--port", "8000"]

Building

    docker build . -t hw_lec2:1.0

Running

    docker run -p 8000:8000 hw_lec2:1.0

## Level 2

Dockerfile with optimization and caching
    

    ARG IMAGE_VERSION="3.13"
    FROM python:${IMAGE_VERSION}
    WORKDIR /app
    COPY requirements.txt .
    RUN pip install --no-cache-dir --upgrade -r ./requirements.txt
    COPY . .
    EXPOSE 8000
    CMD ["uvicorn", "app:app", "--host=0.0.0.0", "--port", "8000"]

Reduce image size by using alpine or slim images instead of defaulf.
Changing argument IMAGE_VERSION and building.

    

    docker build . -t hw_lec2:1.1
    docker build . --build-arg IMAGE_VERSION="3.13-slim" -t hw_lec2:1.1-slim
    docker build . --build-arg IMAGE_VERSION="3.13-alpine" -t hw_lec2:1.1-alpine

Size of images:

    docker images
    REPOSITORY                       TAG          IMAGE ID       CREATED              SIZE 
    hw_lec2                          1.1          521142d326b1   15 seconds ago       1.13GB
    hw_lec2                          1.1-slim     97bca6a0c290   7 seconds ago        141MB
    hw_lec2                          1.1-alpine   38e835038e8f   7 seconds ago        75.7MB


## Level 3

Dockerfile with multi-stage building, runtime use slim image and user without root privileges.

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

