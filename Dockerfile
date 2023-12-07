
# Stage 1: Build the Go binary
FROM golang:1.19 AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN --mount=type=cache,target=/tmp CGO_ENABLED=0 GOOS=linux go build -o chrnr-cli .


# Stage 2: Create a minimal runtime image
FROM scratch

COPY --from=builder /chrnr-cli .

CMD ["./chrnr-cli"]

