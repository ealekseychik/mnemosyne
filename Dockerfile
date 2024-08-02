# Stage 1: Build the application
FROM golang:1.22.3 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /mnemosyne

# Stage 2: Run the application
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /mnemosyne .
COPY .env .
COPY static/ ./static/

CMD ["./mnemosyne"]
