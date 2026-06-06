# ---- builder ----
    FROM golang:1.22 AS builder

    WORKDIR /app
    
    COPY go.mod go.sum ./
    RUN go mod download
    
    COPY . .
    
    RUN CGO_ENABLED=0 GOOS=linux go build -o app
    
    # ---- runtime ----
    FROM gcr.io/distroless/base
    
    WORKDIR /app
    
    COPY --from=builder /app/app /app
    
    EXPOSE 8080
    
    CMD ["/app/app"]