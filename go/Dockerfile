# BUILDER
FROM golang:1.26.3 AS builder

WORKDIR /app

COPY go.mod .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o server .

# RUNTIME
FROM scratch

WORKDIR /app

COPY --from=builder /app/server /app/server

EXPOSE 8090
USER 101

CMD ["/app/server"] 