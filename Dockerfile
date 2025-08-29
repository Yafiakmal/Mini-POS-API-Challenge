# Stage 1: build binary
FROM golang:alpine3.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o build/main cmd/main.go

# Stage 2: final image
FROM alpine:latest

WORKDIR /app

# Install Postgres client + timezone data
RUN apk add --no-cache postgresql-client tzdata

# Set timezone default (optional)
ENV TZ=Asia/Jakarta

# Copy binary & wait script
COPY --from=builder /app/build/main .
# COPY --from=builder /app/.env .
# COPY wait-for-postgres.sh .
# RUN chmod +x wait-for-postgres.sh

EXPOSE 8080

CMD ["./main"]


