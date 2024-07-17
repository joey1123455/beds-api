# Build stage
FROM golang:latest as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
COPY .env .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api ./cmd/api/.


# Intermediate stage for downloading additional packages
FROM alpine:latest as intermediate
RUN apk update && apk add --no-cache git ca-certificates make gcc musl-dev tzdata wget
RUN apk update && apk add postgresql-client
RUN wget --no-check-certificate https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz && \
    tar -xvf migrate.linux-amd64.tar.gz && \
    mv migrate /usr/local/bin/migrate && \
    rm migrate.linux-amd64.tar.gz

# Final stage
FROM alpine:latest
RUN apk update && apk add --no-cache bash
WORKDIR /root/
COPY --from=builder /app/api .
COPY --from=builder /app/wait-for-it.sh .
COPY --from=builder /app/internal/db/migrations /root/internal/db/migrations
COPY --from=builder /app/internal/mailer/templates /root/internal/mailer/templates
COPY --from=builder /app/.env /root/.env
COPY --from=intermediate /usr/local/bin/migrate /usr/local/bin/migrate

# Make the shell script executable
RUN chmod +x wait-for-it.sh

ENV DB_URL="postgresql://root:secret@db:5432/flux_core?sslmode=disable"

EXPOSE 8080
CMD ["sh", "-c", "./wait-for-it.sh db:5432 -- migrate -database \"${DB_URL}\" -path internal/db/migrations up && ./api"]
