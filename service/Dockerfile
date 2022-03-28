FROM golang:1.18-alpine as builder
COPY . /app
WORKDIR /app/service
RUN go get ./...
RUN CGO_ENABLE=0 GOOS=linux go build -o service ./cmd/main.go

FROM alpine
WORKDIR /app
COPY --from=builder /app/service/service service
CMD ["./service"]
