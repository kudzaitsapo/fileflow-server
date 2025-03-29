FROM golang:1.24.0 AS builder

WORKDIR /bin
COPY . .

# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server cmd/server/*.go

# FROM scratch
# WORKDIR /bin
# # Copy CA certificates
# COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# COPY --from=builder /bin/server .

# EXPOSE 8195

# CMD ["./server"]


RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
COPY air.toml ./
RUN go mod download

CMD ["air", "-c", "air.toml"]