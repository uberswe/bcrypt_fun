FROM golang:1.16.5-alpine AS builder

WORKDIR /app
COPY . .

RUN apk --no-cache add ca-certificates

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o "main" -ldflags="-w -s" ./cmd/bcrypt/main.go

FROM scratch

COPY --from=builder /app/main /usr/bin/
COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs/

CMD ["main"]

EXPOSE 80