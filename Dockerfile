
# Stage 1: Build the Go binary
FROM golang:1.16 AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN --mount=type=cache,target=/tmp CGO_ENABLED=0 GOOS=linux go build -o chrnr-cli .


# Stage 2: Create a minimal runtime image
FROM scratch

WORKDIR /app

COPY --from=builder /app/chrnr-cli .

CMD ["./app/chrnr-cli"]

