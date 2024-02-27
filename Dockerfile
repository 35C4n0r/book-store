# Stage 1: Build executable
FROM golang:1.21.0 as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o app cmd/server/main.go

# Stage 2: Create final image
FROM alpine:latest
WORKDIR /root
COPY --from=builder /app/app .
CMD ["./app"]
