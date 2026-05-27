FROM golang:alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG SERVICE_NAME
RUN CGO_ENABLED=0 GOOS=linux go build -o service ./${SERVICE_NAME}

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/service .
CMD ["./service"]