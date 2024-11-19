FROM golang:1.23-alpine AS builder

RUN apk update && apk add --no-cache bash tzdata gcc musl-dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=1 GOOS=linux GOARCH=amd64

RUN go build -o goatrobotics

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/goatrobotics .
COPY --from=builder /app/config.json .
RUN mkdir UI
COPY --from=builder /app/UI UI/

RUN mkdir -p logs && touch logs/Audit.audit logs/GOATROBOTICS.log

ENTRYPOINT ["./goatrobotics"]
